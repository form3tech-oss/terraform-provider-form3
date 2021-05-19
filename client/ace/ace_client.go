// Code generated by go-swagger; DO NOT EDIT.

package ace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteRolesRoleIDAcesAceID(params *DeleteRolesRoleIDAcesAceIDParams, opts ...ClientOption) (*DeleteRolesRoleIDAcesAceIDNoContent, error)

	GetRolesRoleIDAces(params *GetRolesRoleIDAcesParams, opts ...ClientOption) (*GetRolesRoleIDAcesOK, error)

	GetRolesRoleIDAcesAceID(params *GetRolesRoleIDAcesAceIDParams, opts ...ClientOption) (*GetRolesRoleIDAcesAceIDOK, error)

	PostRolesRoleIDAces(params *PostRolesRoleIDAcesParams, opts ...ClientOption) (*PostRolesRoleIDAcesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteRolesRoleIDAcesAceID deletes access control entry
*/
func (a *Client) DeleteRolesRoleIDAcesAceID(params *DeleteRolesRoleIDAcesAceIDParams, opts ...ClientOption) (*DeleteRolesRoleIDAcesAceIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRolesRoleIDAcesAceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRolesRoleIDAcesAceID",
		Method:             "DELETE",
		PathPattern:        "/roles/{role_id}/aces/{ace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRolesRoleIDAcesAceIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteRolesRoleIDAcesAceIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteRolesRoleIDAcesAceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRolesRoleIDAces lists all access controls for role
*/
func (a *Client) GetRolesRoleIDAces(params *GetRolesRoleIDAcesParams, opts ...ClientOption) (*GetRolesRoleIDAcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesRoleIDAcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRolesRoleIDAces",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/aces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRolesRoleIDAcesReader{formats: a.formats},
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
	success, ok := result.(*GetRolesRoleIDAcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRolesRoleIDAces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRolesRoleIDAcesAceID fetches access control entry
*/
func (a *Client) GetRolesRoleIDAcesAceID(params *GetRolesRoleIDAcesAceIDParams, opts ...ClientOption) (*GetRolesRoleIDAcesAceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesRoleIDAcesAceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRolesRoleIDAcesAceID",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/aces/{ace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRolesRoleIDAcesAceIDReader{formats: a.formats},
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
	success, ok := result.(*GetRolesRoleIDAcesAceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRolesRoleIDAcesAceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostRolesRoleIDAces creates access control entry
*/
func (a *Client) PostRolesRoleIDAces(params *PostRolesRoleIDAcesParams, opts ...ClientOption) (*PostRolesRoleIDAcesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRolesRoleIDAcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostRolesRoleIDAces",
		Method:             "POST",
		PathPattern:        "/roles/{role_id}/aces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRolesRoleIDAcesReader{formats: a.formats},
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
	success, ok := result.(*PostRolesRoleIDAcesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostRolesRoleIDAces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
