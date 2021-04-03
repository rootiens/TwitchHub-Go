package services

type LiteratureResponse struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type CreateLitPayload struct {
	Title       string `json:"Title" validate:"required, min=3,max=150"`
	Description string `json:"Description validate:"required"`
}

type LiteratureCreateResponse struct {
	Literature *LiteratureResponse `json:"Literature"`
}
