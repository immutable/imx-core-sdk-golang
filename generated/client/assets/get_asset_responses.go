// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generated/models"
)

// GetAssetReader is a Reader for the GetAsset structure.
type GetAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAssetOK creates a GetAssetOK with default headers values
func NewGetAssetOK() *GetAssetOK {
	return &GetAssetOK{}
}

/* GetAssetOK describes a response with status code 200, with default header values.

OK
*/
type GetAssetOK struct {
	Payload *models.Asset
}

func (o *GetAssetOK) Error() string {
	return fmt.Sprintf("[GET /v1/assets/{token_address}/{token_id}][%d] getAssetOK  %+v", 200, o.Payload)
}
func (o *GetAssetOK) GetPayload() *models.Asset {
	return o.Payload
}

func (o *GetAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Asset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
