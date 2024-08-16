package controllers

import "github.com/goravel/framework/contracts/http"

type CreditCard struct {
	CreditCardNumber string `json:"CreditCardNumber" form:"CreditCardNumber"`
}

type LuhnController struct{}

func NewLuhnController() *LuhnController {
	return &LuhnController{}
}

func (r *LuhnController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
