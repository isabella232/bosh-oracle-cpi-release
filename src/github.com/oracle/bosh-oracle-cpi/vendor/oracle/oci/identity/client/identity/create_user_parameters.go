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

	"oracle/oci/identity/models"
)

// NewCreateUserParams creates a new CreateUserParams object
// with the default values initialized.
func NewCreateUserParams() *CreateUserParams {
	var ()
	return &CreateUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserParamsWithTimeout creates a new CreateUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserParamsWithTimeout(timeout time.Duration) *CreateUserParams {
	var ()
	return &CreateUserParams{

		timeout: timeout,
	}
}

// NewCreateUserParamsWithContext creates a new CreateUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserParamsWithContext(ctx context.Context) *CreateUserParams {
	var ()
	return &CreateUserParams{

		Context: ctx,
	}
}

// NewCreateUserParamsWithHTTPClient creates a new CreateUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserParamsWithHTTPClient(client *http.Client) *CreateUserParams {
	var ()
	return &CreateUserParams{
		HTTPClient: client,
	}
}

/*CreateUserParams contains all the parameters to send to the API endpoint
for the create user operation typically these are written to a http.Request
*/
type CreateUserParams struct {

	/*CreateUserDetails
	  Request object for creating a new user.

	*/
	CreateUserDetails *models.CreateUserDetails
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

// WithTimeout adds the timeout to the create user params
func (o *CreateUserParams) WithTimeout(timeout time.Duration) *CreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user params
func (o *CreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user params
func (o *CreateUserParams) WithContext(ctx context.Context) *CreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user params
func (o *CreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) WithHTTPClient(client *http.Client) *CreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateUserDetails adds the createUserDetails to the create user params
func (o *CreateUserParams) WithCreateUserDetails(createUserDetails *models.CreateUserDetails) *CreateUserParams {
	o.SetCreateUserDetails(createUserDetails)
	return o
}

// SetCreateUserDetails adds the createUserDetails to the create user params
func (o *CreateUserParams) SetCreateUserDetails(createUserDetails *models.CreateUserDetails) {
	o.CreateUserDetails = createUserDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create user params
func (o *CreateUserParams) WithOpcRetryToken(opcRetryToken *string) *CreateUserParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create user params
func (o *CreateUserParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateUserDetails == nil {
		o.CreateUserDetails = new(models.CreateUserDetails)
	}

	if err := r.SetBodyParam(o.CreateUserDetails); err != nil {
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
