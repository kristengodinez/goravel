package controllers

import (
	"unicode"

	"github.com/goravel/framework/contracts/http"
)

/*
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json'
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json' --header 'Content-Type: application/json' --data-raw '{"creditCardNumber": "123"}'
*/

type CreditCard struct {
	CreditCardNumber string `json:"CreditCardNumber" form:"CreditCardNumber"`
}

type LuhnController struct{}

func NewLuhnController() *LuhnController {
	return &LuhnController{}
}

func (r *LuhnController) Json(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"creditCardNumber": "required",
	})
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}
	if validator.Fails() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": validator.Errors().All(),
		})
	}

	var creditCard CreditCard
	if err := validator.Bind(&creditCard); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}

	GetCardValidation(creditCard.CreditCardNumber)

	return ctx.Response().Success().Json(http.Json{
		"creditCardNumber": creditCard.CreditCardNumber,
		"isValid":          GetCardValidation(creditCard.CreditCardNumber),
	})
}

func GetCardValidation(number string) bool {
	var factor int = 2
	var sum int = 0
	var product int = 1
	for _, ch := range number {
		if unicode.IsDigit(ch) {
			product = (int(ch) - '0') * factor

			if product >= 10 {
				sum = sum + 1 + product%10
			} else {
				sum = sum + product
			}

			if factor == 2 {
				factor = 1
			} else {
				factor = 2
			}
		}
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
