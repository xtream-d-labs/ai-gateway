package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
)

// Image state
const (
	StableImage   = "stable"
	PullingImage  = "pulling"
	BuildingImage = "building"
	Singularity   = "singularity"
	RescaleSend   = "rescale-send"
	RescaleStart  = "rescale-start"
)

// Image represents cached image information
type Image struct {
	Tag     string
	Status  string
	Started time.Time
}

// Job represents cached job information
type Job struct {
	ID          string
	Status      string
	DockerImage string
	PythonFile  string
	Workspaces  []string
	Commands    []string
	CoreType    string
	Cores       int64
	Started     time.Time
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

// GetJobs returns jobs under BUILDING Singularity images / Rescale Job state
func GetJobs() ([]*Job, error) {
	jobs, err := getJobs(SingularityStoreKey)
	if err != nil {
		return nil, err
	}
	additionals, err := getJobs(RescaleJobStoreKey)
	if err != nil {
		return nil, err
	}
	return append(jobs, additionals...), nil
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

// SetSingularityJobMeta add a task to BUILD Singularity image
func SetSingularityJobMeta(job *Job) error {
	if job.Status != Singularity {
		return fmt.Errorf("%s is not Singularity job", job.ID)
	}
	return setJob(SingularityStoreKey, job)
}

// RemoveSingularityJobs removes job
func RemoveSingularityJobs(ID string) error {
	return removeJob(SingularityStoreKey, ID)
}

// UpdateJob update job status
func UpdateJob(ID, from, to, status string) error {
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

// RemoveRescaleJob removes a job you specified
func RemoveRescaleJob(ID string) error {
	return removeJob(RescaleJobStoreKey, ID)
}

// Cache keys
const (
	PullingImageStoreKey  = "pulling-images"
	BuildingImageStoreKey = "building-images"
	SingularityStoreKey   = "singularity"
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
