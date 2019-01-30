package models

import (
	u "API/utils"
	"errors"

	"github.com/jinzhu/gorm"
)

type Gallery struct {
	gorm.Model
	Slug   string `json:"slug"`
	Event  string `json:"event";sql:"-"`
	Year   int    `json:"year"`
	Images []File `gorm:"foreignkey:GalleryID"`
}

func (gallery *Gallery) Validate() (map[string]interface{}, bool) {

	if gallery.Year < 2000 {
		return u.Message(false, "Year is required"), false
	}

	temp := &Gallery{}

	err := GetDB().Table("galleries").Where("year = ?", gallery.Year).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (gallery *Gallery) Create() map[string]interface{} {

	if resp, ok := gallery.Validate(); !ok {
		return resp
	}

	GetDB().Create(gallery)

	if gallery.ID <= 0 {
		return u.Message(false, "Failed to create gallery, connection error.")
	}

	response := u.Message(true, "Gallery has been created")
	response["gallery"] = gallery
	return response
}

func GetFirstGallery() map[string]interface{} {

	gal := &Gallery{}
	GetDB().Table("galleries").Where("slug = ?", "test").First(gal)

	response := u.Message(true, "Gallery found")
	response["gallery"] = gal
	return response
}

func GetGalleryBySlug(slug string) (*Gallery, error) {

	gal := &Gallery{}
	GetDB().Table("galleries").Where("slug = ?", slug).First(&gal)

	if gal.Year == 0 { //Gallery not found!
		return gal, errors.New("Gallery not found")
	}
	return gal, nil
}

func GetGalleryById(id uint) (*Gallery, error) {

	gal := &Gallery{}
	GetDB().Table("galleries").Where("id = ?", id).First(&gal)

	if gal.Year == 0 { //Gallery not found!
		return gal, errors.New("Gallery not found")
	}
	return gal, nil
}

func (gallery *Gallery) AddFile(file *File) {
	GetDB().Model(gallery).Association("Images").Append(file)
}

func GetImagesFromFirstGallery() map[string]interface{} {
	files := []File{File{}}
	gal := &Gallery{}
	GetDB().Table("galleries").First(&gal)
	GetDB().Table("files").Where("gallery_id = ?", gal.ID).Find(&files)

	response := u.Message(true, "Files found")
	gal.Images = files
	response["gallery"] = gal
	return response
}

func GetImagesFromGalleryBySlug(slug string) map[string]interface{} {
	files := []File{File{}}
	gal := &Gallery{}
	GetDB().Table("galleries").Where("slug = ?", slug).First(&gal)
	GetDB().Table("files").Where("gallery_id = ?", gal.ID).Find(&files)

	response := u.Message(true, "Files found")
	gal.Images = files
	response["gallery"] = gal
	return response
}
