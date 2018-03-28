// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetInstanceParams creates a new GetInstanceParams object
// with the default values initialized.
func NewGetInstanceParams() *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceParamsWithTimeout creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceParamsWithTimeout(timeout time.Duration) *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		timeout: timeout,
	}
}

// NewGetInstanceParamsWithContext creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceParamsWithContext(ctx context.Context) *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		Context: ctx,
	}
}

// NewGetInstanceParamsWithHTTPClient creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceParamsWithHTTPClient(client *http.Client) *GetInstanceParams {
	var ()
	return &GetInstanceParams{
		HTTPClient: client,
	}
}

/*GetInstanceParams contains all the parameters to send to the API endpoint
for the get instance operation typically these are written to a http.Request
*/
type GetInstanceParams struct {

	/*InstanceID
	  The OCID of the instance.

	*/
	InstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance params
func (o *GetInstanceParams) WithTimeout(timeout time.Duration) *GetInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance params
func (o *GetInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance params
func (o *GetInstanceParams) WithContext(ctx context.Context) *GetInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance params
func (o *GetInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance params
func (o *GetInstanceParams) WithHTTPClient(client *http.Client) *GetInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance params
func (o *GetInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceID adds the instanceID to the get instance params
func (o *GetInstanceParams) WithInstanceID(instanceID string) *GetInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get instance params
func (o *GetInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
