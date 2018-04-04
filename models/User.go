package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID     uint `gorm:"primary_key"`
	Name   string
	Avatar string
	Sex    int
}

func (User) TableName() string {
	return "users"
}