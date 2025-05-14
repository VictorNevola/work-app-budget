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
		Quantity    uint64 `bson:"quantity"`
		Total       uint64 `bson:"total"`
	}

	Totals struct {
		AreasTotais uint64 `bson:"areas_totais"`
		Discount    uint64 `bson:"discount"`
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
		Quantity    uint64 `bson:"quantity"`
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
		BaseModel           `bson:",inline"`
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
	quote := &Quote{
		Description:         entity.Description,
		EstimatedInitDate:   entity.EstimatedInitDate,
		EstimatedFinishDate: entity.EstimatedFinishDate,
		Status:              entity.Status,
	}

	quote.Company = Company{
		ID:    entity.Company.CompanyID,
		Name:  entity.Company.Name,
		TaxID: entity.Company.TaxID,
		Address: Address{
			Street:       entity.Company.Address.Street,
			Number:       entity.Company.Address.Number,
			Complement:   entity.Company.Address.Complement,
			Neighborhood: entity.Company.Address.Neighborhood,
			City:         entity.Company.Address.City,
			State:        entity.Company.Address.State,
			PostalCode:   entity.Company.Address.PostalCode,
			Country:      entity.Company.Address.Country,
		},
		Contact: Contact{
			Email: entity.Company.Contact.Email,
			Phone: entity.Company.Contact.Phone,
		},
	}

	quote.Customer = Customer{
		Name:  entity.Customer.Name,
		TaxID: entity.Customer.TaxID,
		Address: Address{
			Street:       entity.Customer.Address.Street,
			Number:       entity.Customer.Address.Number,
			Complement:   entity.Customer.Address.Complement,
			Neighborhood: entity.Customer.Address.Neighborhood,
			City:         entity.Customer.Address.City,
			State:        entity.Customer.Address.State,
			PostalCode:   entity.Customer.Address.PostalCode,
			Country:      entity.Customer.Address.Country,
		},
		Contact: Contact{
			Email: entity.Customer.Contact.Email,
			Phone: entity.Customer.Contact.Phone,
		},
	}

	for _, area := range entity.Areas {
		quote.Areas = append(quote.Areas, AreasValues{
			Description: area.Description,
			AreaName:    area.AreaName,
			UnitMeasure: area.UnitMeasure,
			Quantity:    area.Quantity,
			Total:       area.Total,
		})
	}

	for _, material := range entity.Materials {
		quote.Materials = append(quote.Materials, Material{
			Name:        material.Name,
			UnitMeasure: material.UnitMeasure,
			Quantity:    material.Quantity,
		})
	}

	quote.Totals = Totals{
		AreasTotais: entity.Totals.AreasTotais,
		Discount:    entity.Totals.Discount,
	}

	return quote
}
