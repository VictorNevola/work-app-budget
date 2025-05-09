package customer

import (
	"github.com/VictorNevola/work-app-budget/internal/domain/address"
	"github.com/VictorNevola/work-app-budget/internal/domain/contact"
)

type (
	Entity struct {
		Name    string
		TaxID   string
		Contact contact.Entity
		Address address.Entity
	}
)
