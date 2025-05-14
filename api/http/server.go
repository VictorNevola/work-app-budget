package http

import (
	httpQuote "github.com/VictorNevola/work-app-budget/api/http/quote"
	"github.com/VictorNevola/work-app-budget/internal/adapters/db/mongo"
	domainQuote "github.com/VictorNevola/work-app-budget/internal/domain/quote"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewFiberServer(
	database *mongo.Connection,
	validator *validator.Validate,
) {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	// repositories
	quotesRepositories := mongo.NewQuoteRepository(database)

	// services
	quotesServices := domainQuote.NewService(quotesRepositories)

	// handlers
	httpQuote.NewQuoteHandler(app, quotesServices, validator)

	app.Listen(":3000")
}
