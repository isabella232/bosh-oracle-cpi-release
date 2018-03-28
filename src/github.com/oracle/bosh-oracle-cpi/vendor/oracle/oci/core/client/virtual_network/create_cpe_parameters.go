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

// NewCreateCpeParams creates a new CreateCpeParams object
// with the default values initialized.
func NewCreateCpeParams() *CreateCpeParams {
	var ()
	return &CreateCpeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCpeParamsWithTimeout creates a new CreateCpeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCpeParamsWithTimeout(timeout time.Duration) *CreateCpeParams {
	var ()
	return &CreateCpeParams{

		timeout: timeout,
	}
}

// NewCreateCpeParamsWithContext creates a new CreateCpeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCpeParamsWithContext(ctx context.Context) *CreateCpeParams {
	var ()
	return &CreateCpeParams{

		Context: ctx,
	}
}

// NewCreateCpeParamsWithHTTPClient creates a new CreateCpeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCpeParamsWithHTTPClient(client *http.Client) *CreateCpeParams {
	var ()
	return &CreateCpeParams{
		HTTPClient: client,
	}
}

/*CreateCpeParams contains all the parameters to send to the API endpoint
for the create cpe operation typically these are written to a http.Request
*/
type CreateCpeParams struct {

	/*CreateCpeDetails
	  Details for creating a CPE.

	*/
	CreateCpeDetails *models.CreateCpeDetails
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

// WithTimeout adds the timeout to the create cpe params
func (o *CreateCpeParams) WithTimeout(timeout time.Duration) *CreateCpeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cpe params
func (o *CreateCpeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cpe params
func (o *CreateCpeParams) WithContext(ctx context.Context) *CreateCpeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cpe params
func (o *CreateCpeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cpe params
func (o *CreateCpeParams) WithHTTPClient(client *http.Client) *CreateCpeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cpe params
func (o *CreateCpeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateCpeDetails adds the createCpeDetails to the create cpe params
func (o *CreateCpeParams) WithCreateCpeDetails(createCpeDetails *models.CreateCpeDetails) *CreateCpeParams {
	o.SetCreateCpeDetails(createCpeDetails)
	return o
}

// SetCreateCpeDetails adds the createCpeDetails to the create cpe params
func (o *CreateCpeParams) SetCreateCpeDetails(createCpeDetails *models.CreateCpeDetails) {
	o.CreateCpeDetails = createCpeDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create cpe params
func (o *CreateCpeParams) WithOpcRetryToken(opcRetryToken *string) *CreateCpeParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create cpe params
func (o *CreateCpeParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCpeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateCpeDetails != nil {
		if err := r.SetBodyParam(o.CreateCpeDetails); err != nil {
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
