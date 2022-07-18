// Code generated by go-swagger; DO NOT EDIT.

package mints

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

// MintTokensReader is a Reader for the MintTokens structure.
type MintTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MintTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMintTokensOK()
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

// NewMintTokensOK creates a MintTokensOK with default headers values
func NewMintTokensOK() *MintTokensOK {
	return &MintTokensOK{}
}

/* MintTokensOK describes a response with status code 200, with default header values.

OK
*/
type MintTokensOK struct {

	/* The mint limit available to the project for each four-week window.
	 */
	MintLimit string

	/* The expiry date of the current four-week window.
	 */
	MintLimitReset string

	/* The amount of mints remaining for current four-week window.
	 */
	MintRemaining string

	Payload *models.MintTokensResponse
}

func (o *MintTokensOK) Error() string {
	return fmt.Sprintf("[POST /v2/mints][%d] mintTokensOK  %+v", 200, o.Payload)
}
func (o *MintTokensOK) GetPayload() *models.MintTokensResponse {
	return o.Payload
}

func (o *MintTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Mint_Limit
	hdrMintLimit := response.GetHeader("Mint_Limit")

	if hdrMintLimit != "" {
		o.MintLimit = hdrMintLimit
	}

	// hydrates response header Mint_Limit_Reset
	hdrMintLimitReset := response.GetHeader("Mint_Limit_Reset")

	if hdrMintLimitReset != "" {
		o.MintLimitReset = hdrMintLimitReset
	}

	// hydrates response header Mint_Remaining
	hdrMintRemaining := response.GetHeader("Mint_Remaining")

	if hdrMintRemaining != "" {
		o.MintRemaining = hdrMintRemaining
	}

	o.Payload = new(models.MintTokensResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
