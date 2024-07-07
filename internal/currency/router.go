package currency

import (
	"github.com/gofiber/fiber/v2"
)

func RouterCurrency(router fiber.Router) {
	var currencyHandler = NewHandler()
	router.Get("/gold/:name", currencyHandler.GetGoldRateByName)
	router.Get("/:name", currencyHandler.GetCurrencyRateByName)
}
