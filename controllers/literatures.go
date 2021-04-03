package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rootiens/TwitchHub-Go/models"
	"github.com/rootiens/TwitchHub-Go/services"
	// "github.com/rootiens/TwitchHub-Go/utils"
)

func CreateLiterature(c *fiber.Ctx) error {
	b := new(services.CreateLitPayload)

	c.BodyParser(b)

	d := &models.Literature{
		Title:       b.Title,
		Description: b.Description,
	}

	if err := models.CreateLiterature(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&services.LiteratureCreateResponse{
		Literature: &services.LiteratureResponse{
			Title:       d.Title,
			Description: d.Description,
		},
	})
}
