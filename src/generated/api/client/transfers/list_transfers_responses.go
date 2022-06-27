// Code generated by go-swagger; DO NOT EDIT.

package transfers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"immutable.com/imx-core-sdk-golang/api/models"
)

// ListTransfersReader is a Reader for the ListTransfers structure.
type ListTransfersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTransfersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTransfersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTransfersOK creates a ListTransfersOK with default headers values
func NewListTransfersOK() *ListTransfersOK {
	return &ListTransfersOK{}
}

/* ListTransfersOK describes a response with status code 200, with default header values.

OK
*/
type ListTransfersOK struct {
	Payload *models.ListTransfersResponse
}

func (o *ListTransfersOK) Error() string {
	return fmt.Sprintf("[GET /v1/transfers][%d] listTransfersOK  %+v", 200, o.Payload)
}
func (o *ListTransfersOK) GetPayload() *models.ListTransfersResponse {
	return o.Payload
}

func (o *ListTransfersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTransfersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
