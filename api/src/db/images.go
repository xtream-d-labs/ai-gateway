package db

import (
	"github.com/jinzhu/gorm"
)

// Image represents cached image information
type Image struct {
	gorm.Model
	Tag    string `gorm:"index:tag;"`
	Action string `gorm:"index:action;"`
	Owner  string
}

// ImageAction defines how handle a image
type ImageAction string

// ImageActions
const (
	ImageActionPulling  ImageAction = "pulling"
	ImageActionBuilding ImageAction = "building"
)

// Create persists its data to database
func (image *Image) Create() error {
	return db.Create(image).Error
}

// FindImages search images based on the ImageType
func FindImages(action ImageAction) ([]*Image, error) {
	result := []*Image{}
	if err := db.Where("action = ?", action).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// RemovePullingImage removes from PULLING cache
func RemovePullingImage(name string) error {
	return removeImage(name, ImageActionPulling)
}

// RemoveBuildingImage removes from BUILDING cache
func RemoveBuildingImage(name string) error {
	return removeImage(name, ImageActionBuilding)
}

func removeImage(tag string, action ImageAction) error {
	return db.Delete(Image{}, "tag = ? and action = ?", tag, action).Error
}
