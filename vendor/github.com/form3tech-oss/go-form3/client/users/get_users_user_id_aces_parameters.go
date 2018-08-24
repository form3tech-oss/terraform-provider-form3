// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersUserIDAcesParams creates a new GetUsersUserIDAcesParams object
// with the default values initialized.
func NewGetUsersUserIDAcesParams() *GetUsersUserIDAcesParams {
	var ()
	return &GetUsersUserIDAcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDAcesParamsWithTimeout creates a new GetUsersUserIDAcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUserIDAcesParamsWithTimeout(timeout time.Duration) *GetUsersUserIDAcesParams {
	var ()
	return &GetUsersUserIDAcesParams{

		timeout: timeout,
	}
}

// NewGetUsersUserIDAcesParamsWithContext creates a new GetUsersUserIDAcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUserIDAcesParamsWithContext(ctx context.Context) *GetUsersUserIDAcesParams {
	var ()
	return &GetUsersUserIDAcesParams{

		Context: ctx,
	}
}

// NewGetUsersUserIDAcesParamsWithHTTPClient creates a new GetUsersUserIDAcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUserIDAcesParamsWithHTTPClient(client *http.Client) *GetUsersUserIDAcesParams {
	var ()
	return &GetUsersUserIDAcesParams{
		HTTPClient: client,
	}
}

/*GetUsersUserIDAcesParams contains all the parameters to send to the API endpoint
for the get users user ID aces operation typically these are written to a http.Request
*/
type GetUsersUserIDAcesParams struct {

	/*FilterAction
	  Access action

	*/
	FilterAction *string
	/*FilterRecordType
	  Record type

	*/
	FilterRecordType *string
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithTimeout(timeout time.Duration) *GetUsersUserIDAcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithContext(ctx context.Context) *GetUsersUserIDAcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithHTTPClient(client *http.Client) *GetUsersUserIDAcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterAction adds the filterAction to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithFilterAction(filterAction *string) *GetUsersUserIDAcesParams {
	o.SetFilterAction(filterAction)
	return o
}

// SetFilterAction adds the filterAction to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetFilterAction(filterAction *string) {
	o.FilterAction = filterAction
}

// WithFilterRecordType adds the filterRecordType to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithFilterRecordType(filterRecordType *string) *GetUsersUserIDAcesParams {
	o.SetFilterRecordType(filterRecordType)
	return o
}

// SetFilterRecordType adds the filterRecordType to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetFilterRecordType(filterRecordType *string) {
	o.FilterRecordType = filterRecordType
}

// WithUserID adds the userID to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) WithUserID(userID strfmt.UUID) *GetUsersUserIDAcesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID aces params
func (o *GetUsersUserIDAcesParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDAcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAction != nil {

		// query param filter[action]
		var qrFilterAction string
		if o.FilterAction != nil {
			qrFilterAction = *o.FilterAction
		}
		qFilterAction := qrFilterAction
		if qFilterAction != "" {
			if err := r.SetQueryParam("filter[action]", qFilterAction); err != nil {
				return err
			}
		}

	}

	if o.FilterRecordType != nil {

		// query param filter[record_type]
		var qrFilterRecordType string
		if o.FilterRecordType != nil {
			qrFilterRecordType = *o.FilterRecordType
		}
		qFilterRecordType := qrFilterRecordType
		if qFilterRecordType != "" {
			if err := r.SetQueryParam("filter[record_type]", qFilterRecordType); err != nil {
				return err
			}
		}

	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
