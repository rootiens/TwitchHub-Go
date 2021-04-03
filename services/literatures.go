package services

import (
	"TwitchHub/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type LiteratureResponse struct {
	Message string `json:"message"`
	// Title       string `json:"title"`
	// Description string `json:"description"`
	Data models.Literature
}

type LiteraturePayload struct {
	Title       string `validate="required,min=5,max=150`
	Description string `validate="required,min=5,max=150`
}

func (l LiteraturePayload) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required, validation.Length(5, 20)),
		validation.Field(&l.Description, validation.Required, validation.Length(5, 50)),
	)
}
