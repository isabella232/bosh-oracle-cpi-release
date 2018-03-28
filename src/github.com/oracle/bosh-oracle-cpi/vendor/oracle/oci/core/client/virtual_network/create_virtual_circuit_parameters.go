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

// NewCreateVirtualCircuitParams creates a new CreateVirtualCircuitParams object
// with the default values initialized.
func NewCreateVirtualCircuitParams() *CreateVirtualCircuitParams {
	var ()
	return &CreateVirtualCircuitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualCircuitParamsWithTimeout creates a new CreateVirtualCircuitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVirtualCircuitParamsWithTimeout(timeout time.Duration) *CreateVirtualCircuitParams {
	var ()
	return &CreateVirtualCircuitParams{

		timeout: timeout,
	}
}

// NewCreateVirtualCircuitParamsWithContext creates a new CreateVirtualCircuitParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVirtualCircuitParamsWithContext(ctx context.Context) *CreateVirtualCircuitParams {
	var ()
	return &CreateVirtualCircuitParams{

		Context: ctx,
	}
}

// NewCreateVirtualCircuitParamsWithHTTPClient creates a new CreateVirtualCircuitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVirtualCircuitParamsWithHTTPClient(client *http.Client) *CreateVirtualCircuitParams {
	var ()
	return &CreateVirtualCircuitParams{
		HTTPClient: client,
	}
}

/*CreateVirtualCircuitParams contains all the parameters to send to the API endpoint
for the create virtual circuit operation typically these are written to a http.Request
*/
type CreateVirtualCircuitParams struct {

	/*CreateVirtualCircuitDetails
	  Details to create a VirtualCircuit.

	*/
	CreateVirtualCircuitDetails *models.CreateVirtualCircuitDetails
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

// WithTimeout adds the timeout to the create virtual circuit params
func (o *CreateVirtualCircuitParams) WithTimeout(timeout time.Duration) *CreateVirtualCircuitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual circuit params
func (o *CreateVirtualCircuitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual circuit params
func (o *CreateVirtualCircuitParams) WithContext(ctx context.Context) *CreateVirtualCircuitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual circuit params
func (o *CreateVirtualCircuitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual circuit params
func (o *CreateVirtualCircuitParams) WithHTTPClient(client *http.Client) *CreateVirtualCircuitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual circuit params
func (o *CreateVirtualCircuitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateVirtualCircuitDetails adds the createVirtualCircuitDetails to the create virtual circuit params
func (o *CreateVirtualCircuitParams) WithCreateVirtualCircuitDetails(createVirtualCircuitDetails *models.CreateVirtualCircuitDetails) *CreateVirtualCircuitParams {
	o.SetCreateVirtualCircuitDetails(createVirtualCircuitDetails)
	return o
}

// SetCreateVirtualCircuitDetails adds the createVirtualCircuitDetails to the create virtual circuit params
func (o *CreateVirtualCircuitParams) SetCreateVirtualCircuitDetails(createVirtualCircuitDetails *models.CreateVirtualCircuitDetails) {
	o.CreateVirtualCircuitDetails = createVirtualCircuitDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create virtual circuit params
func (o *CreateVirtualCircuitParams) WithOpcRetryToken(opcRetryToken *string) *CreateVirtualCircuitParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create virtual circuit params
func (o *CreateVirtualCircuitParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualCircuitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateVirtualCircuitDetails != nil {
		if err := r.SetBodyParam(o.CreateVirtualCircuitDetails); err != nil {
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
