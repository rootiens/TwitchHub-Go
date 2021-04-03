package models

import (
	"github.com/rootiens/TwitchHub-Go/utils"
	"gorm.io/gorm"
)

type Literature struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func CreateLiterature(l *Literature) *gorm.DB {
	db := utils.DatabaseConnection()
	return db.Create(l)
}
