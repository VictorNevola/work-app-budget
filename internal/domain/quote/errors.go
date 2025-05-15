package quote

import "errors"

var (
	ErrNotFound                 = errors.New("not Found")
	ErrInternalServer           = errors.New("internel server error")
	ErrDiscountGreaterThanTotal = errors.New("the discount cannot be greater than the budget value")
)
