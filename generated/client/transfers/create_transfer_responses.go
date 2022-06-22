// Code generated by go-swagger; DO NOT EDIT.

package transfers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// CreateTransferReader is a Reader for the CreateTransfer structure.
type CreateTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTransferOK creates a CreateTransferOK with default headers values
func NewCreateTransferOK() *CreateTransferOK {
	return &CreateTransferOK{}
}

/* CreateTransferOK describes a response with status code 200, with default header values.

OK
*/
type CreateTransferOK struct {
	Payload *models.CreateTransferResponse
}

func (o *CreateTransferOK) Error() string {
	return fmt.Sprintf("[POST /v2/transfers][%d] createTransferOK  %+v", 200, o.Payload)
}
func (o *CreateTransferOK) GetPayload() *models.CreateTransferResponse {
	return o.Payload
}

func (o *CreateTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTransferResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
