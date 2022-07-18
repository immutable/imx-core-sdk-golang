// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// GetSignableCancelOrderReader is a Reader for the GetSignableCancelOrder structure.
type GetSignableCancelOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignableCancelOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignableCancelOrderOK()
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

// NewGetSignableCancelOrderOK creates a GetSignableCancelOrderOK with default headers values
func NewGetSignableCancelOrderOK() *GetSignableCancelOrderOK {
	return &GetSignableCancelOrderOK{}
}

/* GetSignableCancelOrderOK describes a response with status code 200, with default header values.

OK
*/
type GetSignableCancelOrderOK struct {
	Payload *models.GetSignableCancelOrderResponse
}

func (o *GetSignableCancelOrderOK) Error() string {
	return fmt.Sprintf("[POST /v1/signable-cancel-order-details][%d] getSignableCancelOrderOK  %+v", 200, o.Payload)
}
func (o *GetSignableCancelOrderOK) GetPayload() *models.GetSignableCancelOrderResponse {
	return o.Payload
}

func (o *GetSignableCancelOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSignableCancelOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
