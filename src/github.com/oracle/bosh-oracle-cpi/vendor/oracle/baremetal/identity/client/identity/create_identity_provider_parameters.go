package identity

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

	"oracle/baremetal/identity/models"
)

// NewCreateIdentityProviderParams creates a new CreateIdentityProviderParams object
// with the default values initialized.
func NewCreateIdentityProviderParams() *CreateIdentityProviderParams {
	var ()
	return &CreateIdentityProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIdentityProviderParamsWithTimeout creates a new CreateIdentityProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateIdentityProviderParamsWithTimeout(timeout time.Duration) *CreateIdentityProviderParams {
	var ()
	return &CreateIdentityProviderParams{

		timeout: timeout,
	}
}

// NewCreateIdentityProviderParamsWithContext creates a new CreateIdentityProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateIdentityProviderParamsWithContext(ctx context.Context) *CreateIdentityProviderParams {
	var ()
	return &CreateIdentityProviderParams{

		Context: ctx,
	}
}

// NewCreateIdentityProviderParamsWithHTTPClient creates a new CreateIdentityProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateIdentityProviderParamsWithHTTPClient(client *http.Client) *CreateIdentityProviderParams {
	var ()
	return &CreateIdentityProviderParams{
		HTTPClient: client,
	}
}

/*CreateIdentityProviderParams contains all the parameters to send to the API endpoint
for the create identity provider operation typically these are written to a http.Request
*/
type CreateIdentityProviderParams struct {

	/*CreateIdentityProviderDetails
	  Request object for creating a new SAML2 identity provider.

	*/
	CreateIdentityProviderDetails models.CreateIdentityProviderDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create identity provider params
func (o *CreateIdentityProviderParams) WithTimeout(timeout time.Duration) *CreateIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create identity provider params
func (o *CreateIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create identity provider params
func (o *CreateIdentityProviderParams) WithContext(ctx context.Context) *CreateIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create identity provider params
func (o *CreateIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create identity provider params
func (o *CreateIdentityProviderParams) WithHTTPClient(client *http.Client) *CreateIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create identity provider params
func (o *CreateIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateIdentityProviderDetails adds the createIdentityProviderDetails to the create identity provider params
func (o *CreateIdentityProviderParams) WithCreateIdentityProviderDetails(createIdentityProviderDetails models.CreateIdentityProviderDetails) *CreateIdentityProviderParams {
	o.SetCreateIdentityProviderDetails(createIdentityProviderDetails)
	return o
}

// SetCreateIdentityProviderDetails adds the createIdentityProviderDetails to the create identity provider params
func (o *CreateIdentityProviderParams) SetCreateIdentityProviderDetails(createIdentityProviderDetails models.CreateIdentityProviderDetails) {
	o.CreateIdentityProviderDetails = createIdentityProviderDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create identity provider params
func (o *CreateIdentityProviderParams) WithOpcRetryToken(opcRetryToken *string) *CreateIdentityProviderParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create identity provider params
func (o *CreateIdentityProviderParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CreateIdentityProviderDetails); err != nil {
		return err
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
