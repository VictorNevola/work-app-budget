package quote

import (
	"time"

	"github.com/VictorNevola/work-app-budget/internal/domain/company"
	"github.com/VictorNevola/work-app-budget/internal/domain/customer"
	"github.com/VictorNevola/work-app-budget/internal/domain/material"
)

type (
	AreasValues struct {
		Description string
		AreaName    string
		UnitMeasure string
		Quantity    string
		Total       uint64
	}

	Totals struct {
		AreasTotais float64
		Discount    float64
	}

	Entity struct {
		Description         string
		EstimatedInitDate   time.Time
		EstimatedFinishDate time.Time
		CreatedAt           time.Time
		UpdatedAt           time.Time
		Status              Status
		Areas               []AreasValues
		Totals              Totals
		Company             company.Entity
		Customer            customer.Entity
		Materials           []material.Entity
	}
)
