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

// ListTradesReader is a Reader for the ListTrades structure.
type ListTradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTradesOK creates a ListTradesOK with default headers values
func NewListTradesOK() *ListTradesOK {
	return &ListTradesOK{}
}

/* ListTradesOK describes a response with status code 200, with default header values.

OK
*/
type ListTradesOK struct {
	Payload *models.ListTradesResponse
}

func (o *ListTradesOK) Error() string {
	return fmt.Sprintf("[GET /v1/trades][%d] listTradesOK  %+v", 200, o.Payload)
}
func (o *ListTradesOK) GetPayload() *models.ListTradesResponse {
	return o.Payload
}

func (o *ListTradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTradesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
