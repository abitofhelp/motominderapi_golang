// Package request contains the request messages for the use cases.
package request

import (
	"github.com/go-ozzo/ozzo-validation"
)

// ListMotorcyclesRequest is a simple dto containing the required data for the ListMotorcyclesInteractor.
type ListMotorcyclesRequest struct {
}

// NewListMotorcyclesRequest creates a new instance of a ListMotorcyclesRequest.
// Returns (nil, error) when there is an error, otherwise (ListMotorcyclesRequest, nil).
func NewListMotorcyclesRequest() (*ListMotorcyclesRequest, error) {

	listRequest := &ListMotorcyclesRequest{}

	err := listRequest.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return listRequest, nil
}

// Validate verifies that a ListMotorcyclesRequest's fields contain valid data.
// Returns (an instance of ListMotorcyclesRequest, nil) on success, otherwise (nil, error)
func (request ListMotorcyclesRequest) Validate() error {
	return validation.ValidateStruct(&request)
}
