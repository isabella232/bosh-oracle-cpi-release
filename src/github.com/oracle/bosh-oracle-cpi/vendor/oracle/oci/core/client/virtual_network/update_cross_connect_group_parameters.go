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

	"oracle/oci/core/models"
)

// NewUpdateCrossConnectGroupParams creates a new UpdateCrossConnectGroupParams object
// with the default values initialized.
func NewUpdateCrossConnectGroupParams() *UpdateCrossConnectGroupParams {
	var ()
	return &UpdateCrossConnectGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCrossConnectGroupParamsWithTimeout creates a new UpdateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCrossConnectGroupParamsWithTimeout(timeout time.Duration) *UpdateCrossConnectGroupParams {
	var ()
	return &UpdateCrossConnectGroupParams{

		timeout: timeout,
	}
}

// NewUpdateCrossConnectGroupParamsWithContext creates a new UpdateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCrossConnectGroupParamsWithContext(ctx context.Context) *UpdateCrossConnectGroupParams {
	var ()
	return &UpdateCrossConnectGroupParams{

		Context: ctx,
	}
}

// NewUpdateCrossConnectGroupParamsWithHTTPClient creates a new UpdateCrossConnectGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCrossConnectGroupParamsWithHTTPClient(client *http.Client) *UpdateCrossConnectGroupParams {
	var ()
	return &UpdateCrossConnectGroupParams{
		HTTPClient: client,
	}
}

/*UpdateCrossConnectGroupParams contains all the parameters to send to the API endpoint
for the update cross connect group operation typically these are written to a http.Request
*/
type UpdateCrossConnectGroupParams struct {

	/*UpdateCrossConnectGroupDetails
	  Update CrossConnectGroup fields

	*/
	UpdateCrossConnectGroupDetails *models.UpdateCrossConnectGroupDetails
	/*CrossConnectGroupID
	  The OCID of the cross-connect group.

	*/
	CrossConnectGroupID string
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithTimeout(timeout time.Duration) *UpdateCrossConnectGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithContext(ctx context.Context) *UpdateCrossConnectGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithHTTPClient(client *http.Client) *UpdateCrossConnectGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateCrossConnectGroupDetails adds the updateCrossConnectGroupDetails to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithUpdateCrossConnectGroupDetails(updateCrossConnectGroupDetails *models.UpdateCrossConnectGroupDetails) *UpdateCrossConnectGroupParams {
	o.SetUpdateCrossConnectGroupDetails(updateCrossConnectGroupDetails)
	return o
}

// SetUpdateCrossConnectGroupDetails adds the updateCrossConnectGroupDetails to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetUpdateCrossConnectGroupDetails(updateCrossConnectGroupDetails *models.UpdateCrossConnectGroupDetails) {
	o.UpdateCrossConnectGroupDetails = updateCrossConnectGroupDetails
}

// WithCrossConnectGroupID adds the crossConnectGroupID to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithCrossConnectGroupID(crossConnectGroupID string) *UpdateCrossConnectGroupParams {
	o.SetCrossConnectGroupID(crossConnectGroupID)
	return o
}

// SetCrossConnectGroupID adds the crossConnectGroupId to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetCrossConnectGroupID(crossConnectGroupID string) {
	o.CrossConnectGroupID = crossConnectGroupID
}

// WithIfMatch adds the ifMatch to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) WithIfMatch(ifMatch *string) *UpdateCrossConnectGroupParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update cross connect group params
func (o *UpdateCrossConnectGroupParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCrossConnectGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateCrossConnectGroupDetails == nil {
		o.UpdateCrossConnectGroupDetails = new(models.UpdateCrossConnectGroupDetails)
	}

	if err := r.SetBodyParam(o.UpdateCrossConnectGroupDetails); err != nil {
		return err
	}

	// path param crossConnectGroupId
	if err := r.SetPathParam("crossConnectGroupId", o.CrossConnectGroupID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
