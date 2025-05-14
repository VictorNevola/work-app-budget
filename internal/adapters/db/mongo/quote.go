package mongo

import (
	"context"

	"github.com/VictorNevola/work-app-budget/internal/adapters/db/models"
	"github.com/VictorNevola/work-app-budget/internal/domain/quote"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuoteRepository struct {
	dbCollection *mongo.Collection
}

var _ quote.Repository = (*QuoteRepository)(nil)

func NewQuoteRepository(db *Connection) *QuoteRepository {
	return &QuoteRepository{
		dbCollection: db.Collection("Quotes"),
	}
}

func (p *QuoteRepository) Create(ctx context.Context, quote *quote.Entity) error {
	model := new(models.Quote).FromEntity(quote)
	model.NewID()
	model.SetCreatedAt()

	return nil
}
