// Code generated by go-swagger; DO NOT EDIT.

package withdrawals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"immutable.com/imx-core-sdk-golang/api/models"
)

// NewCreateWithdrawalParams creates a new CreateWithdrawalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWithdrawalParams() *CreateWithdrawalParams {
	return &CreateWithdrawalParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateWithdrawalParamsWithTimeout creates a new CreateWithdrawalParams object
// with the ability to set a timeout on a request.
func NewCreateWithdrawalParamsWithTimeout(timeout time.Duration) *CreateWithdrawalParams {
	return &CreateWithdrawalParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateWithdrawalParamsWithContext creates a new CreateWithdrawalParams object
// with the ability to set a context for a request.
func NewCreateWithdrawalParamsWithContext(ctx context.Context) *CreateWithdrawalParams {
	return &CreateWithdrawalParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateWithdrawalParamsWithHTTPClient creates a new CreateWithdrawalParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWithdrawalParamsWithHTTPClient(client *http.Client) *CreateWithdrawalParams {
	return &CreateWithdrawalParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* CreateWithdrawalParams contains all the parameters to send to the API endpoint
   for the create withdrawal operation.

   Typically these are written to a http.Request.
*/
type CreateWithdrawalParams struct {

	/* CreateWithdrawalRequest.

	   create a withdrawal
	*/
	CreateWithdrawalRequest *models.CreateWithdrawalRequest

	/* XImxEthAddress.

	   eth address
	*/
	XImxEthAddress *string

	/* XImxEthSignature.

	   eth signature
	*/
	XImxEthSignature *string

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create withdrawal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWithdrawalParams) WithDefaults() *CreateWithdrawalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create withdrawal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWithdrawalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create withdrawal params
func (o *CreateWithdrawalParams) WithTimeout(timeout time.Duration) *CreateWithdrawalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create withdrawal params
func (o *CreateWithdrawalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create withdrawal params
func (o *CreateWithdrawalParams) WithContext(ctx context.Context) *CreateWithdrawalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create withdrawal params
func (o *CreateWithdrawalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to create withdrawal params.
func (o *CreateWithdrawalParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the create withdrawal params
func (o *CreateWithdrawalParams) WithHTTPClient(client *http.Client) *CreateWithdrawalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create withdrawal params
func (o *CreateWithdrawalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateWithdrawalRequest adds the createWithdrawalRequest to the create withdrawal params
func (o *CreateWithdrawalParams) WithCreateWithdrawalRequest(createWithdrawalRequest *models.CreateWithdrawalRequest) *CreateWithdrawalParams {
	o.SetCreateWithdrawalRequest(createWithdrawalRequest)
	return o
}

// SetCreateWithdrawalRequest adds the createWithdrawalRequest to the create withdrawal params
func (o *CreateWithdrawalParams) SetCreateWithdrawalRequest(createWithdrawalRequest *models.CreateWithdrawalRequest) {
	o.CreateWithdrawalRequest = createWithdrawalRequest
}

// WithXImxEthAddress adds the xImxEthAddress to the create withdrawal params
func (o *CreateWithdrawalParams) WithXImxEthAddress(xImxEthAddress *string) *CreateWithdrawalParams {
	o.SetXImxEthAddress(xImxEthAddress)
	return o
}

// SetXImxEthAddress adds the xImxEthAddress to the create withdrawal params
func (o *CreateWithdrawalParams) SetXImxEthAddress(xImxEthAddress *string) {
	o.XImxEthAddress = xImxEthAddress
}

// WithXImxEthSignature adds the xImxEthSignature to the create withdrawal params
func (o *CreateWithdrawalParams) WithXImxEthSignature(xImxEthSignature *string) *CreateWithdrawalParams {
	o.SetXImxEthSignature(xImxEthSignature)
	return o
}

// SetXImxEthSignature adds the xImxEthSignature to the create withdrawal params
func (o *CreateWithdrawalParams) SetXImxEthSignature(xImxEthSignature *string) {
	o.XImxEthSignature = xImxEthSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWithdrawalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}

	if err := r.SetHeaderParam("x-sdk-version", "imx-core-sdk-golang-"); err != nil {
		return err
	}

	for key, val := range o.AdditionalHeaderParams {
		if err := r.SetHeaderParam(key, val); err != nil {
			return err
		}
	}

	var res []error
	if o.CreateWithdrawalRequest != nil {
		if err := r.SetBodyParam(o.CreateWithdrawalRequest); err != nil {
			return err
		}
	}

	if o.XImxEthAddress != nil {

		// header param x-imx-eth-address
		if err := r.SetHeaderParam("x-imx-eth-address", *o.XImxEthAddress); err != nil {
			return err
		}
	}

	if o.XImxEthSignature != nil {

		// header param x-imx-eth-signature
		if err := r.SetHeaderParam("x-imx-eth-signature", *o.XImxEthSignature); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
