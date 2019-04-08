package bindings

import "errors"

// Validatable - Interface for Validatable
type Validatable interface {
	Validate() error
}

// ErrNotValidatable - Error for Validatable
var ErrNotValidatable = errors.New("Type is not validatable")

// Validator - Struct of Validator
type Validator struct{}

// Validate - If 'i' has Validate function implemented, execute it
func (v *Validator) Validate(i interface{}) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return ErrNotValidatable
}
