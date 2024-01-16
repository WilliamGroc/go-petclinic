package models

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name string `json:"name"`
	Age  int `json:"age"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}