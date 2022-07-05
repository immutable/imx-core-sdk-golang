// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// NewCreateOrderParams creates a new CreateOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrderParams() *CreateOrderParams {
	return &CreateOrderParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateOrderParamsWithTimeout creates a new CreateOrderParams object
// with the ability to set a timeout on a request.
func NewCreateOrderParamsWithTimeout(timeout time.Duration) *CreateOrderParams {
	return &CreateOrderParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateOrderParamsWithContext creates a new CreateOrderParams object
// with the ability to set a context for a request.
func NewCreateOrderParamsWithContext(ctx context.Context) *CreateOrderParams {
	return &CreateOrderParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewCreateOrderParamsWithHTTPClient creates a new CreateOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrderParamsWithHTTPClient(client *http.Client) *CreateOrderParams {
	return &CreateOrderParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* CreateOrderParams contains all the parameters to send to the API endpoint
   for the create order operation.

   Typically these are written to a http.Request.
*/
type CreateOrderParams struct {

	/* CreateOrderRequest.

	   create an order
	*/
	CreateOrderRequest *models.CreateOrderRequest

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

// WithDefaults hydrates default values in the create order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrderParams) WithDefaults() *CreateOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create order params
func (o *CreateOrderParams) WithTimeout(timeout time.Duration) *CreateOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create order params
func (o *CreateOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create order params
func (o *CreateOrderParams) WithContext(ctx context.Context) *CreateOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create order params
func (o *CreateOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to create order params.
func (o *CreateOrderParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the create order params
func (o *CreateOrderParams) WithHTTPClient(client *http.Client) *CreateOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create order params
func (o *CreateOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrderRequest adds the createOrderRequest to the create order params
func (o *CreateOrderParams) WithCreateOrderRequest(createOrderRequest *models.CreateOrderRequest) *CreateOrderParams {
	o.SetCreateOrderRequest(createOrderRequest)
	return o
}

// SetCreateOrderRequest adds the createOrderRequest to the create order params
func (o *CreateOrderParams) SetCreateOrderRequest(createOrderRequest *models.CreateOrderRequest) {
	o.CreateOrderRequest = createOrderRequest
}

// WithXImxEthAddress adds the xImxEthAddress to the create order params
func (o *CreateOrderParams) WithXImxEthAddress(xImxEthAddress *string) *CreateOrderParams {
	o.SetXImxEthAddress(xImxEthAddress)
	return o
}

// SetXImxEthAddress adds the xImxEthAddress to the create order params
func (o *CreateOrderParams) SetXImxEthAddress(xImxEthAddress *string) {
	o.XImxEthAddress = xImxEthAddress
}

// WithXImxEthSignature adds the xImxEthSignature to the create order params
func (o *CreateOrderParams) WithXImxEthSignature(xImxEthSignature *string) *CreateOrderParams {
	o.SetXImxEthSignature(xImxEthSignature)
	return o
}

// SetXImxEthSignature adds the xImxEthSignature to the create order params
func (o *CreateOrderParams) SetXImxEthSignature(xImxEthSignature *string) {
	o.XImxEthSignature = xImxEthSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.CreateOrderRequest != nil {
		if err := r.SetBodyParam(o.CreateOrderRequest); err != nil {
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
