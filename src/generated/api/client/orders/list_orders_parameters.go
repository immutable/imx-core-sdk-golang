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
	"github.com/go-openapi/swag"
)

// NewListOrdersParams creates a new ListOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOrdersParams() *ListOrdersParams {
	return &ListOrdersParams{
		timeout:                cr.DefaultTimeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListOrdersParamsWithTimeout creates a new ListOrdersParams object
// with the ability to set a timeout on a request.
func NewListOrdersParamsWithTimeout(timeout time.Duration) *ListOrdersParams {
	return &ListOrdersParams{
		timeout:                timeout,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListOrdersParamsWithContext creates a new ListOrdersParams object
// with the ability to set a context for a request.
func NewListOrdersParamsWithContext(ctx context.Context) *ListOrdersParams {
	return &ListOrdersParams{
		Context:                ctx,
		AdditionalHeaderParams: make(map[string]string),
	}
}

// NewListOrdersParamsWithHTTPClient creates a new ListOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOrdersParamsWithHTTPClient(client *http.Client) *ListOrdersParams {
	return &ListOrdersParams{
		HTTPClient:             client,
		AdditionalHeaderParams: make(map[string]string),
	}
}

/* ListOrdersParams contains all the parameters to send to the API endpoint
   for the list orders operation.

   Typically these are written to a http.Request.
*/
type ListOrdersParams struct {

	/* AuxiliaryFeePercentages.

	   Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients
	*/
	AuxiliaryFeePercentages *string

	/* AuxiliaryFeeRecipients.

	   Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages
	*/
	AuxiliaryFeeRecipients *string

	/* BuyAssetID.

	   Internal IMX ID of the asset this order buys
	*/
	BuyAssetID *string

	/* BuyMaxQuantity.

	   Max quantity for the asset this order buys
	*/
	BuyMaxQuantity *string

	/* BuyMetadata.

	   JSON-encoded metadata filters for the asset this order buys
	*/
	BuyMetadata *string

	/* BuyMinQuantity.

	   Min quantity for the asset this order buys
	*/
	BuyMinQuantity *string

	/* BuyTokenAddress.

	   Comma separated string of token addresses of the asset this order buys
	*/
	BuyTokenAddress *string

	/* BuyTokenID.

	   ERC721 Token ID of the asset this order buys
	*/
	BuyTokenID *string

	/* BuyTokenName.

	   Token name of the asset this order buys
	*/
	BuyTokenName *string

	/* BuyTokenType.

	   Token type of the asset this order buys
	*/
	BuyTokenType *string

	/* Cursor.

	   Cursor
	*/
	Cursor *string

	/* Direction.

	   Direction to sort (asc/desc)
	*/
	Direction *string

	/* MaxTimestamp.

	   Maximum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	MaxTimestamp *string

	/* MinTimestamp.

	   Minimum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	MinTimestamp *string

	/* OrderBy.

	   Property to sort by
	*/
	OrderBy *string

	/* PageSize.

	   Page size of the result
	*/
	PageSize *int64

	/* SellAssetID.

	   Internal IMX ID of the asset this order sells
	*/
	SellAssetID *string

	/* SellMaxQuantity.

	   Max quantity for the asset this order sells
	*/
	SellMaxQuantity *string

	/* SellMetadata.

	   JSON-encoded metadata filters for the asset this order sells
	*/
	SellMetadata *string

	/* SellMinQuantity.

	   Min quantity for the asset this order sells
	*/
	SellMinQuantity *string

	/* SellTokenAddress.

	   Comma separated string of token addresses of the asset this order sells
	*/
	SellTokenAddress *string

	/* SellTokenID.

	   ERC721 Token ID of the asset this order sells
	*/
	SellTokenID *string

	/* SellTokenName.

	   Token name of the asset this order sells
	*/
	SellTokenName *string

	/* SellTokenType.

	   Token type of the asset this order sells
	*/
	SellTokenType *string

	/* Status.

	   Status of this order
	*/
	Status *string

	/* UpdatedMaxTimestamp.

	   Maximum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	UpdatedMaxTimestamp *string

	/* UpdatedMinTimestamp.

	   Minimum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'
	*/
	UpdatedMinTimestamp *string

	/* User.

	   Ethereum address of the user who submitted this order
	*/
	User *string

	AdditionalHeaderParams map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrdersParams) WithDefaults() *ListOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrdersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) WithTimeout(timeout time.Duration) *ListOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list orders params
func (o *ListOrdersParams) WithContext(ctx context.Context) *ListOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list orders params
func (o *ListOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// AddCustomHeader provides option to add custom header parameters to list orders params.
func (o *ListOrdersParams) AddCustomHeader(key string, value string) {
	o.AdditionalHeaderParams[key] = value
}

// WithHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) WithHTTPClient(client *http.Client) *ListOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuxiliaryFeePercentages adds the auxiliaryFeePercentages to the list orders params
func (o *ListOrdersParams) WithAuxiliaryFeePercentages(auxiliaryFeePercentages *string) *ListOrdersParams {
	o.SetAuxiliaryFeePercentages(auxiliaryFeePercentages)
	return o
}

// SetAuxiliaryFeePercentages adds the auxiliaryFeePercentages to the list orders params
func (o *ListOrdersParams) SetAuxiliaryFeePercentages(auxiliaryFeePercentages *string) {
	o.AuxiliaryFeePercentages = auxiliaryFeePercentages
}

// WithAuxiliaryFeeRecipients adds the auxiliaryFeeRecipients to the list orders params
func (o *ListOrdersParams) WithAuxiliaryFeeRecipients(auxiliaryFeeRecipients *string) *ListOrdersParams {
	o.SetAuxiliaryFeeRecipients(auxiliaryFeeRecipients)
	return o
}

// SetAuxiliaryFeeRecipients adds the auxiliaryFeeRecipients to the list orders params
func (o *ListOrdersParams) SetAuxiliaryFeeRecipients(auxiliaryFeeRecipients *string) {
	o.AuxiliaryFeeRecipients = auxiliaryFeeRecipients
}

// WithBuyAssetID adds the buyAssetID to the list orders params
func (o *ListOrdersParams) WithBuyAssetID(buyAssetID *string) *ListOrdersParams {
	o.SetBuyAssetID(buyAssetID)
	return o
}

// SetBuyAssetID adds the buyAssetId to the list orders params
func (o *ListOrdersParams) SetBuyAssetID(buyAssetID *string) {
	o.BuyAssetID = buyAssetID
}

// WithBuyMaxQuantity adds the buyMaxQuantity to the list orders params
func (o *ListOrdersParams) WithBuyMaxQuantity(buyMaxQuantity *string) *ListOrdersParams {
	o.SetBuyMaxQuantity(buyMaxQuantity)
	return o
}

// SetBuyMaxQuantity adds the buyMaxQuantity to the list orders params
func (o *ListOrdersParams) SetBuyMaxQuantity(buyMaxQuantity *string) {
	o.BuyMaxQuantity = buyMaxQuantity
}

// WithBuyMetadata adds the buyMetadata to the list orders params
func (o *ListOrdersParams) WithBuyMetadata(buyMetadata *string) *ListOrdersParams {
	o.SetBuyMetadata(buyMetadata)
	return o
}

// SetBuyMetadata adds the buyMetadata to the list orders params
func (o *ListOrdersParams) SetBuyMetadata(buyMetadata *string) {
	o.BuyMetadata = buyMetadata
}

// WithBuyMinQuantity adds the buyMinQuantity to the list orders params
func (o *ListOrdersParams) WithBuyMinQuantity(buyMinQuantity *string) *ListOrdersParams {
	o.SetBuyMinQuantity(buyMinQuantity)
	return o
}

// SetBuyMinQuantity adds the buyMinQuantity to the list orders params
func (o *ListOrdersParams) SetBuyMinQuantity(buyMinQuantity *string) {
	o.BuyMinQuantity = buyMinQuantity
}

// WithBuyTokenAddress adds the buyTokenAddress to the list orders params
func (o *ListOrdersParams) WithBuyTokenAddress(buyTokenAddress *string) *ListOrdersParams {
	o.SetBuyTokenAddress(buyTokenAddress)
	return o
}

// SetBuyTokenAddress adds the buyTokenAddress to the list orders params
func (o *ListOrdersParams) SetBuyTokenAddress(buyTokenAddress *string) {
	o.BuyTokenAddress = buyTokenAddress
}

// WithBuyTokenID adds the buyTokenID to the list orders params
func (o *ListOrdersParams) WithBuyTokenID(buyTokenID *string) *ListOrdersParams {
	o.SetBuyTokenID(buyTokenID)
	return o
}

// SetBuyTokenID adds the buyTokenId to the list orders params
func (o *ListOrdersParams) SetBuyTokenID(buyTokenID *string) {
	o.BuyTokenID = buyTokenID
}

// WithBuyTokenName adds the buyTokenName to the list orders params
func (o *ListOrdersParams) WithBuyTokenName(buyTokenName *string) *ListOrdersParams {
	o.SetBuyTokenName(buyTokenName)
	return o
}

// SetBuyTokenName adds the buyTokenName to the list orders params
func (o *ListOrdersParams) SetBuyTokenName(buyTokenName *string) {
	o.BuyTokenName = buyTokenName
}

// WithBuyTokenType adds the buyTokenType to the list orders params
func (o *ListOrdersParams) WithBuyTokenType(buyTokenType *string) *ListOrdersParams {
	o.SetBuyTokenType(buyTokenType)
	return o
}

// SetBuyTokenType adds the buyTokenType to the list orders params
func (o *ListOrdersParams) SetBuyTokenType(buyTokenType *string) {
	o.BuyTokenType = buyTokenType
}

// WithCursor adds the cursor to the list orders params
func (o *ListOrdersParams) WithCursor(cursor *string) *ListOrdersParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list orders params
func (o *ListOrdersParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithDirection adds the direction to the list orders params
func (o *ListOrdersParams) WithDirection(direction *string) *ListOrdersParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the list orders params
func (o *ListOrdersParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithMaxTimestamp adds the maxTimestamp to the list orders params
func (o *ListOrdersParams) WithMaxTimestamp(maxTimestamp *string) *ListOrdersParams {
	o.SetMaxTimestamp(maxTimestamp)
	return o
}

// SetMaxTimestamp adds the maxTimestamp to the list orders params
func (o *ListOrdersParams) SetMaxTimestamp(maxTimestamp *string) {
	o.MaxTimestamp = maxTimestamp
}

// WithMinTimestamp adds the minTimestamp to the list orders params
func (o *ListOrdersParams) WithMinTimestamp(minTimestamp *string) *ListOrdersParams {
	o.SetMinTimestamp(minTimestamp)
	return o
}

// SetMinTimestamp adds the minTimestamp to the list orders params
func (o *ListOrdersParams) SetMinTimestamp(minTimestamp *string) {
	o.MinTimestamp = minTimestamp
}

// WithOrderBy adds the orderBy to the list orders params
func (o *ListOrdersParams) WithOrderBy(orderBy *string) *ListOrdersParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list orders params
func (o *ListOrdersParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPageSize adds the pageSize to the list orders params
func (o *ListOrdersParams) WithPageSize(pageSize *int64) *ListOrdersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list orders params
func (o *ListOrdersParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSellAssetID adds the sellAssetID to the list orders params
func (o *ListOrdersParams) WithSellAssetID(sellAssetID *string) *ListOrdersParams {
	o.SetSellAssetID(sellAssetID)
	return o
}

// SetSellAssetID adds the sellAssetId to the list orders params
func (o *ListOrdersParams) SetSellAssetID(sellAssetID *string) {
	o.SellAssetID = sellAssetID
}

// WithSellMaxQuantity adds the sellMaxQuantity to the list orders params
func (o *ListOrdersParams) WithSellMaxQuantity(sellMaxQuantity *string) *ListOrdersParams {
	o.SetSellMaxQuantity(sellMaxQuantity)
	return o
}

// SetSellMaxQuantity adds the sellMaxQuantity to the list orders params
func (o *ListOrdersParams) SetSellMaxQuantity(sellMaxQuantity *string) {
	o.SellMaxQuantity = sellMaxQuantity
}

// WithSellMetadata adds the sellMetadata to the list orders params
func (o *ListOrdersParams) WithSellMetadata(sellMetadata *string) *ListOrdersParams {
	o.SetSellMetadata(sellMetadata)
	return o
}

// SetSellMetadata adds the sellMetadata to the list orders params
func (o *ListOrdersParams) SetSellMetadata(sellMetadata *string) {
	o.SellMetadata = sellMetadata
}

// WithSellMinQuantity adds the sellMinQuantity to the list orders params
func (o *ListOrdersParams) WithSellMinQuantity(sellMinQuantity *string) *ListOrdersParams {
	o.SetSellMinQuantity(sellMinQuantity)
	return o
}

// SetSellMinQuantity adds the sellMinQuantity to the list orders params
func (o *ListOrdersParams) SetSellMinQuantity(sellMinQuantity *string) {
	o.SellMinQuantity = sellMinQuantity
}

// WithSellTokenAddress adds the sellTokenAddress to the list orders params
func (o *ListOrdersParams) WithSellTokenAddress(sellTokenAddress *string) *ListOrdersParams {
	o.SetSellTokenAddress(sellTokenAddress)
	return o
}

// SetSellTokenAddress adds the sellTokenAddress to the list orders params
func (o *ListOrdersParams) SetSellTokenAddress(sellTokenAddress *string) {
	o.SellTokenAddress = sellTokenAddress
}

// WithSellTokenID adds the sellTokenID to the list orders params
func (o *ListOrdersParams) WithSellTokenID(sellTokenID *string) *ListOrdersParams {
	o.SetSellTokenID(sellTokenID)
	return o
}

// SetSellTokenID adds the sellTokenId to the list orders params
func (o *ListOrdersParams) SetSellTokenID(sellTokenID *string) {
	o.SellTokenID = sellTokenID
}

// WithSellTokenName adds the sellTokenName to the list orders params
func (o *ListOrdersParams) WithSellTokenName(sellTokenName *string) *ListOrdersParams {
	o.SetSellTokenName(sellTokenName)
	return o
}

// SetSellTokenName adds the sellTokenName to the list orders params
func (o *ListOrdersParams) SetSellTokenName(sellTokenName *string) {
	o.SellTokenName = sellTokenName
}

// WithSellTokenType adds the sellTokenType to the list orders params
func (o *ListOrdersParams) WithSellTokenType(sellTokenType *string) *ListOrdersParams {
	o.SetSellTokenType(sellTokenType)
	return o
}

// SetSellTokenType adds the sellTokenType to the list orders params
func (o *ListOrdersParams) SetSellTokenType(sellTokenType *string) {
	o.SellTokenType = sellTokenType
}

// WithStatus adds the status to the list orders params
func (o *ListOrdersParams) WithStatus(status *string) *ListOrdersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list orders params
func (o *ListOrdersParams) SetStatus(status *string) {
	o.Status = status
}

// WithUpdatedMaxTimestamp adds the updatedMaxTimestamp to the list orders params
func (o *ListOrdersParams) WithUpdatedMaxTimestamp(updatedMaxTimestamp *string) *ListOrdersParams {
	o.SetUpdatedMaxTimestamp(updatedMaxTimestamp)
	return o
}

// SetUpdatedMaxTimestamp adds the updatedMaxTimestamp to the list orders params
func (o *ListOrdersParams) SetUpdatedMaxTimestamp(updatedMaxTimestamp *string) {
	o.UpdatedMaxTimestamp = updatedMaxTimestamp
}

// WithUpdatedMinTimestamp adds the updatedMinTimestamp to the list orders params
func (o *ListOrdersParams) WithUpdatedMinTimestamp(updatedMinTimestamp *string) *ListOrdersParams {
	o.SetUpdatedMinTimestamp(updatedMinTimestamp)
	return o
}

// SetUpdatedMinTimestamp adds the updatedMinTimestamp to the list orders params
func (o *ListOrdersParams) SetUpdatedMinTimestamp(updatedMinTimestamp *string) {
	o.UpdatedMinTimestamp = updatedMinTimestamp
}

// WithUser adds the user to the list orders params
func (o *ListOrdersParams) WithUser(user *string) *ListOrdersParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list orders params
func (o *ListOrdersParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AuxiliaryFeePercentages != nil {

		// query param auxiliary_fee_percentages
		var qrAuxiliaryFeePercentages string

		if o.AuxiliaryFeePercentages != nil {
			qrAuxiliaryFeePercentages = *o.AuxiliaryFeePercentages
		}
		qAuxiliaryFeePercentages := qrAuxiliaryFeePercentages
		if qAuxiliaryFeePercentages != "" {

			if err := r.SetQueryParam("auxiliary_fee_percentages", qAuxiliaryFeePercentages); err != nil {
				return err
			}
		}
	}

	if o.AuxiliaryFeeRecipients != nil {

		// query param auxiliary_fee_recipients
		var qrAuxiliaryFeeRecipients string

		if o.AuxiliaryFeeRecipients != nil {
			qrAuxiliaryFeeRecipients = *o.AuxiliaryFeeRecipients
		}
		qAuxiliaryFeeRecipients := qrAuxiliaryFeeRecipients
		if qAuxiliaryFeeRecipients != "" {

			if err := r.SetQueryParam("auxiliary_fee_recipients", qAuxiliaryFeeRecipients); err != nil {
				return err
			}
		}
	}

	if o.BuyAssetID != nil {

		// query param buy_asset_id
		var qrBuyAssetID string

		if o.BuyAssetID != nil {
			qrBuyAssetID = *o.BuyAssetID
		}
		qBuyAssetID := qrBuyAssetID
		if qBuyAssetID != "" {

			if err := r.SetQueryParam("buy_asset_id", qBuyAssetID); err != nil {
				return err
			}
		}
	}

	if o.BuyMaxQuantity != nil {

		// query param buy_max_quantity
		var qrBuyMaxQuantity string

		if o.BuyMaxQuantity != nil {
			qrBuyMaxQuantity = *o.BuyMaxQuantity
		}
		qBuyMaxQuantity := qrBuyMaxQuantity
		if qBuyMaxQuantity != "" {

			if err := r.SetQueryParam("buy_max_quantity", qBuyMaxQuantity); err != nil {
				return err
			}
		}
	}

	if o.BuyMetadata != nil {

		// query param buy_metadata
		var qrBuyMetadata string

		if o.BuyMetadata != nil {
			qrBuyMetadata = *o.BuyMetadata
		}
		qBuyMetadata := qrBuyMetadata
		if qBuyMetadata != "" {

			if err := r.SetQueryParam("buy_metadata", qBuyMetadata); err != nil {
				return err
			}
		}
	}

	if o.BuyMinQuantity != nil {

		// query param buy_min_quantity
		var qrBuyMinQuantity string

		if o.BuyMinQuantity != nil {
			qrBuyMinQuantity = *o.BuyMinQuantity
		}
		qBuyMinQuantity := qrBuyMinQuantity
		if qBuyMinQuantity != "" {

			if err := r.SetQueryParam("buy_min_quantity", qBuyMinQuantity); err != nil {
				return err
			}
		}
	}

	if o.BuyTokenAddress != nil {

		// query param buy_token_address
		var qrBuyTokenAddress string

		if o.BuyTokenAddress != nil {
			qrBuyTokenAddress = *o.BuyTokenAddress
		}
		qBuyTokenAddress := qrBuyTokenAddress
		if qBuyTokenAddress != "" {

			if err := r.SetQueryParam("buy_token_address", qBuyTokenAddress); err != nil {
				return err
			}
		}
	}

	if o.BuyTokenID != nil {

		// query param buy_token_id
		var qrBuyTokenID string

		if o.BuyTokenID != nil {
			qrBuyTokenID = *o.BuyTokenID
		}
		qBuyTokenID := qrBuyTokenID
		if qBuyTokenID != "" {

			if err := r.SetQueryParam("buy_token_id", qBuyTokenID); err != nil {
				return err
			}
		}
	}

	if o.BuyTokenName != nil {

		// query param buy_token_name
		var qrBuyTokenName string

		if o.BuyTokenName != nil {
			qrBuyTokenName = *o.BuyTokenName
		}
		qBuyTokenName := qrBuyTokenName
		if qBuyTokenName != "" {

			if err := r.SetQueryParam("buy_token_name", qBuyTokenName); err != nil {
				return err
			}
		}
	}

	if o.BuyTokenType != nil {

		// query param buy_token_type
		var qrBuyTokenType string

		if o.BuyTokenType != nil {
			qrBuyTokenType = *o.BuyTokenType
		}
		qBuyTokenType := qrBuyTokenType
		if qBuyTokenType != "" {

			if err := r.SetQueryParam("buy_token_type", qBuyTokenType); err != nil {
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

	if o.MaxTimestamp != nil {

		// query param max_timestamp
		var qrMaxTimestamp string

		if o.MaxTimestamp != nil {
			qrMaxTimestamp = *o.MaxTimestamp
		}
		qMaxTimestamp := qrMaxTimestamp
		if qMaxTimestamp != "" {

			if err := r.SetQueryParam("max_timestamp", qMaxTimestamp); err != nil {
				return err
			}
		}
	}

	if o.MinTimestamp != nil {

		// query param min_timestamp
		var qrMinTimestamp string

		if o.MinTimestamp != nil {
			qrMinTimestamp = *o.MinTimestamp
		}
		qMinTimestamp := qrMinTimestamp
		if qMinTimestamp != "" {

			if err := r.SetQueryParam("min_timestamp", qMinTimestamp); err != nil {
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

	if o.SellAssetID != nil {

		// query param sell_asset_id
		var qrSellAssetID string

		if o.SellAssetID != nil {
			qrSellAssetID = *o.SellAssetID
		}
		qSellAssetID := qrSellAssetID
		if qSellAssetID != "" {

			if err := r.SetQueryParam("sell_asset_id", qSellAssetID); err != nil {
				return err
			}
		}
	}

	if o.SellMaxQuantity != nil {

		// query param sell_max_quantity
		var qrSellMaxQuantity string

		if o.SellMaxQuantity != nil {
			qrSellMaxQuantity = *o.SellMaxQuantity
		}
		qSellMaxQuantity := qrSellMaxQuantity
		if qSellMaxQuantity != "" {

			if err := r.SetQueryParam("sell_max_quantity", qSellMaxQuantity); err != nil {
				return err
			}
		}
	}

	if o.SellMetadata != nil {

		// query param sell_metadata
		var qrSellMetadata string

		if o.SellMetadata != nil {
			qrSellMetadata = *o.SellMetadata
		}
		qSellMetadata := qrSellMetadata
		if qSellMetadata != "" {

			if err := r.SetQueryParam("sell_metadata", qSellMetadata); err != nil {
				return err
			}
		}
	}

	if o.SellMinQuantity != nil {

		// query param sell_min_quantity
		var qrSellMinQuantity string

		if o.SellMinQuantity != nil {
			qrSellMinQuantity = *o.SellMinQuantity
		}
		qSellMinQuantity := qrSellMinQuantity
		if qSellMinQuantity != "" {

			if err := r.SetQueryParam("sell_min_quantity", qSellMinQuantity); err != nil {
				return err
			}
		}
	}

	if o.SellTokenAddress != nil {

		// query param sell_token_address
		var qrSellTokenAddress string

		if o.SellTokenAddress != nil {
			qrSellTokenAddress = *o.SellTokenAddress
		}
		qSellTokenAddress := qrSellTokenAddress
		if qSellTokenAddress != "" {

			if err := r.SetQueryParam("sell_token_address", qSellTokenAddress); err != nil {
				return err
			}
		}
	}

	if o.SellTokenID != nil {

		// query param sell_token_id
		var qrSellTokenID string

		if o.SellTokenID != nil {
			qrSellTokenID = *o.SellTokenID
		}
		qSellTokenID := qrSellTokenID
		if qSellTokenID != "" {

			if err := r.SetQueryParam("sell_token_id", qSellTokenID); err != nil {
				return err
			}
		}
	}

	if o.SellTokenName != nil {

		// query param sell_token_name
		var qrSellTokenName string

		if o.SellTokenName != nil {
			qrSellTokenName = *o.SellTokenName
		}
		qSellTokenName := qrSellTokenName
		if qSellTokenName != "" {

			if err := r.SetQueryParam("sell_token_name", qSellTokenName); err != nil {
				return err
			}
		}
	}

	if o.SellTokenType != nil {

		// query param sell_token_type
		var qrSellTokenType string

		if o.SellTokenType != nil {
			qrSellTokenType = *o.SellTokenType
		}
		qSellTokenType := qrSellTokenType
		if qSellTokenType != "" {

			if err := r.SetQueryParam("sell_token_type", qSellTokenType); err != nil {
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
