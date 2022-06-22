// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// CancelOrderReader is a Reader for the CancelOrder structure.
type CancelOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelOrderOK creates a CancelOrderOK with default headers values
func NewCancelOrderOK() *CancelOrderOK {
	return &CancelOrderOK{}
}

/* CancelOrderOK describes a response with status code 200, with default header values.

OK
*/
type CancelOrderOK struct {
	Payload *models.CancelOrderResponse
}

func (o *CancelOrderOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/orders/{id}][%d] cancelOrderOK  %+v", 200, o.Payload)
}
func (o *CancelOrderOK) GetPayload() *models.CancelOrderResponse {
	return o.Payload
}

func (o *CancelOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
