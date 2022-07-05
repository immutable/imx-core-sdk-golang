// Code generated by go-swagger; DO NOT EDIT.

package balances

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

// NewListBalancesParams creates a new ListBalancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBalancesParams() *ListBalancesParams {
	return &ListBalancesParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListBalancesParamsWithTimeout creates a new ListBalancesParams object
// with the ability to set a timeout on a request.
func NewListBalancesParamsWithTimeout(timeout time.Duration) *ListBalancesParams {
	return &ListBalancesParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListBalancesParamsWithContext creates a new ListBalancesParams object
// with the ability to set a context for a request.
func NewListBalancesParamsWithContext(ctx context.Context) *ListBalancesParams {
	return &ListBalancesParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListBalancesParamsWithHTTPClient creates a new ListBalancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBalancesParamsWithHTTPClient(client *http.Client) *ListBalancesParams {
	return &ListBalancesParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* ListBalancesParams contains all the parameters to send to the API endpoint
   for the list balances operation.

   Typically these are written to a http.Request.
*/
type ListBalancesParams struct {

	/* Owner.

	   Ethereum wallet address for user
	*/
	Owner string

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list balances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBalancesParams) WithDefaults() *ListBalancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list balances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBalancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list balances params
func (o *ListBalancesParams) WithTimeout(timeout time.Duration) *ListBalancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list balances params
func (o *ListBalancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list balances params
func (o *ListBalancesParams) WithContext(ctx context.Context) *ListBalancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list balances params
func (o *ListBalancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to list balances params.
func (o *ListBalancesParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the list balances params
func (o *ListBalancesParams) WithHTTPClient(client *http.Client) *ListBalancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list balances params
func (o *ListBalancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the list balances params
func (o *ListBalancesParams) WithOwner(owner string) *ListBalancesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list balances params
func (o *ListBalancesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ListBalancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
