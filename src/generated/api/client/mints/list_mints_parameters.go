// Code generated by go-swagger; DO NOT EDIT.

package mints

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
	"github.com/go-openapi/swag"
)

// NewListMintsParams creates a new ListMintsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMintsParams() *ListMintsParams {
	return &ListMintsParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListMintsParamsWithTimeout creates a new ListMintsParams object
// with the ability to set a timeout on a request.
func NewListMintsParamsWithTimeout(timeout time.Duration) *ListMintsParams {
	return &ListMintsParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListMintsParamsWithContext creates a new ListMintsParams object
// with the ability to set a context for a request.
func NewListMintsParamsWithContext(ctx context.Context) *ListMintsParams {
	return &ListMintsParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListMintsParamsWithHTTPClient creates a new ListMintsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMintsParamsWithHTTPClient(client *http.Client) *ListMintsParams {
	return &ListMintsParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* ListMintsParams contains all the parameters to send to the API endpoint
   for the list mints operation.

   Typically these are written to a http.Request.
*/
type ListMintsParams struct {

	/* AssetID.

	   Internal IMX ID of the minted asset
	*/
	AssetID *string

	/* Cursor.

	   Cursor
	*/
	Cursor *string

	/* Direction.

	   Direction to sort (asc/desc)
	*/
	Direction *string

	/* MaxQuantity.

	   Max quantity for the minted asset
	*/
	MaxQuantity *string

	/* Metadata.

	   JSON-encoded metadata filters for the minted asset
	*/
	Metadata *string

	/* MinQuantity.

	   Min quantity for the minted asset
	*/
	MinQuantity *string

	/* OrderBy.

	   Property to sort by
	*/
	OrderBy *string

	/* PageSize.

	   Page size of the result
	*/
	PageSize *int64

	/* Status.

	   Status of this mint
	*/
	Status *string

	/* TokenAddress.

	   Token address of the minted asset
	*/
	TokenAddress *string

	/* TokenID.

	   ERC721 Token ID of the minted asset
	*/
	TokenID *string

	/* TokenName.

	   Token Name of the minted asset
	*/
	TokenName *string

	/* TokenType.

	   Token type of the minted asset
	*/
	TokenType *string

	/* UpdatedMaxTimestamp.

	   Maximum timestamp for this mint, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	UpdatedMaxTimestamp *string

	/* UpdatedMinTimestamp.

	   Minimum timestamp for this mint, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	UpdatedMinTimestamp *string

	/* User.

	   Ethereum address of the user who submitted this mint
	*/
	User *string

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list mints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMintsParams) WithDefaults() *ListMintsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list mints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMintsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list mints params
func (o *ListMintsParams) WithTimeout(timeout time.Duration) *ListMintsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mints params
func (o *ListMintsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mints params
func (o *ListMintsParams) WithContext(ctx context.Context) *ListMintsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mints params
func (o *ListMintsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to list mints params.
func (o *ListMintsParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the list mints params
func (o *ListMintsParams) WithHTTPClient(client *http.Client) *ListMintsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mints params
func (o *ListMintsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the list mints params
func (o *ListMintsParams) WithAssetID(assetID *string) *ListMintsParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the list mints params
func (o *ListMintsParams) SetAssetID(assetID *string) {
	o.AssetID = assetID
}

// WithCursor adds the cursor to the list mints params
func (o *ListMintsParams) WithCursor(cursor *string) *ListMintsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list mints params
func (o *ListMintsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithDirection adds the direction to the list mints params
func (o *ListMintsParams) WithDirection(direction *string) *ListMintsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the list mints params
func (o *ListMintsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithMaxQuantity adds the maxQuantity to the list mints params
func (o *ListMintsParams) WithMaxQuantity(maxQuantity *string) *ListMintsParams {
	o.SetMaxQuantity(maxQuantity)
	return o
}

// SetMaxQuantity adds the maxQuantity to the list mints params
func (o *ListMintsParams) SetMaxQuantity(maxQuantity *string) {
	o.MaxQuantity = maxQuantity
}

// WithMetadata adds the metadata to the list mints params
func (o *ListMintsParams) WithMetadata(metadata *string) *ListMintsParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the list mints params
func (o *ListMintsParams) SetMetadata(metadata *string) {
	o.Metadata = metadata
}

// WithMinQuantity adds the minQuantity to the list mints params
func (o *ListMintsParams) WithMinQuantity(minQuantity *string) *ListMintsParams {
	o.SetMinQuantity(minQuantity)
	return o
}

// SetMinQuantity adds the minQuantity to the list mints params
func (o *ListMintsParams) SetMinQuantity(minQuantity *string) {
	o.MinQuantity = minQuantity
}

// WithOrderBy adds the orderBy to the list mints params
func (o *ListMintsParams) WithOrderBy(orderBy *string) *ListMintsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list mints params
func (o *ListMintsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPageSize adds the pageSize to the list mints params
func (o *ListMintsParams) WithPageSize(pageSize *int64) *ListMintsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list mints params
func (o *ListMintsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithStatus adds the status to the list mints params
func (o *ListMintsParams) WithStatus(status *string) *ListMintsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list mints params
func (o *ListMintsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTokenAddress adds the tokenAddress to the list mints params
func (o *ListMintsParams) WithTokenAddress(tokenAddress *string) *ListMintsParams {
	o.SetTokenAddress(tokenAddress)
	return o
}

// SetTokenAddress adds the tokenAddress to the list mints params
func (o *ListMintsParams) SetTokenAddress(tokenAddress *string) {
	o.TokenAddress = tokenAddress
}

// WithTokenID adds the tokenID to the list mints params
func (o *ListMintsParams) WithTokenID(tokenID *string) *ListMintsParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the list mints params
func (o *ListMintsParams) SetTokenID(tokenID *string) {
	o.TokenID = tokenID
}

// WithTokenName adds the tokenName to the list mints params
func (o *ListMintsParams) WithTokenName(tokenName *string) *ListMintsParams {
	o.SetTokenName(tokenName)
	return o
}

// SetTokenName adds the tokenName to the list mints params
func (o *ListMintsParams) SetTokenName(tokenName *string) {
	o.TokenName = tokenName
}

// WithTokenType adds the tokenType to the list mints params
func (o *ListMintsParams) WithTokenType(tokenType *string) *ListMintsParams {
	o.SetTokenType(tokenType)
	return o
}

// SetTokenType adds the tokenType to the list mints params
func (o *ListMintsParams) SetTokenType(tokenType *string) {
	o.TokenType = tokenType
}

// WithUpdatedMaxTimestamp adds the updatedMaxTimestamp to the list mints params
func (o *ListMintsParams) WithUpdatedMaxTimestamp(updatedMaxTimestamp *string) *ListMintsParams {
	o.SetUpdatedMaxTimestamp(updatedMaxTimestamp)
	return o
}

// SetUpdatedMaxTimestamp adds the updatedMaxTimestamp to the list mints params
func (o *ListMintsParams) SetUpdatedMaxTimestamp(updatedMaxTimestamp *string) {
	o.UpdatedMaxTimestamp = updatedMaxTimestamp
}

// WithUpdatedMinTimestamp adds the updatedMinTimestamp to the list mints params
func (o *ListMintsParams) WithUpdatedMinTimestamp(updatedMinTimestamp *string) *ListMintsParams {
	o.SetUpdatedMinTimestamp(updatedMinTimestamp)
	return o
}

// SetUpdatedMinTimestamp adds the updatedMinTimestamp to the list mints params
func (o *ListMintsParams) SetUpdatedMinTimestamp(updatedMinTimestamp *string) {
	o.UpdatedMinTimestamp = updatedMinTimestamp
}

// WithUser adds the user to the list mints params
func (o *ListMintsParams) WithUser(user *string) *ListMintsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list mints params
func (o *ListMintsParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListMintsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AssetID != nil {

		// query param asset_id
		var qrAssetID string

		if o.AssetID != nil {
			qrAssetID = *o.AssetID
		}
		qAssetID := qrAssetID
		if qAssetID != "" {

			if err := r.SetQueryParam("asset_id", qAssetID); err != nil {
				return err
			}
		}
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}
	}

	if o.MaxQuantity != nil {

		// query param max_quantity
		var qrMaxQuantity string

		if o.MaxQuantity != nil {
			qrMaxQuantity = *o.MaxQuantity
		}
		qMaxQuantity := qrMaxQuantity
		if qMaxQuantity != "" {

			if err := r.SetQueryParam("max_quantity", qMaxQuantity); err != nil {
				return err
			}
		}
	}

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string

		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {

			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
				return err
			}
		}
	}

	if o.MinQuantity != nil {

		// query param min_quantity
		var qrMinQuantity string

		if o.MinQuantity != nil {
			qrMinQuantity = *o.MinQuantity
		}
		qMinQuantity := qrMinQuantity
		if qMinQuantity != "" {

			if err := r.SetQueryParam("min_quantity", qMinQuantity); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.TokenAddress != nil {

		// query param token_address
		var qrTokenAddress string

		if o.TokenAddress != nil {
			qrTokenAddress = *o.TokenAddress
		}
		qTokenAddress := qrTokenAddress
		if qTokenAddress != "" {

			if err := r.SetQueryParam("token_address", qTokenAddress); err != nil {
				return err
			}
		}
	}

	if o.TokenID != nil {

		// query param token_id
		var qrTokenID string

		if o.TokenID != nil {
			qrTokenID = *o.TokenID
		}
		qTokenID := qrTokenID
		if qTokenID != "" {

			if err := r.SetQueryParam("token_id", qTokenID); err != nil {
				return err
			}
		}
	}

	if o.TokenName != nil {

		// query param token_name
		var qrTokenName string

		if o.TokenName != nil {
			qrTokenName = *o.TokenName
		}
		qTokenName := qrTokenName
		if qTokenName != "" {

			if err := r.SetQueryParam("token_name", qTokenName); err != nil {
				return err
			}
		}
	}

	if o.TokenType != nil {

		// query param token_type
		var qrTokenType string

		if o.TokenType != nil {
			qrTokenType = *o.TokenType
		}
		qTokenType := qrTokenType
		if qTokenType != "" {

			if err := r.SetQueryParam("token_type", qTokenType); err != nil {
				return err
			}
		}
	}

	if o.UpdatedMaxTimestamp != nil {

		// query param updated_max_timestamp
		var qrUpdatedMaxTimestamp string

		if o.UpdatedMaxTimestamp != nil {
			qrUpdatedMaxTimestamp = *o.UpdatedMaxTimestamp
		}
		qUpdatedMaxTimestamp := qrUpdatedMaxTimestamp
		if qUpdatedMaxTimestamp != "" {

			if err := r.SetQueryParam("updated_max_timestamp", qUpdatedMaxTimestamp); err != nil {
				return err
			}
		}
	}

	if o.UpdatedMinTimestamp != nil {

		// query param updated_min_timestamp
		var qrUpdatedMinTimestamp string

		if o.UpdatedMinTimestamp != nil {
			qrUpdatedMinTimestamp = *o.UpdatedMinTimestamp
		}
		qUpdatedMinTimestamp := qrUpdatedMinTimestamp
		if qUpdatedMinTimestamp != "" {

			if err := r.SetQueryParam("updated_min_timestamp", qUpdatedMinTimestamp); err != nil {
				return err
			}
		}
	}

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
