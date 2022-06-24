// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"immutable.com/imx-core-sdk-golang/generated/models"
)

// ListCollectionFiltersReader is a Reader for the ListCollectionFilters structure.
type ListCollectionFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCollectionFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCollectionFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCollectionFiltersOK creates a ListCollectionFiltersOK with default headers values
func NewListCollectionFiltersOK() *ListCollectionFiltersOK {
	return &ListCollectionFiltersOK{}
}

/* ListCollectionFiltersOK describes a response with status code 200, with default header values.

OK
*/
type ListCollectionFiltersOK struct {
	Payload *models.CollectionFilter
}

func (o *ListCollectionFiltersOK) Error() string {
	return fmt.Sprintf("[GET /v1/collections/{address}/filters][%d] listCollectionFiltersOK  %+v", 200, o.Payload)
}
func (o *ListCollectionFiltersOK) GetPayload() *models.CollectionFilter {
	return o.Payload
}

func (o *ListCollectionFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectionFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}