package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	luhnController := controllers.NewLuhnController()
	facades.Route().Post("/credit_card_validation/json", luhnController.Json)
}
