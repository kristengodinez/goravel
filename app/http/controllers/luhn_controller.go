package controllers

import "github.com/goravel/framework/contracts/http"

type LuhnController struct{}

func NewLuhnController() *LuhnController {
	return &LuhnController{}
}

func (r *LuhnController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
