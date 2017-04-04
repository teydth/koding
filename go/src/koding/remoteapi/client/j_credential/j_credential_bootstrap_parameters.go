package j_credential

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

// NewJCredentialBootstrapParams creates a new JCredentialBootstrapParams object
// with the default values initialized.
func NewJCredentialBootstrapParams() *JCredentialBootstrapParams {
	var ()
	return &JCredentialBootstrapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJCredentialBootstrapParamsWithTimeout creates a new JCredentialBootstrapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJCredentialBootstrapParamsWithTimeout(timeout time.Duration) *JCredentialBootstrapParams {
	var ()
	return &JCredentialBootstrapParams{

		timeout: timeout,
	}
}

// NewJCredentialBootstrapParamsWithContext creates a new JCredentialBootstrapParams object
// with the default values initialized, and the ability to set a context for a request
func NewJCredentialBootstrapParamsWithContext(ctx context.Context) *JCredentialBootstrapParams {
	var ()
	return &JCredentialBootstrapParams{

		Context: ctx,
	}
}

/*JCredentialBootstrapParams contains all the parameters to send to the API endpoint
for the j credential bootstrap operation typically these are written to a http.Request
*/
type JCredentialBootstrapParams struct {

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

// WithTimeout adds the timeout to the j credential bootstrap params
func (o *JCredentialBootstrapParams) WithTimeout(timeout time.Duration) *JCredentialBootstrapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j credential bootstrap params
func (o *JCredentialBootstrapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j credential bootstrap params
func (o *JCredentialBootstrapParams) WithContext(ctx context.Context) *JCredentialBootstrapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j credential bootstrap params
func (o *JCredentialBootstrapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j credential bootstrap params
func (o *JCredentialBootstrapParams) WithBody(body models.DefaultSelector) *JCredentialBootstrapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j credential bootstrap params
func (o *JCredentialBootstrapParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j credential bootstrap params
func (o *JCredentialBootstrapParams) WithID(id string) *JCredentialBootstrapParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j credential bootstrap params
func (o *JCredentialBootstrapParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JCredentialBootstrapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
