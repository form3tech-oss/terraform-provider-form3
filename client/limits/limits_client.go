// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new limits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for limits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteLimitsID deletes limit
*/
func (a *Client) DeleteLimitsID(params *DeleteLimitsIDParams) (*DeleteLimitsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLimitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLimitsID",
		Method:             "DELETE",
		PathPattern:        "/limits/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLimitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLimitsIDNoContent), nil

}

/*
GetLimits lists limits
*/
func (a *Client) GetLimits(params *GetLimitsParams) (*GetLimitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLimitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLimits",
		Method:             "GET",
		PathPattern:        "/limits",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLimitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLimitsOK), nil

}

/*
GetLimitsID fetches limit
*/
func (a *Client) GetLimitsID(params *GetLimitsIDParams) (*GetLimitsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLimitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLimitsID",
		Method:             "GET",
		PathPattern:        "/limits/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLimitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLimitsIDOK), nil

}

/*
PatchLimitsID amends limit
*/
func (a *Client) PatchLimitsID(params *PatchLimitsIDParams) (*PatchLimitsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLimitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLimitsID",
		Method:             "PATCH",
		PathPattern:        "/limits/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLimitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLimitsIDOK), nil

}

/*
PostLimits creates a limit
*/
func (a *Client) PostLimits(params *PostLimitsParams) (*PostLimitsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLimitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLimits",
		Method:             "POST",
		PathPattern:        "/limits",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLimitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLimitsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}