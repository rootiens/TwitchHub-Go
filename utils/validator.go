package utils

// import "github.com/go-playground/validator"

// "fmt"
// "strings"

// "github.com/go-playground/validator/v10"
// "github.com/gofiber/fiber/v2"

// var validate = validator.New()

// type ErrorResponse struct {
// 	FailedField string
// 	Tag         string
// 	Value       string
// }

// func Validate(payload interface{}) []*ErrorResponse {
// 	var errors []*ErrorResponse
// 	validate := validator.New()
// 	err := validate.Struct(payload)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var element ErrorResponse
// 			element.FailedField = err.StructNamespace()
// 			element.Tag = err.Tag()
// 			element.Value = err.Param()
// 			errors = append(errors, &element)
// 		}
// 	}
// 	return errors
// }

// // Validate validates the input struct
// func Validate(payload interface{}) *fiber.Error {
// 	err := validate.Struct(payload)

// 	if err != nil {
// 		var errors []string
// 		for _, err := range err.(validator.ValidationErrors) {
// 			errors = append(
// 				errors,
// 				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
// 			)
// 		}

// 		return &fiber.Error{
// 			Code:    fiber.StatusBadRequest,
// 			Message: strings.Join(errors, ","),
// 		}
// 	}

// 	return nil
// }
