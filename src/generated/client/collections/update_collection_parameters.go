// Code generated by go-swagger; DO NOT EDIT.

package collections

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

	"immutable.com/imx-core-sdk-golang/generated/models"
)

// NewUpdateCollectionParams creates a new UpdateCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCollectionParams() *UpdateCollectionParams {
	return &UpdateCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCollectionParamsWithTimeout creates a new UpdateCollectionParams object
// with the ability to set a timeout on a request.
func NewUpdateCollectionParamsWithTimeout(timeout time.Duration) *UpdateCollectionParams {
	return &UpdateCollectionParams{
		timeout: timeout,
	}
}

// NewUpdateCollectionParamsWithContext creates a new UpdateCollectionParams object
// with the ability to set a context for a request.
func NewUpdateCollectionParamsWithContext(ctx context.Context) *UpdateCollectionParams {
	return &UpdateCollectionParams{
		Context: ctx,
	}
}

// NewUpdateCollectionParamsWithHTTPClient creates a new UpdateCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCollectionParamsWithHTTPClient(client *http.Client) *UpdateCollectionParams {
	return &UpdateCollectionParams{
		HTTPClient: client,
	}
}

/* UpdateCollectionParams contains all the parameters to send to the API endpoint
   for the update collection operation.

   Typically these are written to a http.Request.
*/
type UpdateCollectionParams struct {

	/* IMXSignature.

	   String created by signing wallet address and timestamp
	*/
	IMXSignature string

	/* IMXTimestamp.

	   Unix Epoc timestamp
	*/
	IMXTimestamp string

	/* UpdateCollectionRequest.

	   update a collection
	*/
	UpdateCollectionRequest *models.UpdateCollectionRequest

	/* Address.

	   Collection contract address
	*/
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectionParams) WithDefaults() *UpdateCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update collection params
func (o *UpdateCollectionParams) WithTimeout(timeout time.Duration) *UpdateCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update collection params
func (o *UpdateCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update collection params
func (o *UpdateCollectionParams) WithContext(ctx context.Context) *UpdateCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update collection params
func (o *UpdateCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update collection params
func (o *UpdateCollectionParams) WithHTTPClient(client *http.Client) *UpdateCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update collection params
func (o *UpdateCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIMXSignature adds the iMXSignature to the update collection params
func (o *UpdateCollectionParams) WithIMXSignature(iMXSignature string) *UpdateCollectionParams {
	o.SetIMXSignature(iMXSignature)
	return o
}

// SetIMXSignature adds the iMXSignature to the update collection params
func (o *UpdateCollectionParams) SetIMXSignature(iMXSignature string) {
	o.IMXSignature = iMXSignature
}

// WithIMXTimestamp adds the iMXTimestamp to the update collection params
func (o *UpdateCollectionParams) WithIMXTimestamp(iMXTimestamp string) *UpdateCollectionParams {
	o.SetIMXTimestamp(iMXTimestamp)
	return o
}

// SetIMXTimestamp adds the iMXTimestamp to the update collection params
func (o *UpdateCollectionParams) SetIMXTimestamp(iMXTimestamp string) {
	o.IMXTimestamp = iMXTimestamp
}

// WithUpdateCollectionRequest adds the updateCollectionRequest to the update collection params
func (o *UpdateCollectionParams) WithUpdateCollectionRequest(updateCollectionRequest *models.UpdateCollectionRequest) *UpdateCollectionParams {
	o.SetUpdateCollectionRequest(updateCollectionRequest)
	return o
}

// SetUpdateCollectionRequest adds the updateCollectionRequest to the update collection params
func (o *UpdateCollectionParams) SetUpdateCollectionRequest(updateCollectionRequest *models.UpdateCollectionRequest) {
	o.UpdateCollectionRequest = updateCollectionRequest
}

// WithAddress adds the address to the update collection params
func (o *UpdateCollectionParams) WithAddress(address string) *UpdateCollectionParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the update collection params
func (o *UpdateCollectionParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param IMX-Signature
	if err := r.SetHeaderParam("IMX-Signature", o.IMXSignature); err != nil {
		return err
	}

	// header param IMX-Timestamp
	if err := r.SetHeaderParam("IMX-Timestamp", o.IMXTimestamp); err != nil {
		return err
	}
	if o.UpdateCollectionRequest != nil {
		if err := r.SetBodyParam(o.UpdateCollectionRequest); err != nil {
			return err
		}
	}

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}