// Code generated by go-swagger; DO NOT EDIT.

package balances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// GetBalanceReader is a Reader for the GetBalance structure.
type GetBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBalanceOK creates a GetBalanceOK with default headers values
func NewGetBalanceOK() *GetBalanceOK {
	return &GetBalanceOK{}
}

/* GetBalanceOK describes a response with status code 200, with default header values.

OK
*/
type GetBalanceOK struct {
	Payload *models.Balance
}

func (o *GetBalanceOK) Error() string {
	return fmt.Sprintf("[GET /v2/balances/{owner}/{address}][%d] getBalanceOK  %+v", 200, o.Payload)
}
func (o *GetBalanceOK) GetPayload() *models.Balance {
	return o.Payload
}

func (o *GetBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Balance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
