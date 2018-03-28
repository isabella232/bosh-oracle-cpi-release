// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

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

	models "oracle/oci/core/models"
)

// NewCreatePrivateIPParams creates a new CreatePrivateIPParams object
// with the default values initialized.
func NewCreatePrivateIPParams() *CreatePrivateIPParams {
	var ()
	return &CreatePrivateIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivateIPParamsWithTimeout creates a new CreatePrivateIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePrivateIPParamsWithTimeout(timeout time.Duration) *CreatePrivateIPParams {
	var ()
	return &CreatePrivateIPParams{

		timeout: timeout,
	}
}

// NewCreatePrivateIPParamsWithContext creates a new CreatePrivateIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePrivateIPParamsWithContext(ctx context.Context) *CreatePrivateIPParams {
	var ()
	return &CreatePrivateIPParams{

		Context: ctx,
	}
}

// NewCreatePrivateIPParamsWithHTTPClient creates a new CreatePrivateIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePrivateIPParamsWithHTTPClient(client *http.Client) *CreatePrivateIPParams {
	var ()
	return &CreatePrivateIPParams{
		HTTPClient: client,
	}
}

/*CreatePrivateIPParams contains all the parameters to send to the API endpoint
for the create private Ip operation typically these are written to a http.Request
*/
type CreatePrivateIPParams struct {

	/*CreatePrivateIPDetails
	  Create private IP details.

	*/
	CreatePrivateIPDetails *models.CreatePrivateIPDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create private Ip params
func (o *CreatePrivateIPParams) WithTimeout(timeout time.Duration) *CreatePrivateIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create private Ip params
func (o *CreatePrivateIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create private Ip params
func (o *CreatePrivateIPParams) WithContext(ctx context.Context) *CreatePrivateIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create private Ip params
func (o *CreatePrivateIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create private Ip params
func (o *CreatePrivateIPParams) WithHTTPClient(client *http.Client) *CreatePrivateIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create private Ip params
func (o *CreatePrivateIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatePrivateIPDetails adds the createPrivateIPDetails to the create private Ip params
func (o *CreatePrivateIPParams) WithCreatePrivateIPDetails(createPrivateIPDetails *models.CreatePrivateIPDetails) *CreatePrivateIPParams {
	o.SetCreatePrivateIPDetails(createPrivateIPDetails)
	return o
}

// SetCreatePrivateIPDetails adds the createPrivateIpDetails to the create private Ip params
func (o *CreatePrivateIPParams) SetCreatePrivateIPDetails(createPrivateIPDetails *models.CreatePrivateIPDetails) {
	o.CreatePrivateIPDetails = createPrivateIPDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create private Ip params
func (o *CreatePrivateIPParams) WithOpcRetryToken(opcRetryToken *string) *CreatePrivateIPParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create private Ip params
func (o *CreatePrivateIPParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivateIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatePrivateIPDetails != nil {
		if err := r.SetBodyParam(o.CreatePrivateIPDetails); err != nil {
			return err
		}
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
