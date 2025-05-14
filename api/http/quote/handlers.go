package quote

import (
	"net/http"

	"github.com/VictorNevola/work-app-budget/api/http/helpers"
	domainQuote "github.com/VictorNevola/work-app-budget/internal/domain/quote"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	Handler struct {
		service   domainQuote.Service
		validator *validator.Validate
	}
)

func NewQuoteHandler(
	app *fiber.App,
	service domainQuote.Service,
	validator *validator.Validate,
) {
	hander := Handler{
		service,
		validator,
	}

	app.Route("/quote", func(r fiber.Router) {
		r.Post("/", hander.CreateQuote)
	})
}

func (h *Handler) CreateQuote(c *fiber.Ctx) error {
	var QuoteDTO domainQuote.CreateDTO
	if err := c.BodyParser(&QuoteDTO); err != nil {
		return helpers.HandlerError(c, http.StatusBadGateway, err)
	}

	if err := validator.New().Struct(QuoteDTO); err != nil {
		return helpers.HandlerError(c, http.StatusBadGateway, err)
	}

	err := h.service.Create(c.Context(), &QuoteDTO)
	if err != nil {
		return helpers.HandlerError(c, http.StatusInternalServerError, err)
	}

	return c.SendStatus(http.StatusOK)
}
