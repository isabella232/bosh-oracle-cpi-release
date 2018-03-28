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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTerminateInstanceParams creates a new TerminateInstanceParams object
// with the default values initialized.
func NewTerminateInstanceParams() *TerminateInstanceParams {
	var ()
	return &TerminateInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTerminateInstanceParamsWithTimeout creates a new TerminateInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTerminateInstanceParamsWithTimeout(timeout time.Duration) *TerminateInstanceParams {
	var ()
	return &TerminateInstanceParams{

		timeout: timeout,
	}
}

// NewTerminateInstanceParamsWithContext creates a new TerminateInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewTerminateInstanceParamsWithContext(ctx context.Context) *TerminateInstanceParams {
	var ()
	return &TerminateInstanceParams{

		Context: ctx,
	}
}

// NewTerminateInstanceParamsWithHTTPClient creates a new TerminateInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTerminateInstanceParamsWithHTTPClient(client *http.Client) *TerminateInstanceParams {
	var ()
	return &TerminateInstanceParams{
		HTTPClient: client,
	}
}

/*TerminateInstanceParams contains all the parameters to send to the API endpoint
for the terminate instance operation typically these are written to a http.Request
*/
type TerminateInstanceParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*InstanceID
	  The OCID of the instance.

	*/
	InstanceID string
	/*PreserveBootVolume
	  Specifies whether to delete or preserve the boot volume when terminating an instance.
	The default value is false.


	*/
	PreserveBootVolume *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the terminate instance params
func (o *TerminateInstanceParams) WithTimeout(timeout time.Duration) *TerminateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terminate instance params
func (o *TerminateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terminate instance params
func (o *TerminateInstanceParams) WithContext(ctx context.Context) *TerminateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terminate instance params
func (o *TerminateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terminate instance params
func (o *TerminateInstanceParams) WithHTTPClient(client *http.Client) *TerminateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terminate instance params
func (o *TerminateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the terminate instance params
func (o *TerminateInstanceParams) WithIfMatch(ifMatch *string) *TerminateInstanceParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the terminate instance params
func (o *TerminateInstanceParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithInstanceID adds the instanceID to the terminate instance params
func (o *TerminateInstanceParams) WithInstanceID(instanceID string) *TerminateInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the terminate instance params
func (o *TerminateInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithPreserveBootVolume adds the preserveBootVolume to the terminate instance params
func (o *TerminateInstanceParams) WithPreserveBootVolume(preserveBootVolume *bool) *TerminateInstanceParams {
	o.SetPreserveBootVolume(preserveBootVolume)
	return o
}

// SetPreserveBootVolume adds the preserveBootVolume to the terminate instance params
func (o *TerminateInstanceParams) SetPreserveBootVolume(preserveBootVolume *bool) {
	o.PreserveBootVolume = preserveBootVolume
}

// WriteToRequest writes these params to a swagger request
func (o *TerminateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
		return err
	}

	if o.PreserveBootVolume != nil {

		// query param preserveBootVolume
		var qrPreserveBootVolume bool
		if o.PreserveBootVolume != nil {
			qrPreserveBootVolume = *o.PreserveBootVolume
		}
		qPreserveBootVolume := swag.FormatBool(qrPreserveBootVolume)
		if qPreserveBootVolume != "" {
			if err := r.SetQueryParam("preserveBootVolume", qPreserveBootVolume); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
