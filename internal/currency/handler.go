package currency

import (
	"currency-tracker/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	GetGoldRateByName(c *fiber.Ctx) error
}

type Handler struct {
	s *Service
}

func NewHandler() *Handler {
	return &Handler{s: NewService()}
}

func (h *Handler) GetGoldRateByName(c *fiber.Ctx) error {
	goldName := c.Params("name")
	recipe := h.s.GetGoldRateByName(goldName)
	return utils.DataResponseCreated(c, recipe)
}

func (h *Handler) GetCurrencyRateByName(c *fiber.Ctx) error {
	currencyName := c.Params("name")
	recipe := h.s.GetCurrencyRateByName(currencyName)
	return utils.DataResponseCreated(c, recipe)
}
