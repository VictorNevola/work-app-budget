package quote

import (
	"time"

	"github.com/VictorNevola/work-app-budget/internal/domain/company"
	"github.com/VictorNevola/work-app-budget/internal/domain/customer"
	"github.com/VictorNevola/work-app-budget/internal/domain/material"
)

const (
	minTotal = 0
)

type (
	AreasValues struct {
		Description string
		AreaName    string
		UnitMeasure string
		Quantity    int64
		UnitValue   int64
		Total       int64
	}

	Totals struct {
		AreasTotais int64
		Discount    int64
		Total       int64
	}

	Entity struct {
		Description         string
		EstimatedInitDate   time.Time
		EstimatedFinishDate time.Time
		CreatedAt           time.Time
		UpdatedAt           time.Time
		Status              Status
		Areas               []AreasValues
		Discount            int64
		Totals              Totals
		Company             company.Entity
		Customer            customer.Entity
		Materials           []material.Entity
	}
)

func (e *Entity) CalcTotals() error {
	var areasTotal int64
	for _, area := range e.Areas {
		areasTotal += (area.Quantity * area.UnitValue)
	}

	e.Totals = Totals{
		Total:       areasTotal - e.Discount,
		Discount:    e.Discount,
		AreasTotais: areasTotal,
	}

	if e.Totals.Total < minTotal {
		return ErrDiscountGreaterThanTotal
	}

	return nil
}
