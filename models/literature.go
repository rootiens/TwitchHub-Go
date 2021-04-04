package models

import (
	"fmt"

	"github.com/rootiens/TwitchHub-Go/utils"

	"gorm.io/gorm"
)

type Literature struct {
	gorm.Model
	Title       string
	Description string
}

func (l *Literature) GetAllLiteratures() []Literature {
	var literatures []Literature

	db := utils.DatabaseConnection()

	db.Find(&literatures)

	return literatures

}

func (l *Literature) CreateLiterature() (int64, error) {

	db := utils.DatabaseConnection()

	res := db.Create(&l)

	return res.RowsAffected, res.Error

}

func (l *Literature) GetLiterature() Literature {
	var literature Literature
	fmt.Println(l)

	db := utils.DatabaseConnection()
	db.Where("id = ?", l).First(&literature)
	return literature
}
