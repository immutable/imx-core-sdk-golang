// Code generated by go-swagger; DO NOT EDIT.

package tokens

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
)

// NewGetTokenParams creates a new GetTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTokenParams() *GetTokenParams {
	return &GetTokenParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetTokenParamsWithTimeout creates a new GetTokenParams object
// with the ability to set a timeout on a request.
func NewGetTokenParamsWithTimeout(timeout time.Duration) *GetTokenParams {
	return &GetTokenParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetTokenParamsWithContext creates a new GetTokenParams object
// with the ability to set a context for a request.
func NewGetTokenParamsWithContext(ctx context.Context) *GetTokenParams {
	return &GetTokenParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetTokenParamsWithHTTPClient creates a new GetTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTokenParamsWithHTTPClient(client *http.Client) *GetTokenParams {
	return &GetTokenParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* GetTokenParams contains all the parameters to send to the API endpoint
   for the get token operation.

   Typically these are written to a http.Request.
*/
type GetTokenParams struct {

	/* Address.

	   Token Contract Address
	*/
	Address string

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTokenParams) WithDefaults() *GetTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get token params
func (o *GetTokenParams) WithTimeout(timeout time.Duration) *GetTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get token params
func (o *GetTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get token params
func (o *GetTokenParams) WithContext(ctx context.Context) *GetTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get token params
func (o *GetTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to get token params.
func (o *GetTokenParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the get token params
func (o *GetTokenParams) WithHTTPClient(client *http.Client) *GetTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get token params
func (o *GetTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get token params
func (o *GetTokenParams) WithAddress(address string) *GetTokenParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get token params
func (o *GetTokenParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *GetTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
