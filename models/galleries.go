package models

import (
	u "API/utils"

	"github.com/jinzhu/gorm"
)

type Gallery struct {
	gorm.Model
	Event  string `json:"event";sql:"-"`
	Year   int    `json:"year"`
	Images []File `json:"images";sql:"-"`
}

func (gallery *Gallery) Validate() (map[string]interface{}, bool) {

	if gallery.Year < 2000 {
		return u.Message(false, "Year is required"), false
	}

	//Email must be unique
	temp := &Gallery{}

	//check for errors and duplicate emails
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

func GetGallery(u uint) *Gallery {

	gal := &Gallery{}
	GetDB().Table("accounts").Where("id = ?", u).First(gal)
	if gal.Year == 0 { //Gallery not found!
		return nil
	}
	return gal
}

// func UpdateGallery(gallery *Gallery) {
// 	gallery.UpdatedAt = time.Now()
// 	stmt, err :=

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = stmt.Exec(car.Manufacturer, car.Design, car.Style, car.Doors, car.UpdatedAt, car.Id)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
