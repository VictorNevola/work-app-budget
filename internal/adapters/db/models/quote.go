package models

import (
	"time"

	domain "github.com/VictorNevola/work-app-budget/internal/domain/quote"
)

type (
	AreasValues struct {
		Description string `bson:"description"`
		AreaName    string `bson:"area_name"`
		UnitMeasure string `bson:"unit_measure"`
		Quantity    string `bson:"quantity"`
		Total       uint64 `bson:"total"`
	}

	Totals struct {
		AreasTotais float64 `bson:"areas_totais"`
		Discount    float64 `bson:"discount"`
	}

	Customer struct {
		Name    string  `bson:"name"`
		TaxID   string  `bson:"tax_id"`
		Contact Contact `bson:"contact"`
		Address Address `bson:"address"`
	}

	Material struct {
		Name        string `bson:"name"`
		UnitMeasure string `bson:"unit_measure"`
		Quantity    string `bson:"quantity"`
	}

	Company struct {
		ID      int     `bson:"company_id"`
		Name    string  `bson:"name"`
		TaxID   string  `bson:"tax_id"`
		Address Address `bson:"address"`
		Contact Contact `bson:"contact"`
	}

	Contact struct {
		Email string `bson:"email"`
		Phone string `bson:"phone"`
	}

	Address struct {
		Street       string `bson:"street"`
		Number       string `bson:"number"`
		Complement   string `bson:"complement"`
		Neighborhood string `bson:"neighborhood"`
		City         string `bson:"city"`
		State        string `bson:"state"`
		PostalCode   string `bson:"postal_code"`
		Country      string `bson:"country"`
	}

	Quote struct {
		BaseModel
		Description         string        `bson:"description"`
		EstimatedInitDate   time.Time     `bson:"estimated_init_date"`
		EstimatedFinishDate time.Time     `bson:"estimated_finish_date"`
		Status              domain.Status `bson:"status"`
		Areas               []AreasValues `bson:"areas"`
		Materials           []Material    `bson:"materials"`
		Totals              Totals        `bson:"totals"`
		Company             Company       `bson:"company"`
		Customer            Customer      `bson:"customer"`
	}
)

func (q *Quote) FromEntity(entity *domain.Entity) *Quote {
	return &Quote{
		Description:         entity.Description,
		EstimatedInitDate:   entity.EstimatedInitDate,
		EstimatedFinishDate: entity.EstimatedFinishDate,
		Status:              entity.Status,
		Totals:              Totals(entity.Totals),
	}
}
