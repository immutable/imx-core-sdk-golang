// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetSignableRegistrationParams creates a new GetSignableRegistrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSignableRegistrationParams() *GetSignableRegistrationParams {
	return &GetSignableRegistrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSignableRegistrationParamsWithTimeout creates a new GetSignableRegistrationParams object
// with the ability to set a timeout on a request.
func NewGetSignableRegistrationParamsWithTimeout(timeout time.Duration) *GetSignableRegistrationParams {
	return &GetSignableRegistrationParams{
		timeout: timeout,
	}
}

// NewGetSignableRegistrationParamsWithContext creates a new GetSignableRegistrationParams object
// with the ability to set a context for a request.
func NewGetSignableRegistrationParamsWithContext(ctx context.Context) *GetSignableRegistrationParams {
	return &GetSignableRegistrationParams{
		Context: ctx,
	}
}

// NewGetSignableRegistrationParamsWithHTTPClient creates a new GetSignableRegistrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSignableRegistrationParamsWithHTTPClient(client *http.Client) *GetSignableRegistrationParams {
	return &GetSignableRegistrationParams{
		HTTPClient: client,
	}
}

/* GetSignableRegistrationParams contains all the parameters to send to the API endpoint
   for the get signable registration operation.

   Typically these are written to a http.Request.
*/
type GetSignableRegistrationParams struct {

	/* GetSignableRegistrationRequest.

	   Register User
	*/
	GetSignableRegistrationRequest *models.GetSignableRegistrationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get signable registration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableRegistrationParams) WithDefaults() *GetSignableRegistrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get signable registration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableRegistrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get signable registration params
func (o *GetSignableRegistrationParams) WithTimeout(timeout time.Duration) *GetSignableRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get signable registration params
func (o *GetSignableRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get signable registration params
func (o *GetSignableRegistrationParams) WithContext(ctx context.Context) *GetSignableRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get signable registration params
func (o *GetSignableRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get signable registration params
func (o *GetSignableRegistrationParams) WithHTTPClient(client *http.Client) *GetSignableRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get signable registration params
func (o *GetSignableRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetSignableRegistrationRequest adds the getSignableRegistrationRequest to the get signable registration params
func (o *GetSignableRegistrationParams) WithGetSignableRegistrationRequest(getSignableRegistrationRequest *models.GetSignableRegistrationRequest) *GetSignableRegistrationParams {
	o.SetGetSignableRegistrationRequest(getSignableRegistrationRequest)
	return o
}

// SetGetSignableRegistrationRequest adds the getSignableRegistrationRequest to the get signable registration params
func (o *GetSignableRegistrationParams) SetGetSignableRegistrationRequest(getSignableRegistrationRequest *models.GetSignableRegistrationRequest) {
	o.GetSignableRegistrationRequest = getSignableRegistrationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetSignableRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GetSignableRegistrationRequest != nil {
		if err := r.SetBodyParam(o.GetSignableRegistrationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}