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

// NewUpdateVnicParams creates a new UpdateVnicParams object
// with the default values initialized.
func NewUpdateVnicParams() *UpdateVnicParams {
	var ()
	return &UpdateVnicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVnicParamsWithTimeout creates a new UpdateVnicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVnicParamsWithTimeout(timeout time.Duration) *UpdateVnicParams {
	var ()
	return &UpdateVnicParams{

		timeout: timeout,
	}
}

// NewUpdateVnicParamsWithContext creates a new UpdateVnicParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVnicParamsWithContext(ctx context.Context) *UpdateVnicParams {
	var ()
	return &UpdateVnicParams{

		Context: ctx,
	}
}

// NewUpdateVnicParamsWithHTTPClient creates a new UpdateVnicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVnicParamsWithHTTPClient(client *http.Client) *UpdateVnicParams {
	var ()
	return &UpdateVnicParams{
		HTTPClient: client,
	}
}

/*UpdateVnicParams contains all the parameters to send to the API endpoint
for the update vnic operation typically these are written to a http.Request
*/
type UpdateVnicParams struct {

	/*UpdateVnicDetails
	  Details object for updating a VNIC.

	*/
	UpdateVnicDetails *models.UpdateVnicDetails
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*VnicID
	  The OCID of the VNIC.

	*/
	VnicID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update vnic params
func (o *UpdateVnicParams) WithTimeout(timeout time.Duration) *UpdateVnicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vnic params
func (o *UpdateVnicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vnic params
func (o *UpdateVnicParams) WithContext(ctx context.Context) *UpdateVnicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vnic params
func (o *UpdateVnicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vnic params
func (o *UpdateVnicParams) WithHTTPClient(client *http.Client) *UpdateVnicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vnic params
func (o *UpdateVnicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateVnicDetails adds the updateVnicDetails to the update vnic params
func (o *UpdateVnicParams) WithUpdateVnicDetails(updateVnicDetails *models.UpdateVnicDetails) *UpdateVnicParams {
	o.SetUpdateVnicDetails(updateVnicDetails)
	return o
}

// SetUpdateVnicDetails adds the updateVnicDetails to the update vnic params
func (o *UpdateVnicParams) SetUpdateVnicDetails(updateVnicDetails *models.UpdateVnicDetails) {
	o.UpdateVnicDetails = updateVnicDetails
}

// WithIfMatch adds the ifMatch to the update vnic params
func (o *UpdateVnicParams) WithIfMatch(ifMatch *string) *UpdateVnicParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update vnic params
func (o *UpdateVnicParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithVnicID adds the vnicID to the update vnic params
func (o *UpdateVnicParams) WithVnicID(vnicID string) *UpdateVnicParams {
	o.SetVnicID(vnicID)
	return o
}

// SetVnicID adds the vnicId to the update vnic params
func (o *UpdateVnicParams) SetVnicID(vnicID string) {
	o.VnicID = vnicID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVnicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateVnicDetails != nil {
		if err := r.SetBodyParam(o.UpdateVnicDetails); err != nil {
			return err
		}
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param vnicId
	if err := r.SetPathParam("vnicId", o.VnicID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
