package db

import (
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/scaleshift/scaleshift/api/src/log"
)

// Error represents error information
type Error struct {
	gorm.Model
	Owner        string  `gorm:"index:owner;"`
	Caption      string  `gorm:"column:caption;"`
	ErrorMessage *string `gorm:"column:error_detail;"`
	OnlineAction *string `gorm:"column:online_action;"`
	ImageAction  *string `gorm:"column:image_action;"`
	ImageTag     *string `gorm:"column:image_tag;"`
	JobAction    *string `gorm:"column:job_action;"`
	JobID        *string `gorm:"column:job_id;"`
}

// OnlineAction defines the actions on the core business logic
type OnlineAction string

// OnlineAction
const (
	OnlineActionWrapJupyter  OnlineAction = "wrap-jupyternotebook"
	OnlineActionRunJupyter   OnlineAction = "run-jupyternotebook"
	OnlineActionSelectData   OnlineAction = "select-data"
	OnlineActionDeleteImages OnlineAction = "delete-images"
	OnlineActionUpdateStatus OnlineAction = "update-dbstatus"
	OnlineActionConvertToSig OnlineAction = "convert-singularity"
)
const (
	anonymous = "anonymous"
)

// FindErrors search errors
func FindErrors(owner string) ([]*Error, error) {
	result := []*Error{}
	owners := []string{anonymous, owner}
	if err := db.Where("owner IN (?)", owners).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// OnlineError persists an online error
func OnlineError(owner, caption string, onlineAction OnlineAction, err error) {
	message := ""
	if err != nil {
		message = err.Error()
	}
	db.Create(&Error{
		Owner:        owner,
		Caption:      caption,
		ErrorMessage: swag.String(message),
		OnlineAction: swag.String(string(onlineAction)),
	})
	log.Error(caption, err, nil)
}

// ImageError persists an image error
func ImageError(owner, caption string, imageAction ImageAction, tag string, err error) {
	message := ""
	if err != nil {
		message = err.Error()
	}
	db.Create(&Error{
		Owner:        owner,
		Caption:      caption,
		ErrorMessage: swag.String(message),
		ImageAction:  swag.String(string(imageAction)),
		ImageTag:     swag.String(tag),
	})
	log.Error(caption, err, nil)
}

// JobError persists a job error
func JobError(owner, caption string, jobAction JobAction, id string, err error) {
	message := ""
	if err != nil {
		message = err.Error()
	}
	db.Create(&Error{
		Owner:        owner,
		Caption:      caption,
		ErrorMessage: swag.String(message),
		JobAction:    swag.String(string(jobAction)),
		JobID:        swag.String(id),
	})
	log.Error(caption, err, nil)
}
