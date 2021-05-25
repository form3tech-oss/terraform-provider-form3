// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewDeleteProductsIDParams creates a new DeleteProductsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProductsIDParams() *DeleteProductsIDParams {
	return &DeleteProductsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProductsIDParamsWithTimeout creates a new DeleteProductsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProductsIDParamsWithTimeout(timeout time.Duration) *DeleteProductsIDParams {
	return &DeleteProductsIDParams{
		timeout: timeout,
	}
}

// NewDeleteProductsIDParamsWithContext creates a new DeleteProductsIDParams object
// with the ability to set a context for a request.
func NewDeleteProductsIDParamsWithContext(ctx context.Context) *DeleteProductsIDParams {
	return &DeleteProductsIDParams{
		Context: ctx,
	}
}

// NewDeleteProductsIDParamsWithHTTPClient creates a new DeleteProductsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProductsIDParamsWithHTTPClient(client *http.Client) *DeleteProductsIDParams {
	return &DeleteProductsIDParams{
		HTTPClient: client,
	}
}

/* DeleteProductsIDParams contains all the parameters to send to the API endpoint
   for the delete products ID operation.

   Typically these are written to a http.Request.
*/
type DeleteProductsIDParams struct {

	/* ID.

	   Association Id

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Version.

	   Version
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete products ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProductsIDParams) WithDefaults() *DeleteProductsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete products ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProductsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete products ID params
func (o *DeleteProductsIDParams) WithTimeout(timeout time.Duration) *DeleteProductsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete products ID params
func (o *DeleteProductsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete products ID params
func (o *DeleteProductsIDParams) WithContext(ctx context.Context) *DeleteProductsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete products ID params
func (o *DeleteProductsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete products ID params
func (o *DeleteProductsIDParams) WithHTTPClient(client *http.Client) *DeleteProductsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete products ID params
func (o *DeleteProductsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete products ID params
func (o *DeleteProductsIDParams) WithID(id strfmt.UUID) *DeleteProductsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete products ID params
func (o *DeleteProductsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete products ID params
func (o *DeleteProductsIDParams) WithVersion(version int64) *DeleteProductsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete products ID params
func (o *DeleteProductsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProductsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
