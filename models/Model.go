package models

import (
	"github.com/jinzhu/gorm"
	"morningo/config"
	"fmt"
)

var Model *gorm.DB

func init() {
	var err error
	log.Println(config.GetEnv().DATABASE.FormatDSN())
	Model, err = gorm.Open("mysql", config.GetEnv().DATABASE.FormatDSN())

	if err != nil {
		panic(err)
	}
}
