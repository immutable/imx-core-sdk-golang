// Code generated by go-swagger; DO NOT EDIT.

package metadata

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

// NewGetMetadataSchemaParams creates a new GetMetadataSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMetadataSchemaParams() *GetMetadataSchemaParams {
	return &GetMetadataSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetadataSchemaParamsWithTimeout creates a new GetMetadataSchemaParams object
// with the ability to set a timeout on a request.
func NewGetMetadataSchemaParamsWithTimeout(timeout time.Duration) *GetMetadataSchemaParams {
	return &GetMetadataSchemaParams{
		timeout: timeout,
	}
}

// NewGetMetadataSchemaParamsWithContext creates a new GetMetadataSchemaParams object
// with the ability to set a context for a request.
func NewGetMetadataSchemaParamsWithContext(ctx context.Context) *GetMetadataSchemaParams {
	return &GetMetadataSchemaParams{
		Context: ctx,
	}
}

// NewGetMetadataSchemaParamsWithHTTPClient creates a new GetMetadataSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMetadataSchemaParamsWithHTTPClient(client *http.Client) *GetMetadataSchemaParams {
	return &GetMetadataSchemaParams{
		HTTPClient: client,
	}
}

/* GetMetadataSchemaParams contains all the parameters to send to the API endpoint
   for the get metadata schema operation.

   Typically these are written to a http.Request.
*/
type GetMetadataSchemaParams struct {

	/* Address.

	   Collection contract address
	*/
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get metadata schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetadataSchemaParams) WithDefaults() *GetMetadataSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get metadata schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetadataSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get metadata schema params
func (o *GetMetadataSchemaParams) WithTimeout(timeout time.Duration) *GetMetadataSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metadata schema params
func (o *GetMetadataSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metadata schema params
func (o *GetMetadataSchemaParams) WithContext(ctx context.Context) *GetMetadataSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metadata schema params
func (o *GetMetadataSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metadata schema params
func (o *GetMetadataSchemaParams) WithHTTPClient(client *http.Client) *GetMetadataSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metadata schema params
func (o *GetMetadataSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get metadata schema params
func (o *GetMetadataSchemaParams) WithAddress(address string) *GetMetadataSchemaParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get metadata schema params
func (o *GetMetadataSchemaParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetadataSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
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
