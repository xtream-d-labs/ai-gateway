package kubernetes

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"

	"github.com/go-openapi/swag"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	containerName = "main"
)

// CreateJob creates a new kubernetes job
func CreateJob(kubeConfig, jobID, namespace, imageName string, cpu, gpu int64) error {
	clientset, e := client(kubeConfig)
	if e != nil {
		return e
	}
	pod := coreV1.PodTemplateSpec{
		Spec: coreV1.PodSpec{
			Containers: []coreV1.Container{
				coreV1.Container{
					Name:  containerName,
					Image: imageName,
					Resources: coreV1.ResourceRequirements{
						Requests: coreV1.ResourceList{
							"cpu": resource.MustParse(fmt.Sprintf("%d", cpu)),
						},
					},
				},
			},
			RestartPolicy: coreV1.RestartPolicyNever,
		},
	}
	if gpu > 0 {
		// https://kubernetes.io/docs/tasks/manage-gpus/scheduling-gpus/
		pod.Spec.Containers[0].Resources.Limits = coreV1.ResourceList{
			"nvidia.com/gpu": resource.MustParse(fmt.Sprintf("%d", gpu)),
		}
	}
	_, e = clientset.BatchV1().Jobs(namespace).Create(&batchV1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("job-%s", jobID),
		},
		Spec: batchV1.JobSpec{
			Completions:  swag.Int32(1),
			Parallelism:  swag.Int32(1),
			BackoffLimit: swag.Int32(1),
			Template:     pod,
		},
	})
	return e
}

// JobStatus returns status of the specified job
func JobStatus(kubeConfig, jobID, namespace string) (*batchV1.JobStatus, error) {
	clientset, e := client(kubeConfig)
	if e != nil {
		return nil, e
	}
	name := fmt.Sprintf("job-%s", jobID)
	job, e := clientset.BatchV1().Jobs(namespace).Get(name, metav1.GetOptions{})
	if e != nil {
		return nil, e
	}
	return &job.Status, nil
}

// PodStatus returns status of the specified pod
func PodStatus(kubeConfig, jobID, namespace string) (*coreV1.Pod, error) {
	clientset, e := client(kubeConfig)
	if e != nil {
		return nil, e
	}
	pods, e := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("job-name=job-%s", jobID),
	})
	if e != nil {
		return nil, e
	}
	if len(pods.Items) == 0 {
		return nil, fmt.Errorf("Not found")
	}
	return &pods.Items[0], nil
}

// Log the log structure
type Log struct {
	Time time.Time
	Log  string
}

var (
	reNewline = regexp.MustCompile(`\n`)
	reDateLog = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.*Z)(.*)`)
	reAnsiSeq = regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
)

// Logs retrieve the pod log
func Logs(kubeConfig, jobID, namespace string) ([]*Log, error) {
	clientset, e := client(kubeConfig)
	if e != nil {
		return nil, e
	}
	pods, e := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("job-name=job-%s", jobID),
	})
	if e != nil {
		return nil, e
	}
	if len(pods.Items) == 0 {
		return nil, fmt.Errorf("Not found")
	}
	req := clientset.CoreV1().Pods(namespace).GetLogs(pods.Items[0].ObjectMeta.Name, &coreV1.PodLogOptions{
		Container:  containerName,
		Timestamps: true,
		Follow:     false,
	})
	out, e := req.Stream()
	if e != nil {
		return nil, e
	}
	defer out.Close()

	buf := bytes.Buffer{}
	io.Copy(&buf, out)

	result := []*Log{}
	for _, line := range reNewline.Split(buf.String(), -1) {
		values := reDateLog.FindStringSubmatch(line)
		if len(values) < 3 {
			continue
		}
		value, e := time.Parse(time.RFC3339Nano, values[1])
		if e != nil {
			continue
		}
		if values[2] == "" {
			continue
		}
		result = append(result, &Log{
			Time: value,
			Log:  reAnsiSeq.ReplaceAllString(values[2], ""),
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Time.Before(result[j].Time)
	})
	return result, nil
}

// DeleteJob deletes the specified job
func DeleteJob(kubeConfig, jobID, namespace string) error {
	clientset, e := client(kubeConfig)
	if e != nil {
		return e
	}
	pods, e := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("job-name=job-%s", jobID),
	})
	if e != nil {
		return e
	}
	if len(pods.Items) == 0 {
		return fmt.Errorf("Not found")
	}
	if e = clientset.CoreV1().Pods(namespace).Delete(pods.Items[0].ObjectMeta.Name, &metav1.DeleteOptions{
		GracePeriodSeconds: swag.Int64(0),
	}); e != nil {
		return e
	}
	name := fmt.Sprintf("job-%s", jobID)
	if e = clientset.BatchV1().Jobs(namespace).Delete(name, &metav1.DeleteOptions{
		GracePeriodSeconds: swag.Int64(0),
	}); e != nil {
		return e
	}
	return nil
}

func client(kubeConfig string) (*kubernetes.Clientset, error) {
	config, e := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfig))
	if e != nil {
		return nil, e
	}
	clientset, e := kubernetes.NewForConfig(config)
	if e != nil {
		return nil, e
	}
	return clientset, nil
}
