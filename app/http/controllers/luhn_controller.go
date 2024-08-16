package controllers

import "github.com/goravel/framework/contracts/http"

type CreditCard struct {
	CreditCardNumber string `json:"CreditCardNumber" form:"CreditCardNumber"`
}

type LuhnController struct{}

func NewLuhnController() *LuhnController {
	return &LuhnController{}
}

func (r *LuhnController) Json(ctx http.Context) http.Response {
	ctx.Request().Validate(map[string]string{
		"creditCardNumber": "required",
	})
	return ctx.Response().Success().Json(http.Json{
		"creditCardNumber": "123",
	})
}
