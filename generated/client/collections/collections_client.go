// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCollection(params *CreateCollectionParams, opts ...ClientOption) (*CreateCollectionCreated, error)

	GetCollection(params *GetCollectionParams, opts ...ClientOption) (*GetCollectionOK, error)

	ListCollectionFilters(params *ListCollectionFiltersParams, opts ...ClientOption) (*ListCollectionFiltersOK, error)

	ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error)

	UpdateCollection(params *UpdateCollectionParams, opts ...ClientOption) (*UpdateCollectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCollection creates collection

  Create collection
*/
func (a *Client) CreateCollection(params *CreateCollectionParams, opts ...ClientOption) (*CreateCollectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCollection",
		Method:             "POST",
		PathPattern:        "/v1/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCollectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCollection gets details of a collection at the given address

  Get details of a collection at the given address
*/
func (a *Client) GetCollection(params *GetCollectionParams, opts ...ClientOption) (*GetCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCollection",
		Method:             "GET",
		PathPattern:        "/v1/collections/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCollectionFilters gets a list of collection filters

  Get a list of collection filters
*/
func (a *Client) ListCollectionFilters(params *ListCollectionFiltersParams, opts ...ClientOption) (*ListCollectionFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCollectionFiltersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listCollectionFilters",
		Method:             "GET",
		PathPattern:        "/v1/collections/{address}/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCollectionFiltersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCollectionFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCollectionFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCollections gets a list of collections

  Get a list of collections
*/
func (a *Client) ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listCollections",
		Method:             "GET",
		PathPattern:        "/v1/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCollectionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCollection updates collection

  Update collection
*/
func (a *Client) UpdateCollection(params *UpdateCollectionParams, opts ...ClientOption) (*UpdateCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCollection",
		Method:             "PATCH",
		PathPattern:        "/v1/collections/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
