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

// NewGetSignableRegistrationOffchainParams creates a new GetSignableRegistrationOffchainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSignableRegistrationOffchainParams() *GetSignableRegistrationOffchainParams {
	return &GetSignableRegistrationOffchainParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetSignableRegistrationOffchainParamsWithTimeout creates a new GetSignableRegistrationOffchainParams object
// with the ability to set a timeout on a request.
func NewGetSignableRegistrationOffchainParamsWithTimeout(timeout time.Duration) *GetSignableRegistrationOffchainParams {
	return &GetSignableRegistrationOffchainParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetSignableRegistrationOffchainParamsWithContext creates a new GetSignableRegistrationOffchainParams object
// with the ability to set a context for a request.
func NewGetSignableRegistrationOffchainParamsWithContext(ctx context.Context) *GetSignableRegistrationOffchainParams {
	return &GetSignableRegistrationOffchainParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewGetSignableRegistrationOffchainParamsWithHTTPClient creates a new GetSignableRegistrationOffchainParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSignableRegistrationOffchainParamsWithHTTPClient(client *http.Client) *GetSignableRegistrationOffchainParams {
	return &GetSignableRegistrationOffchainParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* GetSignableRegistrationOffchainParams contains all the parameters to send to the API endpoint
   for the get signable registration offchain operation.

   Typically these are written to a http.Request.
*/
type GetSignableRegistrationOffchainParams struct {

	/* GetSignableRegistrationRequest.

	   Register User Offchain
	*/
	GetSignableRegistrationRequest *models.GetSignableRegistrationRequest

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get signable registration offchain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableRegistrationOffchainParams) WithDefaults() *GetSignableRegistrationOffchainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get signable registration offchain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignableRegistrationOffchainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) WithTimeout(timeout time.Duration) *GetSignableRegistrationOffchainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) WithContext(ctx context.Context) *GetSignableRegistrationOffchainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to get signable registration offchain params.
func (o *GetSignableRegistrationOffchainParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) WithHTTPClient(client *http.Client) *GetSignableRegistrationOffchainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetSignableRegistrationRequest adds the getSignableRegistrationRequest to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) WithGetSignableRegistrationRequest(getSignableRegistrationRequest *models.GetSignableRegistrationRequest) *GetSignableRegistrationOffchainParams {
	o.SetGetSignableRegistrationRequest(getSignableRegistrationRequest)
	return o
}

// SetGetSignableRegistrationRequest adds the getSignableRegistrationRequest to the get signable registration offchain params
func (o *GetSignableRegistrationOffchainParams) SetGetSignableRegistrationRequest(getSignableRegistrationRequest *models.GetSignableRegistrationRequest) {
	o.GetSignableRegistrationRequest = getSignableRegistrationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetSignableRegistrationOffchainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}

	for key, val := range o.AdditionalHeaderParams {
		if err := r.SetHeaderParam(key, val); err != nil {
			return err
		}
	}

	// Add SDK version header.
	if err := r.SetHeaderParam("x-sdk-version", "imx-core-sdk-golang-0.1.0"); err != nil {
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
