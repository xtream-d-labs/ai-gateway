package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
)

// Image state
const (
	StableImage    = "stable"   // Normal docker images
	PullingImage   = "pulling"  // Pull a specified image
	BuildingImage  = "building" // Build a Jupyter notebook image
	BuildingJob    = "building-job"
	PushingJob     = "pushing-job"
	KubernetesJob  = "k8s-job"
	K8sJobStart    = "k8s-job-start"
	K8sJobPending  = "k8s-job-pending"
	K8sJobRunning  = "k8s-job-runnning"
	K8sSucceeded   = "k8s-job-succeeded"
	K8sFailed      = "k8s-job-failed"
	RescaleSend    = "rescale-send"
	RescaleStart   = "rescale-start"
	RescaleRunning = "rescale-runnning"
	RescaleSucceed = "rescale-succeeded"
	RescaleFailed  = "rescale-failed"
	StatusUnknown  = "unknown"
)

// Image represents cached image information
type Image struct {
	Tag     string
	Status  string
	Started time.Time
}

// PlatformType defines in where the job will be run
type PlatformType int

const (
	PlatformKubernetes PlatformType = iota
	PlatformRescale
)

func (p PlatformType) String() string {
	if p == PlatformKubernetes {
		return models.JobPlatformKubernetes
	}
	return models.JobPlatformRescale
}

// Job represents cached job information
type Job struct {
	Platform    PlatformType
	ID          string
	Status      string
	DockerImage string
	PythonFile  string
	Workspaces  []string
	Commands    []string
	CPU         int64
	Memory      int64
	GPU         int64
	CoreType    string
	Cores       int64
	Started     time.Time
	TargetID    string // Job ID of the target platform
}

// GetPullingImages are images under PULLING state
func GetPullingImages() ([]*Image, error) {
	return getImages(PullingImageStoreKey)
}

// SetPullingImageMeta add an image to PULLING cache
func SetPullingImageMeta(name string) error {
	return setImage(PullingImageStoreKey, &Image{
		Tag:     name,
		Status:  PullingImage,
		Started: time.Now(),
	})
}

// RemovePullingImage removes from PULLING cache
func RemovePullingImage(name string) error {
	return removeImage(PullingImageStoreKey, name)
}

// GetBuildingImages are images under BUILDING state
func GetBuildingImages() ([]*Image, error) {
	return getImages(BuildingImageStoreKey)
}

// SetBuildImageMeta add an image to BUILDING cache
func SetBuildImageMeta(name string) error {
	return setImage(BuildingImageStoreKey, &Image{
		Tag:     name,
		Status:  BuildingImage,
		Started: time.Now(),
	})
}

// RemoveBuildingImage removes from BUILDING cache
func RemoveBuildingImage(name string) error {
	return removeImage(BuildingImageStoreKey, name)
}

// GetJobs returns jobs under BUILDING job images / Rescale status
func GetJobs() ([]*Job, error) {
	jobs, err := getJobs(BuildingJobStoreKey)
	if err != nil {
		return nil, err
	}
	additionals, err := getJobs(PushingJobStoreKey)
	if err != nil {
		return nil, err
	}
	jobs = append(jobs, additionals...)
	additionals, err = getJobs(KubernetesJobStoreKey)
	if err != nil {
		return nil, err
	}
	jobs = append(jobs, additionals...)
	additionals, err = getJobs(RescaleJobStoreKey)
	if err != nil {
		return nil, err
	}
	jobs = append(jobs, additionals...)
	return jobs, nil
}

// GetJob returns job specified
func GetJob(ID string) (*Job, error) {
	jobs, err := GetJobs()
	if err != nil {
		return nil, err
	}
	for _, job := range jobs {
		if job.ID == ID {
			return job, nil
		}
	}
	return nil, fmt.Errorf("A specified job is not found")
}

// SetJobMeta add a task to BUILD job images
func SetJobMeta(job *Job) error {
	if job.Status != BuildingJob {
		return fmt.Errorf("%s is not proper status for building job images", job.ID)
	}
	return setJob(BuildingJobStoreKey, job)
}

// RemoveBuildingJobImagesJobs removes job
func RemoveBuildingJobImagesJobs(ID string) error {
	return removeJob(BuildingJobStoreKey, ID)
}

// RemovePushingJobImageJobs removes job
func RemovePushingJobImageJobs(ID string) error {
	return removeJob(PushingJobStoreKey, ID)
}

// UpdateJob update job status
func UpdateJob(ID, from, to, status string) error {
	return UpdateJobDetail(ID, from, to, status, nil)
}

// UpdateJobDetail update job status and targetID
func UpdateJobDetail(ID, from, to, status string, targetID *string) error {
	return SetValue(func(txn *badger.Txn) error {
		jobs, err := getCachedJobs(txn, from)
		if err != nil {
			return err
		}
		remains := []*Job{}
		var target *Job
		for _, job := range jobs {
			if strings.EqualFold(ID, job.ID) {
				target = job
				target.Status = status
				if targetID != nil {
					target.TargetID = swag.StringValue(targetID)
				}
			} else {
				remains = append(remains, job)
			}
		}
		bytes, err := json.Marshal(remains)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(from), bytes)
		if err != nil {
			return err
		}
		jobs, err = getCachedJobs(txn, to)
		if err != nil {
			return err
		}
		if target == nil {
			return nil
		}
		if idInSlice(target.ID, jobs) {
			return nil
		}
		jobs = append(jobs, target)
		bytes, err = json.Marshal(jobs)
		if err != nil {
			return err
		}
		return txn.Set([]byte(to), bytes)
	})
}

// RemoveJob removes a job you specified
func RemoveJob(ID string) error {
	if err := removeJob(KubernetesJobStoreKey, ID); err != nil {
		return err
	}
	if err := removeJob(RescaleJobStoreKey, ID); err != nil {
		return err
	}
	return nil
}

// Cache keys
const (
	PullingImageStoreKey  = "pulling-images"
	BuildingImageStoreKey = "building-images"
	BuildingJobStoreKey   = "building-job"
	PushingJobStoreKey    = "pushing-job"
	KubernetesJobStoreKey = "kubernetes-job"
	RescaleJobStoreKey    = "rescale-job"
)

func getImages(key string) ([]*Image, error) {
	result := []*Image{}
	if err := GetValue(func(txn *badger.Txn) (err error) {
		result, err = getCachedImages(txn, key)
		return
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func setImage(key string, image *Image) error {
	return SetValue(func(txn *badger.Txn) error {
		images, err := getCachedImages(txn, key)
		if err != nil {
			return err
		}
		if inSlice(image.Tag, images) {
			return nil
		}
		images = append(images, image)
		bytes, err := json.Marshal(images)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), bytes)
	})
}

func removeImage(key, name string) error {
	return SetValue(func(txn *badger.Txn) error {
		images, err := getCachedImages(txn, key)
		if err != nil {
			return err
		}
		result := []*Image{}
		for _, image := range images {
			if !strings.EqualFold(name, image.Tag) {
				result = append(result, image)
			}
		}
		bytes, err := json.Marshal(result)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), bytes)
	})
}

func getCachedImages(txn *badger.Txn, key string) ([]*Image, error) {
	result := []*Image{}
	item, err := txn.Get([]byte(key))
	if err != nil {
		return result, nil // ignore
	}
	val, err := item.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(val, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func inSlice(candidate string, images []*Image) bool {
	for _, image := range images {
		if strings.EqualFold(candidate, image.Tag) {
			return true
		}
	}
	return false
}

func getJobs(key string) ([]*Job, error) {
	result := []*Job{}
	if err := GetValue(func(txn *badger.Txn) (err error) {
		result, err = getCachedJobs(txn, key)
		return
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func setJob(key string, job *Job) error {
	return SetValue(func(txn *badger.Txn) error {
		jobs, err := getCachedJobs(txn, key)
		if err != nil {
			return err
		}
		if idInSlice(job.ID, jobs) {
			return nil
		}
		jobs = append(jobs, job)
		bytes, err := json.Marshal(jobs)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), bytes)
	})
}

func removeJob(key, ID string) error {
	return SetValue(func(txn *badger.Txn) error {
		jobs, err := getCachedJobs(txn, key)
		if err != nil {
			return err
		}
		result := []*Job{}
		for _, job := range jobs {
			if !strings.EqualFold(ID, job.ID) {
				result = append(result, job)
			}
		}
		bytes, err := json.Marshal(result)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), bytes)
	})
}

func getCachedJobs(txn *badger.Txn, key string) ([]*Job, error) {
	result := []*Job{}
	item, err := txn.Get([]byte(key))
	if err != nil {
		return result, nil // ignore
	}
	val, err := item.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(val, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func idInSlice(candidate string, jobs []*Job) bool {
	for _, job := range jobs {
		if strings.EqualFold(candidate, job.ID) {
			return true
		}
	}
	return false
}
