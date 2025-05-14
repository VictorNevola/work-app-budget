package quote

import "context"

type (
	Repository interface {
		Create(ctx context.Context, quote *Entity) error
	}

	Service interface {
		Create(ctx context.Context, quoteDTO *CreateDTO) error
	}
)
