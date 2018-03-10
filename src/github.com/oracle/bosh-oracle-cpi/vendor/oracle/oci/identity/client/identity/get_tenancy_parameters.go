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
)

// NewGetTenancyParams creates a new GetTenancyParams object
// with the default values initialized.
func NewGetTenancyParams() *GetTenancyParams {
	var ()
	return &GetTenancyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenancyParamsWithTimeout creates a new GetTenancyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenancyParamsWithTimeout(timeout time.Duration) *GetTenancyParams {
	var ()
	return &GetTenancyParams{

		timeout: timeout,
	}
}

// NewGetTenancyParamsWithContext creates a new GetTenancyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenancyParamsWithContext(ctx context.Context) *GetTenancyParams {
	var ()
	return &GetTenancyParams{

		Context: ctx,
	}
}

// NewGetTenancyParamsWithHTTPClient creates a new GetTenancyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenancyParamsWithHTTPClient(client *http.Client) *GetTenancyParams {
	var ()
	return &GetTenancyParams{
		HTTPClient: client,
	}
}

/*GetTenancyParams contains all the parameters to send to the API endpoint
for the get tenancy operation typically these are written to a http.Request
*/
type GetTenancyParams struct {

	/*TenancyID
	  The OCID of the tenancy.

	*/
	TenancyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenancy params
func (o *GetTenancyParams) WithTimeout(timeout time.Duration) *GetTenancyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenancy params
func (o *GetTenancyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenancy params
func (o *GetTenancyParams) WithContext(ctx context.Context) *GetTenancyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenancy params
func (o *GetTenancyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenancy params
func (o *GetTenancyParams) WithHTTPClient(client *http.Client) *GetTenancyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenancy params
func (o *GetTenancyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenancyID adds the tenancyID to the get tenancy params
func (o *GetTenancyParams) WithTenancyID(tenancyID string) *GetTenancyParams {
	o.SetTenancyID(tenancyID)
	return o
}

// SetTenancyID adds the tenancyId to the get tenancy params
func (o *GetTenancyParams) SetTenancyID(tenancyID string) {
	o.TenancyID = tenancyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenancyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenancyId
	if err := r.SetPathParam("tenancyId", o.TenancyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
