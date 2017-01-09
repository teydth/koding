package j_account

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

// NewPostRemoteAPIJAccountFetchMetaInformationIDParams creates a new PostRemoteAPIJAccountFetchMetaInformationIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountFetchMetaInformationIDParams() *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchMetaInformationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountFetchMetaInformationIDParamsWithTimeout creates a new PostRemoteAPIJAccountFetchMetaInformationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountFetchMetaInformationIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchMetaInformationIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountFetchMetaInformationIDParamsWithContext creates a new PostRemoteAPIJAccountFetchMetaInformationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountFetchMetaInformationIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchMetaInformationIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountFetchMetaInformationIDParams contains all the parameters to send to the API endpoint
for the post remote API j account fetch meta information ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountFetchMetaInformationIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) WithID(id string) *PostRemoteAPIJAccountFetchMetaInformationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account fetch meta information ID params
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountFetchMetaInformationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
