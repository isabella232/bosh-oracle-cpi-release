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

	"oracle/baremetal/core/models"
)

// NewCreateCrossConnectGroupParams creates a new CreateCrossConnectGroupParams object
// with the default values initialized.
func NewCreateCrossConnectGroupParams() *CreateCrossConnectGroupParams {
	var ()
	return &CreateCrossConnectGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCrossConnectGroupParamsWithTimeout creates a new CreateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCrossConnectGroupParamsWithTimeout(timeout time.Duration) *CreateCrossConnectGroupParams {
	var ()
	return &CreateCrossConnectGroupParams{

		timeout: timeout,
	}
}

// NewCreateCrossConnectGroupParamsWithContext creates a new CreateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCrossConnectGroupParamsWithContext(ctx context.Context) *CreateCrossConnectGroupParams {
	var ()
	return &CreateCrossConnectGroupParams{

		Context: ctx,
	}
}

// NewCreateCrossConnectGroupParamsWithHTTPClient creates a new CreateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCrossConnectGroupParamsWithHTTPClient(client *http.Client) *CreateCrossConnectGroupParams {
	var ()
	return &CreateCrossConnectGroupParams{
		HTTPClient: client,
	}
}

/*CreateCrossConnectGroupParams contains all the parameters to send to the API endpoint
for the create cross connect group operation typically these are written to a http.Request
*/
type CreateCrossConnectGroupParams struct {

	/*CreateCrossConnectGroupDetails
	  Details to create a CrossConnectGroup

	*/
	CreateCrossConnectGroupDetails *models.CreateCrossConnectGroupDetails
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

// WithTimeout adds the timeout to the create cross connect group params
func (o *CreateCrossConnectGroupParams) WithTimeout(timeout time.Duration) *CreateCrossConnectGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cross connect group params
func (o *CreateCrossConnectGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cross connect group params
func (o *CreateCrossConnectGroupParams) WithContext(ctx context.Context) *CreateCrossConnectGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cross connect group params
func (o *CreateCrossConnectGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cross connect group params
func (o *CreateCrossConnectGroupParams) WithHTTPClient(client *http.Client) *CreateCrossConnectGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cross connect group params
func (o *CreateCrossConnectGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateCrossConnectGroupDetails adds the createCrossConnectGroupDetails to the create cross connect group params
func (o *CreateCrossConnectGroupParams) WithCreateCrossConnectGroupDetails(createCrossConnectGroupDetails *models.CreateCrossConnectGroupDetails) *CreateCrossConnectGroupParams {
	o.SetCreateCrossConnectGroupDetails(createCrossConnectGroupDetails)
	return o
}

// SetCreateCrossConnectGroupDetails adds the createCrossConnectGroupDetails to the create cross connect group params
func (o *CreateCrossConnectGroupParams) SetCreateCrossConnectGroupDetails(createCrossConnectGroupDetails *models.CreateCrossConnectGroupDetails) {
	o.CreateCrossConnectGroupDetails = createCrossConnectGroupDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create cross connect group params
func (o *CreateCrossConnectGroupParams) WithOpcRetryToken(opcRetryToken *string) *CreateCrossConnectGroupParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create cross connect group params
func (o *CreateCrossConnectGroupParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCrossConnectGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateCrossConnectGroupDetails == nil {
		o.CreateCrossConnectGroupDetails = new(models.CreateCrossConnectGroupDetails)
	}

	if err := r.SetBodyParam(o.CreateCrossConnectGroupDetails); err != nil {
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
