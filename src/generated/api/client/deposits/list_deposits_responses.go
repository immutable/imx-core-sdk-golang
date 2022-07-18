// Code generated by go-swagger; DO NOT EDIT.

package deposits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"immutable.com/imx-core-sdk-golang/api/models"
)

// ListDepositsReader is a Reader for the ListDeposits structure.
type ListDepositsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDepositsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDepositsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		body, err := io.ReadAll(response.Body())
		if err == nil {
			var message map[string]interface{}
			err = json.Unmarshal(body, &message)
			if err == nil {
				return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", message, response.Code())
			}
		}
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListDepositsOK creates a ListDepositsOK with default headers values
func NewListDepositsOK() *ListDepositsOK {
	return &ListDepositsOK{}
}

/* ListDepositsOK describes a response with status code 200, with default header values.

OK
*/
type ListDepositsOK struct {
	Payload *models.ListDepositsResponse
}

func (o *ListDepositsOK) Error() string {
	return fmt.Sprintf("[GET /v1/deposits][%d] listDepositsOK  %+v", 200, o.Payload)
}
func (o *ListDepositsOK) GetPayload() *models.ListDepositsResponse {
	return o.Payload
}

func (o *ListDepositsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDepositsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
