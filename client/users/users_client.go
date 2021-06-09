// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUsersUserID(params *DeleteUsersUserIDParams) (*DeleteUsersUserIDNoContent, error)

	DeleteUsersUserIDCredentialsClientID(params *DeleteUsersUserIDCredentialsClientIDParams) (*DeleteUsersUserIDCredentialsClientIDNoContent, error)

	DeleteUsersUserIDCredentialsPublicKeyPublicKeyID(params *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) (*DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDNoContent, error)

	DeleteUsersUserIDCredentialsSsoSsoUserID(params *DeleteUsersUserIDCredentialsSsoSsoUserIDParams) (*DeleteUsersUserIDCredentialsSsoSsoUserIDNoContent, error)

	DeleteUsersUserIDRolesRoleID(params *DeleteUsersUserIDRolesRoleIDParams) (*DeleteUsersUserIDRolesRoleIDNoContent, error)

	GetUsers(params *GetUsersParams) (*GetUsersOK, error)

	GetUsersUserID(params *GetUsersUserIDParams) (*GetUsersUserIDOK, error)

	GetUsersUserIDAces(params *GetUsersUserIDAcesParams) (*GetUsersUserIDAcesOK, error)

	GetUsersUserIDCredentials(params *GetUsersUserIDCredentialsParams) (*GetUsersUserIDCredentialsOK, error)

	GetUsersUserIDCredentialsPublicKey(params *GetUsersUserIDCredentialsPublicKeyParams) (*GetUsersUserIDCredentialsPublicKeyOK, error)

	GetUsersUserIDCredentialsPublicKeyPublicKeyID(params *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) (*GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK, error)

	GetUsersUserIDCredentialsSsoSsoUserID(params *GetUsersUserIDCredentialsSsoSsoUserIDParams) (*GetUsersUserIDCredentialsSsoSsoUserIDOK, error)

	GetUsersUserIDRoles(params *GetUsersUserIDRolesParams) (*GetUsersUserIDRolesOK, error)

	PatchUsersUserID(params *PatchUsersUserIDParams) (*PatchUsersUserIDOK, error)

	PostUsers(params *PostUsersParams) (*PostUsersCreated, error)

	PostUsersUserIDCredentials(params *PostUsersUserIDCredentialsParams) (*PostUsersUserIDCredentialsCreated, error)

	PostUsersUserIDCredentialsPublicKey(params *PostUsersUserIDCredentialsPublicKeyParams) (*PostUsersUserIDCredentialsPublicKeyCreated, error)

	PostUsersUserIDCredentialsSso(params *PostUsersUserIDCredentialsSsoParams) (*PostUsersUserIDCredentialsSsoCreated, error)

	PostUsersUserIDRolesRoleID(params *PostUsersUserIDRolesRoleIDParams) (*PostUsersUserIDRolesRoleIDCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUsersUserID deletes user
*/
func (a *Client) DeleteUsersUserID(params *DeleteUsersUserIDParams) (*DeleteUsersUserIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersUserIDCredentialsClientID deletes credential for user
*/
func (a *Client) DeleteUsersUserIDCredentialsClientID(params *DeleteUsersUserIDCredentialsClientIDParams) (*DeleteUsersUserIDCredentialsClientIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDCredentialsClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserIDCredentialsClientID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/credentials/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDCredentialsClientIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDCredentialsClientIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserIDCredentialsClientID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersUserIDCredentialsPublicKeyPublicKeyID deletes public key credential for user
*/
func (a *Client) DeleteUsersUserIDCredentialsPublicKeyPublicKeyID(params *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) (*DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserIDCredentialsPublicKeyPublicKeyID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/credentials/public_key/{public_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserIDCredentialsPublicKeyPublicKeyID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersUserIDCredentialsSsoSsoUserID deletes sso user credential
*/
func (a *Client) DeleteUsersUserIDCredentialsSsoSsoUserID(params *DeleteUsersUserIDCredentialsSsoSsoUserIDParams) (*DeleteUsersUserIDCredentialsSsoSsoUserIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDCredentialsSsoSsoUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserIDCredentialsSsoSsoUserID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/credentials/sso/{sso_user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDCredentialsSsoSsoUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDCredentialsSsoSsoUserIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserIDCredentialsSsoSsoUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersUserIDRolesRoleID removes role from user
*/
func (a *Client) DeleteUsersUserIDRolesRoleID(params *DeleteUsersUserIDRolesRoleIDParams) (*DeleteUsersUserIDRolesRoleIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDRolesRoleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserIDRolesRoleID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDRolesRoleIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDRolesRoleIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserIDRolesRoleID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsers lists all users
*/
func (a *Client) GetUsers(params *GetUsersParams) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserID fetches user
*/
func (a *Client) GetUsersUserID(params *GetUsersUserIDParams) (*GetUsersUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserID",
		Method:             "GET",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDAces gets access control list for user
*/
func (a *Client) GetUsersUserIDAces(params *GetUsersUserIDAcesParams) (*GetUsersUserIDAcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDAcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDAces",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/aces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDAcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDAcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDAces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDCredentials gets all credentials for user
*/
func (a *Client) GetUsersUserIDCredentials(params *GetUsersUserIDCredentialsParams) (*GetUsersUserIDCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDCredentials",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDCredentialsPublicKey fetches public key credentials
*/
func (a *Client) GetUsersUserIDCredentialsPublicKey(params *GetUsersUserIDCredentialsPublicKeyParams) (*GetUsersUserIDCredentialsPublicKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDCredentialsPublicKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDCredentialsPublicKey",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/credentials/public_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDCredentialsPublicKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDCredentialsPublicKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDCredentialsPublicKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDCredentialsPublicKeyPublicKeyID fetches public key credential
*/
func (a *Client) GetUsersUserIDCredentialsPublicKeyPublicKeyID(params *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) (*GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDCredentialsPublicKeyPublicKeyID",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/credentials/public_key/{public_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDCredentialsPublicKeyPublicKeyIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDCredentialsPublicKeyPublicKeyID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDCredentialsSsoSsoUserID fetches sso credential
*/
func (a *Client) GetUsersUserIDCredentialsSsoSsoUserID(params *GetUsersUserIDCredentialsSsoSsoUserIDParams) (*GetUsersUserIDCredentialsSsoSsoUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDCredentialsSsoSsoUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDCredentialsSsoSsoUserID",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/credentials/sso/{sso_user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDCredentialsSsoSsoUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDCredentialsSsoSsoUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDCredentialsSsoSsoUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUserIDRoles gets all roles for user
*/
func (a *Client) GetUsersUserIDRoles(params *GetUsersUserIDRolesParams) (*GetUsersUserIDRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDRoles",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchUsersUserID edits user details
*/
func (a *Client) PatchUsersUserID(params *PatchUsersUserIDParams) (*PatchUsersUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchUsersUserID",
		Method:             "PATCH",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUsersUserIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchUsersUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchUsersUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsers creates user
*/
func (a *Client) PostUsers(params *PostUsersParams) (*PostUsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsers",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersUserIDCredentials adds credentials to user
*/
func (a *Client) PostUsersUserIDCredentials(params *PostUsersUserIDCredentialsParams) (*PostUsersUserIDCredentialsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersUserIDCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsersUserIDCredentials",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersUserIDCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersUserIDCredentialsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsersUserIDCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersUserIDCredentialsPublicKey generates new public key credential for a user
*/
func (a *Client) PostUsersUserIDCredentialsPublicKey(params *PostUsersUserIDCredentialsPublicKeyParams) (*PostUsersUserIDCredentialsPublicKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersUserIDCredentialsPublicKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsersUserIDCredentialsPublicKey",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/credentials/public_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersUserIDCredentialsPublicKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersUserIDCredentialsPublicKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsersUserIDCredentialsPublicKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersUserIDCredentialsSso creates new sso credential for a user
*/
func (a *Client) PostUsersUserIDCredentialsSso(params *PostUsersUserIDCredentialsSsoParams) (*PostUsersUserIDCredentialsSsoCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersUserIDCredentialsSsoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsersUserIDCredentialsSso",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/credentials/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersUserIDCredentialsSsoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersUserIDCredentialsSsoCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsersUserIDCredentialsSso: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersUserIDRolesRoleID adds role to user
*/
func (a *Client) PostUsersUserIDRolesRoleID(params *PostUsersUserIDRolesRoleIDParams) (*PostUsersUserIDRolesRoleIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersUserIDRolesRoleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsersUserIDRolesRoleID",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersUserIDRolesRoleIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersUserIDRolesRoleIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsersUserIDRolesRoleID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
