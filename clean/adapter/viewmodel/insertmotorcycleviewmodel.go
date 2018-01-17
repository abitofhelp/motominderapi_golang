// Package viewmodel translates a response message into a view model.
package viewmodel

import (
	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// InsertMotorcycleViewModel translates a InsertMotorcycleResponse to a InsertMotorcycleViewModel.
// by the Configuration ring.
type InsertMotorcycleViewModel struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// NewInsertMotorcycleViewModel creates a new instance of a InsertMotorcycleViewModel.
// Returns an (instance of InsertMotorcycleViewModel, nil) on success, otherwise (nil, error)
func NewInsertMotorcycleViewModel(id int, message string, err error) (*InsertMotorcycleViewModel, error) {

	viewModel := &InsertMotorcycleViewModel{
		ID:      id,
		Message: message,
		Error:   err,
	}

	msgErr := viewModel.Validate()
	if msgErr != nil {
		// We had an error validating the response message,
		// so we will wrap the original error with the validation error.
		return nil, errors.Wrap(msgErr, viewModel.Error.Error())
	}

	// All okay
	return viewModel, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of InsertMotorcycleViewModel, nil) on success, otherwise (nil, error)
func (viewmodel *InsertMotorcycleViewModel) Handle(responseMessage *response.InsertMotorcycleResponse) (*InsertMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return NewInsertMotorcycleViewModel(constant.InvalidEntityID, responseMessage.Error.Error(), responseMessage.Error)
	}

	return NewInsertMotorcycleViewModel(responseMessage.ID, "Successfully inserted a new motorcycle.", nil)
}

// Validate verifies that a InsertMotorcycleViewModel's fields contain valid data.
// Returns (an instance of InsertMotorcycleViewModel, nil) on success, otherwise (nil, error).
func (viewmodel InsertMotorcycleViewModel) Validate() error {
	return validation.ValidateStruct(&viewmodel,
		// ID is required and it must be non-zero
		validation.Field(&viewmodel.ID, validation.Required, validation.Min(constant.MinEntityID)),
		// Message is required and it cannot be empty or nil.
		validation.Field(&viewmodel.Message, validation.Required, validation.NilOrNotEmpty),
	)
}