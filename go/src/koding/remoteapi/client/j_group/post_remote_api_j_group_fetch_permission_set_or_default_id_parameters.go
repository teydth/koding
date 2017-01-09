package j_group

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams creates a new PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams() *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParamsWithTimeout creates a new PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParamsWithContext creates a new PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams contains all the parameters to send to the API endpoint
for the post remote API j group fetch permission set or default ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) WithID(id string) *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group fetch permission set or default ID params
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupFetchPermissionSetOrDefaultIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
