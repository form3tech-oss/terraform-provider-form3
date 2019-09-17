// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new associations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for associations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteBacsID deletes organisation association for b a c s
*/
func (a *Client) DeleteBacsID(params *DeleteBacsIDParams) (*DeleteBacsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBacsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteBacsID",
		Method:             "DELETE",
		PathPattern:        "/bacs/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBacsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteBacsIDNoContent), nil

}

/*
DeleteConfirmationOfPayeeID deletes organisation association
*/
func (a *Client) DeleteConfirmationOfPayeeID(params *DeleteConfirmationOfPayeeIDParams) (*DeleteConfirmationOfPayeeIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConfirmationOfPayeeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteConfirmationOfPayeeID",
		Method:             "DELETE",
		PathPattern:        "/confirmation-of-payee/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConfirmationOfPayeeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteConfirmationOfPayeeIDNoContent), nil

}

/*
DeleteLhvID deletes organisation lhv association
*/
func (a *Client) DeleteLhvID(params *DeleteLhvIDParams) (*DeleteLhvIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLhvIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLhvID",
		Method:             "DELETE",
		PathPattern:        "/lhv/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLhvIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLhvIDNoContent), nil

}

/*
DeletePayportID deletes service association
*/
func (a *Client) DeletePayportID(params *DeletePayportIDParams) (*DeletePayportIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePayportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePayportID",
		Method:             "DELETE",
		PathPattern:        "/payport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePayportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePayportIDNoContent), nil

}

/*
DeleteProductsID deletes product association
*/
func (a *Client) DeleteProductsID(params *DeleteProductsIDParams) (*DeleteProductsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProductsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProductsID",
		Method:             "DELETE",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProductsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProductsIDNoContent), nil

}

/*
DeleteSepainstantID deletes organisation spea instant association
*/
func (a *Client) DeleteSepainstantID(params *DeleteSepainstantIDParams) (*DeleteSepainstantIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSepainstantIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSepainstantID",
		Method:             "DELETE",
		PathPattern:        "/sepainstant/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSepainstantIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSepainstantIDNoContent), nil

}

/*
DeleteSepasctID deletes organisation spea sct association
*/
func (a *Client) DeleteSepasctID(params *DeleteSepasctIDParams) (*DeleteSepasctIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSepasctIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSepasctID",
		Method:             "DELETE",
		PathPattern:        "/sepasct/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSepasctIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSepasctIDNoContent), nil

}

/*
DeleteStarlingID deletes organisation association
*/
func (a *Client) DeleteStarlingID(params *DeleteStarlingIDParams) (*DeleteStarlingIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStarlingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteStarlingID",
		Method:             "DELETE",
		PathPattern:        "/starling/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStarlingIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStarlingIDNoContent), nil

}

/*
DeleteVocalinkreportID deletes organisation association
*/
func (a *Client) DeleteVocalinkreportID(params *DeleteVocalinkreportIDParams) (*DeleteVocalinkreportIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVocalinkreportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVocalinkreportID",
		Method:             "DELETE",
		PathPattern:        "/vocalinkreport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVocalinkreportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVocalinkreportIDNoContent), nil

}

/*
GetBacs lists all organisation associations for b a c s
*/
func (a *Client) GetBacs(params *GetBacsParams) (*GetBacsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBacsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBacs",
		Method:             "GET",
		PathPattern:        "/bacs",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBacsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBacsOK), nil

}

/*
GetBacsID fetches organisation association for b a c s
*/
func (a *Client) GetBacsID(params *GetBacsIDParams) (*GetBacsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBacsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBacsID",
		Method:             "GET",
		PathPattern:        "/bacs/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBacsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBacsIDOK), nil

}

/*
GetConfirmationOfPayee lists all organisation associations
*/
func (a *Client) GetConfirmationOfPayee(params *GetConfirmationOfPayeeParams) (*GetConfirmationOfPayeeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfirmationOfPayeeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfirmationOfPayee",
		Method:             "GET",
		PathPattern:        "/confirmation-of-payee",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfirmationOfPayeeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfirmationOfPayeeOK), nil

}

/*
GetConfirmationOfPayeeID fetches organisation association
*/
func (a *Client) GetConfirmationOfPayeeID(params *GetConfirmationOfPayeeIDParams) (*GetConfirmationOfPayeeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfirmationOfPayeeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfirmationOfPayeeID",
		Method:             "GET",
		PathPattern:        "/confirmation-of-payee/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfirmationOfPayeeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfirmationOfPayeeIDOK), nil

}

/*
GetLhv lists all organisation lhv associations
*/
func (a *Client) GetLhv(params *GetLhvParams) (*GetLhvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLhvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLhv",
		Method:             "GET",
		PathPattern:        "/lhv",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLhvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLhvOK), nil

}

/*
GetLhvID fetches organisation lhv association
*/
func (a *Client) GetLhvID(params *GetLhvIDParams) (*GetLhvIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLhvIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLhvID",
		Method:             "GET",
		PathPattern:        "/lhv/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLhvIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLhvIDOK), nil

}

/*
GetPayport lists all organisation associations
*/
func (a *Client) GetPayport(params *GetPayportParams) (*GetPayportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayport",
		Method:             "GET",
		PathPattern:        "/payport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPayportOK), nil

}

/*
GetPayportID fetches service association
*/
func (a *Client) GetPayportID(params *GetPayportIDParams) (*GetPayportIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayportID",
		Method:             "GET",
		PathPattern:        "/payport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPayportIDOK), nil

}

/*
GetProducts lists all product associations
*/
func (a *Client) GetProducts(params *GetProductsParams) (*GetProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductsOK), nil

}

/*
GetProductsID fetches product association
*/
func (a *Client) GetProductsID(params *GetProductsIDParams) (*GetProductsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProductsID",
		Method:             "GET",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductsIDOK), nil

}

/*
GetSepainstant lists all organisation sepa instant associations
*/
func (a *Client) GetSepainstant(params *GetSepainstantParams) (*GetSepainstantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSepainstantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSepainstant",
		Method:             "GET",
		PathPattern:        "/sepainstant",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSepainstantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSepainstantOK), nil

}

/*
GetSepainstantID fetches organisation sepa instant association
*/
func (a *Client) GetSepainstantID(params *GetSepainstantIDParams) (*GetSepainstantIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSepainstantIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSepainstantID",
		Method:             "GET",
		PathPattern:        "/sepainstant/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSepainstantIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSepainstantIDOK), nil

}

/*
GetSepasct lists all organisation sepa sct associations
*/
func (a *Client) GetSepasct(params *GetSepasctParams) (*GetSepasctOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSepasctParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSepasct",
		Method:             "GET",
		PathPattern:        "/sepasct",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSepasctReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSepasctOK), nil

}

/*
GetSepasctID fetches organisation sepa sct association
*/
func (a *Client) GetSepasctID(params *GetSepasctIDParams) (*GetSepasctIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSepasctIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSepasctID",
		Method:             "GET",
		PathPattern:        "/sepasct/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSepasctIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSepasctIDOK), nil

}

/*
GetStarling lists all organisation associations
*/
func (a *Client) GetStarling(params *GetStarlingParams) (*GetStarlingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStarlingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStarling",
		Method:             "GET",
		PathPattern:        "/starling",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStarlingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStarlingOK), nil

}

/*
GetStarlingID fetches organisation association
*/
func (a *Client) GetStarlingID(params *GetStarlingIDParams) (*GetStarlingIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStarlingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStarlingID",
		Method:             "GET",
		PathPattern:        "/starling/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStarlingIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStarlingIDOK), nil

}

/*
GetVocalinkreport lists all organisation associations
*/
func (a *Client) GetVocalinkreport(params *GetVocalinkreportParams) (*GetVocalinkreportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVocalinkreportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVocalinkreport",
		Method:             "GET",
		PathPattern:        "/vocalinkreport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVocalinkreportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVocalinkreportOK), nil

}

/*
GetVocalinkreportID fetches organisation association
*/
func (a *Client) GetVocalinkreportID(params *GetVocalinkreportIDParams) (*GetVocalinkreportIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVocalinkreportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVocalinkreportID",
		Method:             "GET",
		PathPattern:        "/vocalinkreport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVocalinkreportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVocalinkreportIDOK), nil

}

/*
PostBacs creates organisation association for b a c s
*/
func (a *Client) PostBacs(params *PostBacsParams) (*PostBacsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBacsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostBacs",
		Method:             "POST",
		PathPattern:        "/bacs",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostBacsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostBacsCreated), nil

}

/*
PostConfirmationOfPayee creates organisation association with confirmation of payee
*/
func (a *Client) PostConfirmationOfPayee(params *PostConfirmationOfPayeeParams) (*PostConfirmationOfPayeeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostConfirmationOfPayeeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostConfirmationOfPayee",
		Method:             "POST",
		PathPattern:        "/confirmation-of-payee",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostConfirmationOfPayeeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostConfirmationOfPayeeCreated), nil

}

/*
PostLhv creates organisation association for lhv
*/
func (a *Client) PostLhv(params *PostLhvParams) (*PostLhvCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLhvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLhv",
		Method:             "POST",
		PathPattern:        "/lhv",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLhvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLhvCreated), nil

}

/*
PostPayport creates payport service association
*/
func (a *Client) PostPayport(params *PostPayportParams) (*PostPayportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPayportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPayport",
		Method:             "POST",
		PathPattern:        "/payport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPayportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPayportCreated), nil

}

/*
PostProducts creates product association
*/
func (a *Client) PostProducts(params *PostProductsParams) (*PostProductsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostProducts",
		Method:             "POST",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostProductsCreated), nil

}

/*
PostSepainstant creates organisation association for sepa instant
*/
func (a *Client) PostSepainstant(params *PostSepainstantParams) (*PostSepainstantCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSepainstantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSepainstant",
		Method:             "POST",
		PathPattern:        "/sepainstant",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSepainstantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSepainstantCreated), nil

}

/*
PostSepasct creates organisation association for sepa sct
*/
func (a *Client) PostSepasct(params *PostSepasctParams) (*PostSepasctCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSepasctParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSepasct",
		Method:             "POST",
		PathPattern:        "/sepasct",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSepasctReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSepasctCreated), nil

}

/*
PostStarling creates organisation association
*/
func (a *Client) PostStarling(params *PostStarlingParams) (*PostStarlingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStarlingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStarling",
		Method:             "POST",
		PathPattern:        "/starling",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStarlingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStarlingCreated), nil

}

/*
PostVocalinkreport creates organisation association
*/
func (a *Client) PostVocalinkreport(params *PostVocalinkreportParams) (*PostVocalinkreportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVocalinkreportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVocalinkreport",
		Method:             "POST",
		PathPattern:        "/vocalinkreport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVocalinkreportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostVocalinkreportCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
