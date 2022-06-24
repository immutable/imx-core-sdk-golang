// Code generated by go-swagger; DO NOT EDIT.

package encoding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new encoding API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for encoding API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EncodeAsset(params *EncodeAssetParams, opts ...ClientOption) (*EncodeAssetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EncodeAsset retrieves the starkex encoded format for a given asset

  Retrieves the Starkex Encoded format for a given asset so that it can be used as parameter for Starkex smart contracts
*/
func (a *Client) EncodeAsset(params *EncodeAssetParams, opts ...ClientOption) (*EncodeAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEncodeAssetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "encodeAsset",
		Method:             "POST",
		PathPattern:        "/v1/encode/{assetType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EncodeAssetReader{formats: a.formats},
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
	success, ok := result.(*EncodeAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for encodeAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}