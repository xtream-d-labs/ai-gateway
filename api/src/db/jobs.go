package db

import (
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/models"
)

// Job represents cached job information
type Job struct {
	gorm.Model
	JobID       string `gorm:"column:job_id;index:jobid;"`
	Action      string `gorm:"index:action;"`
	Platform    int
	Status      string
	DockerImage string `gorm:"column:docker_image;"`
	PythonFile  string `gorm:"column:python_file;"`
	Workspaces  string
	Commands    string
	CPU         int64
	Memory      int64
	GPU         int64
	CoreType    string `gorm:"column:core_type;"`
	Cores       int64
	TargetID    string `gorm:"column:target_id;"` // Job ID of the target platform
}

// JobAction defines how to handle the job
type JobAction string

// JobActions
const (
	JobActionBuilding   JobAction = "building"
	JobActionPushing    JobAction = "pushing"
	JobActionKubernetes JobAction = "kubernetes"
	JobActionRescale    JobAction = "rescale"
)

// PlatformType defines in where the job will be run
type PlatformType int

// PlatformTypes
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

// JobStatus defines the status of a job
type JobStatus string

// JobStatuses
const (
	JobStatusImageBuilding  JobStatus = "building"
	JobStatusImagePushing   JobStatus = "pushing"
	JobStatusK8sInit        JobStatus = "k8s-job-init" // Kubernetes
	JobStatusK8sStarted     JobStatus = "k8s-job-started"
	JobStatusK8sPending     JobStatus = "k8s-job-pending"
	JobStatusK8sRunning     JobStatus = "k8s-job-runnning"
	JobStatusK8sSucceeded   JobStatus = "k8s-job-succeeded"
	JobStatusK8sFailed      JobStatus = "k8s-job-failed"
	JobStatusRescaleSend    JobStatus = "rescale-send" // Rescale
	JobStatusRescaleStarted JobStatus = "rescale-start"
	JobStatusRescaleRunning JobStatus = "rescale-runnning"
	JobStatusRescaleSucceed JobStatus = "rescale-succeeded"
	JobStatusRescaleFailed  JobStatus = "rescale-failed"
	JobStatusUnknown        JobStatus = "unknown"
)

// Value casts its value to the pointer
func (s JobStatus) Value() *string {
	return swag.String(string(s))
}

// Create persists its data to database
func (job *Job) Create() error {
	return db.Create(job).Error
}

// GetJobs returns jobs under BUILDING job images / Rescale status
func GetJobs() ([]*Job, error) {
	result := []*Job{}
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetJob returns job specified
func GetJob(jobID string) (*Job, error) {
	result := &Job{}
	if err := db.Where("job_id = ?", jobID).Find(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveBuildingJobImagesJobs removes job
func RemoveBuildingJobImagesJobs(id string) error {
	return removeJob(id, JobActionBuilding)
}

// RemovePushingJobImageJobs removes job
func RemovePushingJobImageJobs(id string) error {
	return removeJob(id, JobActionPushing)
}

func removeJob(id string, action JobAction) error {
	return db.Delete(Job{}, "job_id = ? and action = ?", id, action).Error
}

// RemoveJob removes a job you specified
func RemoveJob(id string) error {
	return db.Delete(Job{}, "job_id = ?", id).Error
}

// UpdateJob update job status
func UpdateJob(id string, from, to JobAction, status JobStatus, targetID *string) error {
	fields := map[string]interface{}{
		"action": to,
		"status": status,
	}
	if targetID == nil {
		fields["target_id"] = swag.StringValue(targetID)
	}
	return db.Model(&Job{}).Where("job_id = ? and action = ?", id, from).Updates(fields).Error
}
