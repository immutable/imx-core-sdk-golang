// Code generated by go-swagger; DO NOT EDIT.

package trades

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

// NewGetTradeParams creates a new GetTradeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTradeParams() *GetTradeParams {
	return &GetTradeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTradeParamsWithTimeout creates a new GetTradeParams object
// with the ability to set a timeout on a request.
func NewGetTradeParamsWithTimeout(timeout time.Duration) *GetTradeParams {
	return &GetTradeParams{
		timeout: timeout,
	}
}

// NewGetTradeParamsWithContext creates a new GetTradeParams object
// with the ability to set a context for a request.
func NewGetTradeParamsWithContext(ctx context.Context) *GetTradeParams {
	return &GetTradeParams{
		Context: ctx,
	}
}

// NewGetTradeParamsWithHTTPClient creates a new GetTradeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTradeParamsWithHTTPClient(client *http.Client) *GetTradeParams {
	return &GetTradeParams{
		HTTPClient: client,
	}
}

/* GetTradeParams contains all the parameters to send to the API endpoint
   for the get trade operation.

   Typically these are written to a http.Request.
*/
type GetTradeParams struct {

	/* ID.

	   Trade ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTradeParams) WithDefaults() *GetTradeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTradeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get trade params
func (o *GetTradeParams) WithTimeout(timeout time.Duration) *GetTradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trade params
func (o *GetTradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trade params
func (o *GetTradeParams) WithContext(ctx context.Context) *GetTradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trade params
func (o *GetTradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trade params
func (o *GetTradeParams) WithHTTPClient(client *http.Client) *GetTradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trade params
func (o *GetTradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get trade params
func (o *GetTradeParams) WithID(id string) *GetTradeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get trade params
func (o *GetTradeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
