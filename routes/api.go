package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	luhnController := controllers.NewLuhnController()
	facades.Route().Get("/credit_card_validation/{credit_card_number}", luhnController.Show)
}
