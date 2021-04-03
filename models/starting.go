package models

import (
	"TwitchHub/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Starting struct {
	gorm.Model
	Description string `json:Description form:description`
}

func (s *Starting) All() []Starting {
	var startings []Starting
	db := utils.DatabaseConnection()
	db.Raw("select * from startings").Scan(&startings)
	return startings
}

func (s *Starting) NewStarting(c *fiber.Ctx) Starting {
	var starting Starting
	if err := c.BodyParser(starting); err != nil {
		return starting
	}
	return starting
}
