package quote

import (
	"context"
)

type (
	service struct {
		repo Repository
	}
)

func NewService(quoteRepository Repository) *service {
	return &service{
		repo: quoteRepository,
	}
}

func (s *service) Create(ctx context.Context, quoteDTO *CreateDTO) error {
	entity, err := quoteDTO.ToEntity()
	if err != nil {
		return ErrInternalServer
	}

	err = s.repo.Create(ctx, entity)
	if err != nil {
		return ErrInternalServer
	}

	return nil
}
