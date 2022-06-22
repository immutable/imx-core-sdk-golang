// Code generated by go-swagger; DO NOT EDIT.

package trades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// GetSignableTradeReader is a Reader for the GetSignableTrade structure.
type GetSignableTradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignableTradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignableTradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSignableTradeOK creates a GetSignableTradeOK with default headers values
func NewGetSignableTradeOK() *GetSignableTradeOK {
	return &GetSignableTradeOK{}
}

/* GetSignableTradeOK describes a response with status code 200, with default header values.

OK
*/
type GetSignableTradeOK struct {
	Payload *models.GetSignableTradeResponse
}

func (o *GetSignableTradeOK) Error() string {
	return fmt.Sprintf("[POST /v3/signable-trade-details][%d] getSignableTradeOK  %+v", 200, o.Payload)
}
func (o *GetSignableTradeOK) GetPayload() *models.GetSignableTradeResponse {
	return o.Payload
}

func (o *GetSignableTradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSignableTradeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
