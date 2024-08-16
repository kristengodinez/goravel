package controllers

import "github.com/goravel/framework/contracts/http"

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

	return ctx.Response().Success().Json(http.Json{
		"creditCardNumber": creditCard.CreditCardNumber,
	})
}
