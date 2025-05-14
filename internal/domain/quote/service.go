package quote

import "context"

type (
	CreateDTO struct {
		Description         string `json:"description" validate:"required,min=5,max=255"`
		EstimatedInitDate   string `json:"estimated_init_date" validate:"required,datetime=2006-01-02"`
		EstimatedFinishDate string `json:"estimated_finish_date" validate:"required,datetime=2006-01-02"`
	}

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
	return nil
}
