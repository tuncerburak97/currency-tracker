package app

import (
	"currency-tracker/internal/config"
	"currency-tracker/internal/currency"
	"currency-tracker/internal/database"
	"currency-tracker/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Config *config.Config
	Fiber  *fiber.App
}

func NewApp() (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	database.InitInfluxDatabase()
	go currency.InitScheduler()
	app := fiber.New()

	return &App{
		Config: cfg,
		Fiber:  app,
	}, nil
}

func (a *App) Start() {
	a.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	RegisterRoutes(a.Fiber)

	a.Fiber.Static("/ui", "./public")

	err := a.Fiber.Listen(":" + "8080")
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("Starting server on port 8080")

}

func RegisterRoutes(app *fiber.App) {

	// currency
	currencyGroup := app.Group("/currency")
	currency.RouterCurrency(currencyGroup)

}
