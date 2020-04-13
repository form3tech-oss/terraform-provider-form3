// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteKeysKeyID deletes key
*/
func (a *Client) DeleteKeysKeyID(params *DeleteKeysKeyIDParams) (*DeleteKeysKeyIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeysKeyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteKeysKeyID",
		Method:             "DELETE",
		PathPattern:        "/keys/{key_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKeysKeyIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKeysKeyIDNoContent), nil

}

/*
DeleteKeysKeyIDCertificatesCertificateID deletes certificate
*/
func (a *Client) DeleteKeysKeyIDCertificatesCertificateID(params *DeleteKeysKeyIDCertificatesCertificateIDParams) (*DeleteKeysKeyIDCertificatesCertificateIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeysKeyIDCertificatesCertificateIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteKeysKeyIDCertificatesCertificateID",
		Method:             "DELETE",
		PathPattern:        "/keys/{key_id}/certificates/{certificate_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKeysKeyIDCertificatesCertificateIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKeysKeyIDCertificatesCertificateIDNoContent), nil

}

/*
GetKeys lists all keys
*/
func (a *Client) GetKeys(params *GetKeysParams) (*GetKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetKeys",
		Method:             "GET",
		PathPattern:        "/keys",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKeysOK), nil

}

/*
GetKeysKeyID fetches key
*/
func (a *Client) GetKeysKeyID(params *GetKeysKeyIDParams) (*GetKeysKeyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeysKeyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetKeysKeyID",
		Method:             "GET",
		PathPattern:        "/keys/{key_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeysKeyIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKeysKeyIDOK), nil

}

/*
GetKeysKeyIDCertificates lists all certificates
*/
func (a *Client) GetKeysKeyIDCertificates(params *GetKeysKeyIDCertificatesParams) (*GetKeysKeyIDCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeysKeyIDCertificatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetKeysKeyIDCertificates",
		Method:             "GET",
		PathPattern:        "/keys/{key_id}/certificates",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeysKeyIDCertificatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKeysKeyIDCertificatesOK), nil

}

/*
GetKeysKeyIDCertificatesCertificateID fetches certificate
*/
func (a *Client) GetKeysKeyIDCertificatesCertificateID(params *GetKeysKeyIDCertificatesCertificateIDParams) (*GetKeysKeyIDCertificatesCertificateIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeysKeyIDCertificatesCertificateIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetKeysKeyIDCertificatesCertificateID",
		Method:             "GET",
		PathPattern:        "/keys/{key_id}/certificates/{certificate_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeysKeyIDCertificatesCertificateIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKeysKeyIDCertificatesCertificateIDOK), nil

}

/*
PostKeys creates key
*/
func (a *Client) PostKeys(params *PostKeysParams) (*PostKeysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostKeys",
		Method:             "POST",
		PathPattern:        "/keys",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostKeysCreated), nil

}

/*
PostKeysKeyIDCertificates creates certificate
*/
func (a *Client) PostKeysKeyIDCertificates(params *PostKeysKeyIDCertificatesParams) (*PostKeysKeyIDCertificatesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostKeysKeyIDCertificatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostKeysKeyIDCertificates",
		Method:             "POST",
		PathPattern:        "/keys/{key_id}/certificates",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostKeysKeyIDCertificatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostKeysKeyIDCertificatesCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
