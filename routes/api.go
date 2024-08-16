package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	luhnController := controllers.NewLuhnController()
	facades.Route().Post("/credit_card_validation/json", luhnController.Json)
}
