package models

import (
	"github.com/rootiens/TwitchHub-Go/utils"

	"gorm.io/gorm"
)

type Literature struct {
	gorm.Model
	Title       string
	Description string
}

func (l *Literature) GetAll() []Literature {
	var literatures []Literature

	db := utils.DatabaseConnection()

	// db.Raw("select * from literatures").Scan(&literatures)
	db.Find(&literatures)

	return literatures

}

func (l *Literature) Create() (int64, error) {

	db := utils.DatabaseConnection()

	res := db.Create(&l)

	return res.RowsAffected, res.Error

}
