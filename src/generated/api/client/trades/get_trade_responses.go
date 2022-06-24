// Code generated by go-swagger; DO NOT EDIT.

package trades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"immutable.com/imx-core-sdk-golang/api/models"
)

// GetTradeReader is a Reader for the GetTrade structure.
type GetTradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTradeOK creates a GetTradeOK with default headers values
func NewGetTradeOK() *GetTradeOK {
	return &GetTradeOK{}
}

/* GetTradeOK describes a response with status code 200, with default header values.

OK
*/
type GetTradeOK struct {
	Payload *models.Trade
}

func (o *GetTradeOK) Error() string {
	return fmt.Sprintf("[GET /v1/trades/{id}][%d] getTradeOK  %+v", 200, o.Payload)
}
func (o *GetTradeOK) GetPayload() *models.Trade {
	return o.Payload
}

func (o *GetTradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trade)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}