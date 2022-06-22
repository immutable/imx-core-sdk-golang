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

	"generated/models"
)

// NewGetSignableOrderParams creates a new GetSignableOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSignableOrderParams() *GetSignableOrderParams {
	return &GetSignableOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSignableOrderParamsWithTimeout creates a new GetSignableOrderParams object
// with the ability to set a timeout on a request.
func NewGetSignableOrderParamsWithTimeout(timeout time.Duration) *GetSignableOrderParams {
	return &GetSignableOrderParams{
		timeout: timeout,
	}
}

// NewGetSignableOrderParamsWithContext creates a new GetSignableOrderParams object
// with the ability to set a context for a request.
func NewGetSignableOrderParamsWithContext(ctx context.Context) *GetSignableOrderParams {
	return &GetSignableOrderParams{
		Context: ctx,
	}
}

// NewGetSignableOrderParamsWithHTTPClient creates a new GetSignableOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSignableOrderParamsWithHTTPClient(client *http.Client) *GetSignableOrderParams {
	return &GetSignableOrderParams{
		HTTPClient: client,
	}
}

/* GetSignableOrderParams contains all the parameters to send to the API endpoint
   for the get signable order operation.

   Typically these are written to a http.Request.
*/
type GetSignableOrderParams struct {

	/* GetSignableOrderRequestV3.

	   get a signable order
	*/
	GetSignableOrderRequestV3 *models.GetSignableOrderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get signable order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableOrderParams) WithDefaults() *GetSignableOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get signable order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get signable order params
func (o *GetSignableOrderParams) WithTimeout(timeout time.Duration) *GetSignableOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get signable order params
func (o *GetSignableOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get signable order params
func (o *GetSignableOrderParams) WithContext(ctx context.Context) *GetSignableOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get signable order params
func (o *GetSignableOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get signable order params
func (o *GetSignableOrderParams) WithHTTPClient(client *http.Client) *GetSignableOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get signable order params
func (o *GetSignableOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetSignableOrderRequestV3 adds the getSignableOrderRequestV3 to the get signable order params
func (o *GetSignableOrderParams) WithGetSignableOrderRequestV3(getSignableOrderRequestV3 *models.GetSignableOrderRequest) *GetSignableOrderParams {
	o.SetGetSignableOrderRequestV3(getSignableOrderRequestV3)
	return o
}

// SetGetSignableOrderRequestV3 adds the getSignableOrderRequestV3 to the get signable order params
func (o *GetSignableOrderParams) SetGetSignableOrderRequestV3(getSignableOrderRequestV3 *models.GetSignableOrderRequest) {
	o.GetSignableOrderRequestV3 = getSignableOrderRequestV3
}

// WriteToRequest writes these params to a swagger request
func (o *GetSignableOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GetSignableOrderRequestV3 != nil {
		if err := r.SetBodyParam(o.GetSignableOrderRequestV3); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
