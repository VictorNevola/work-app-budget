package company

import (
	"github.com/VictorNevola/work-app-budget/internal/domain/address"
	"github.com/VictorNevola/work-app-budget/internal/domain/contact"
)

type (
	Entity struct {
		ID      int
		Name    string
		TaxID   string
		Address address.Entity
		Contact contact.Entity
	}
)
