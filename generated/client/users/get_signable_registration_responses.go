// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// GetSignableRegistrationReader is a Reader for the GetSignableRegistration structure.
type GetSignableRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignableRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignableRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSignableRegistrationOK creates a GetSignableRegistrationOK with default headers values
func NewGetSignableRegistrationOK() *GetSignableRegistrationOK {
	return &GetSignableRegistrationOK{}
}

/* GetSignableRegistrationOK describes a response with status code 200, with default header values.

OK
*/
type GetSignableRegistrationOK struct {
	Payload *models.GetSignableRegistrationResponse
}

func (o *GetSignableRegistrationOK) Error() string {
	return fmt.Sprintf("[POST /v1/signable-registration][%d] getSignableRegistrationOK  %+v", 200, o.Payload)
}
func (o *GetSignableRegistrationOK) GetPayload() *models.GetSignableRegistrationResponse {
	return o.Payload
}

func (o *GetSignableRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSignableRegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
