package controllers

import (
	"TwitchHub/utils"

	"TwitchHub/models"

	"github.com/gofiber/fiber/v2"
)

type StartingController struct {
	model models.Starting
}

func (s StartingController) Index(ctx *fiber.Ctx) error {

	startings := s.model.All()
	return ctx.JSON(startings)
}

func (s StartingController) NewStarting(ctx *fiber.Ctx) error {

	newstarting := new(models.Starting)
	db := utils.DatabaseConnection()
	if err := ctx.BodyParser(newstarting); err != nil {
		return err
	}
	res := db.Create(&newstarting)
	return ctx.JSON(res)
}
