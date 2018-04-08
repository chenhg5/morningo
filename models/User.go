package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string
	Avatar string
	Sex    int
}

func (User) TableName() string {
	return "users"
}
