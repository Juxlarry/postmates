package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Status    string      `json:"status"`
	Phone     string      `json:"phone"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Location  GeoLocation `json:"location"`
}
