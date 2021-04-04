package controllers

import (
	"strconv"

	"github.com/rootiens/TwitchHub-Go/models"
	"github.com/rootiens/TwitchHub-Go/services"

	// "log"

	"github.com/gofiber/fiber/v2"
)

type LiteratureController struct {
	model models.Literature
}

func (l LiteratureController) Index(ctx *fiber.Ctx) error {
	literatures := l.model.GetAllLiteratures()
	return ctx.JSON(literatures)
}

func (l LiteratureController) Create(ctx *fiber.Ctx) error {
	b := new(services.LiteraturePayload)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}
	validationErrors := b.Validate()

	if validationErrors != nil {
		return ctx.JSON(validationErrors)
	}

	l.model = models.Literature{
		Title:       b.Title,
		Description: b.Description,
	}

	_, modelError := l.model.CreateLiterature()

	if modelError != nil {
		fiber.NewError(fiber.StatusConflict, modelError.Error())
	}

	return ctx.JSON(services.LiteratureResponse{
		Message: "success",
		Data:    l.model,
	})

}

func (l LiteratureController) GetLiterature(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("ID"))

	if err != nil || id < 0 {
		return ctx.SendStatus(404)
	}
	l.model.ID = uint(id)

	l.model.GetLiterature()

	return ctx.JSON(services.LiteratureResponse{
		Message: "success",
		Data:    l.model,
	})
}
