package quote

import (
	"github.com/VictorNevola/work-app-budget/internal/domain/address"
	"github.com/VictorNevola/work-app-budget/internal/domain/company"
	"github.com/VictorNevola/work-app-budget/internal/domain/contact"
	"github.com/VictorNevola/work-app-budget/internal/domain/customer"
	"github.com/VictorNevola/work-app-budget/internal/domain/material"
	"github.com/VictorNevola/work-app-budget/pkg/date"
)

type (
	CreateDTO struct {
		Description         string           `json:"description" validate:"required,min=5,max=255"`
		EstimatedInitDate   string           `json:"estimated_init_date" validate:"required,datetime=2006-01-02"`
		EstimatedFinishDate string           `json:"estimated_finish_date" validate:"required,datetime=2006-01-02"`
		Company             CompanyDTO       `json:"company" validate:"required"`
		Customer            CustomerDTO      `json:"customer" validate:"required"`
		Discount            uint64           `json:"discount" validate:"min=0"`
		Areas               []AreasValuesDTO `json:"areas" validate:"dive"`
		Materials           []MaterialDTO    `json:"materials" validate:"dive"`
	}

	CompanyDTO struct {
		CompanyID int        `json:"company_id" validate:"required,min=1"`
		Name      string     `json:"name" validate:"required,min=2"`
		TaxID     string     `json:"tax_id" validate:"required"`
		Address   AddressDTO `json:"address" validate:"required"`
		Contact   ContactDTO `json:"contact" validate:"required"`
	}

	CustomerDTO struct {
		Name    string     `json:"name" validate:"required,min=2"`
		TaxID   string     `json:"tax_id" validate:"required"`
		Address AddressDTO `json:"address" validate:"required"`
		Contact ContactDTO `json:"contact" validate:"required"`
	}

	AddressDTO struct {
		Street       string `json:"street" validate:"required,min=3"`
		Number       string `json:"number" validate:"required"`
		Complement   string `json:"complement"`
		Neighborhood string `json:"neighborhood" validate:"required"`
		City         string `json:"city" validate:"required"`
		State        string `json:"state" validate:"required"`
		PostalCode   string `json:"postal_code" validate:"required"`
		Country      string `json:"country" validate:"required"`
	}

	ContactDTO struct {
		Email string `json:"email" validate:"required,email"`
		Phone string `json:"phone" validate:"required,e164"`
	}

	AreasValuesDTO struct {
		Description string `json:"description" validate:"required,min=5,max=100"`
		AreaName    string `json:"area_name" validate:"required,min=2,max=50"`
		UnitMeasure string `json:"unit_measure" validate:"required,min=1,max=10"`
		Quantity    uint64 `json:"quantity" validate:"required"`
		UnitValue   uint64 `json:"unit_value" validate:"required,min=0"`
	}

	MaterialDTO struct {
		Name        string `json:"name" validate:"required,min=2,max=100"`
		UnitMeasure string `json:"unit_measure" validate:"required,min=1,max=10"`
		Quantity    uint64 `json:"quantity" validate:"required"`
	}
)

func (dto *CompanyDTO) ToEntity() company.Entity {
	return company.Entity{
		CompanyID: dto.CompanyID,
		Name:      dto.Name,
		TaxID:     dto.TaxID,
		Address: address.Entity{
			Street:       dto.Address.Street,
			Number:       dto.Address.Number,
			Complement:   dto.Address.Complement,
			Neighborhood: dto.Address.Neighborhood,
			City:         dto.Address.City,
			State:        dto.Address.State,
			PostalCode:   dto.Address.PostalCode,
			Country:      dto.Address.Country,
		},
		Contact: contact.Entity{
			Email: dto.Contact.Email,
			Phone: dto.Contact.Phone,
		},
	}
}

func (dto *CustomerDTO) ToEntity() customer.Entity {
	return customer.Entity{
		Name:  dto.Name,
		TaxID: dto.TaxID,
		Address: address.Entity{
			Street:       dto.Address.Street,
			Number:       dto.Address.Number,
			Complement:   dto.Address.Complement,
			Neighborhood: dto.Address.Neighborhood,
			City:         dto.Address.City,
			State:        dto.Address.State,
			PostalCode:   dto.Address.PostalCode,
			Country:      dto.Address.Country,
		},
		Contact: contact.Entity{
			Email: dto.Contact.Email,
			Phone: dto.Contact.Phone,
		},
	}
}

func (dto *CreateDTO) ToEntity() (*Entity, error) {
	initDate, err := date.FromString(dto.EstimatedInitDate)
	if err != nil {
		return nil, ErrInternalServer
	}

	finishDate, err := date.FromString(dto.EstimatedFinishDate)
	if err != nil {
		return nil, ErrInternalServer
	}

	areas := make([]AreasValues, len(dto.Areas))
	for index, a := range dto.Areas {
		areas[index] = AreasValues{
			Description: a.Description,
			AreaName:    a.AreaName,
			UnitMeasure: a.UnitMeasure,
			Quantity:    a.Quantity,
			UnitValue:   a.UnitValue,
		}
	}

	materials := make([]material.Entity, len(dto.Materials))
	for index, m := range dto.Materials {
		materials[index] = material.Entity{
			Name:        m.Name,
			UnitMeasure: m.UnitMeasure,
			Quantity:    m.Quantity,
		}
	}

	companyEntity := dto.Company.ToEntity()
	customerEntity := dto.Customer.ToEntity()

	entity := &Entity{
		Description:         dto.Description,
		EstimatedInitDate:   *initDate,
		EstimatedFinishDate: *finishDate,
		Status:              Pending,
		Areas:               areas,
		Company:             companyEntity,
		Customer:            customerEntity,
		Materials:           materials,
	}

	return entity, nil
}
