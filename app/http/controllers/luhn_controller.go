package controllers

import (
	"unicode"

	"github.com/goravel/framework/contracts/http"
)

/*
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json'
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json' --header 'Content-Type: application/json' --data-raw '{"creditCardNumber": "3379 5135 6110 8795"}'  # True
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json' --header 'Content-Type: application/json' --data-raw '{"creditCardNumber": "3379 5135 6110 8794"}'  # False
curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json' --header 'Content-Type: application/json' --data-raw '{"creditCardNumber": "123"}'                  # False

curl --location --request POST 'http://127.0.0.1:3000/credit_card_validation/json' --header 'Content-Type: application/json' --data-raw '{"numbers": [{"creditCardNumber": "123"}, {"creditCardNumber": "3379 5135 6110 8795"}]}'  # "123": False, "3379 5135 6110 8795": True
*/

type CreditCardNumber struct {
	Number string `json:"creditCardNumber"`
}

type CreditCard struct {
	Numbers []CreditCardNumber `json:"numbers"`
}

type CreditCardResponse struct {
	CreditCardNumber string `json:"creditCardNumber"`
	IsValid          bool   `json:"isValid"`
}

type LuhnController struct{}

func NewLuhnController() *LuhnController {
	return &LuhnController{}
}

func (r *LuhnController) Json(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"numbers": "required",
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

	is_valid_results := []CreditCardResponse{}
	for _, creditCard := range creditCard.Numbers {
		num := creditCard.Number
		is_valid_results = append(is_valid_results, CreditCardResponse{
			CreditCardNumber: num,
			IsValid:          GetCardValidation(num),
		})
	}

	return ctx.Response().Success().Json(is_valid_results)
}

func GetCardValidation(number string) bool {
	var factor int = 2
	var sum int = 0
	var product int = 1
	var count int = 0
	for _, ch := range number {
		if unicode.IsDigit(ch) {
			count += 1
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

	if sum%10 == 0 && count == 16 {
		return true
	}

	return false
}
