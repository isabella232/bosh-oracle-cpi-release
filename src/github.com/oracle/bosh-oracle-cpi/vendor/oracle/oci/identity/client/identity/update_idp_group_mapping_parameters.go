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

// NewUpdateIdpGroupMappingParams creates a new UpdateIdpGroupMappingParams object
// with the default values initialized.
func NewUpdateIdpGroupMappingParams() *UpdateIdpGroupMappingParams {
	var ()
	return &UpdateIdpGroupMappingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIdpGroupMappingParamsWithTimeout creates a new UpdateIdpGroupMappingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateIdpGroupMappingParamsWithTimeout(timeout time.Duration) *UpdateIdpGroupMappingParams {
	var ()
	return &UpdateIdpGroupMappingParams{

		timeout: timeout,
	}
}

// NewUpdateIdpGroupMappingParamsWithContext creates a new UpdateIdpGroupMappingParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateIdpGroupMappingParamsWithContext(ctx context.Context) *UpdateIdpGroupMappingParams {
	var ()
	return &UpdateIdpGroupMappingParams{

		Context: ctx,
	}
}

// NewUpdateIdpGroupMappingParamsWithHTTPClient creates a new UpdateIdpGroupMappingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateIdpGroupMappingParamsWithHTTPClient(client *http.Client) *UpdateIdpGroupMappingParams {
	var ()
	return &UpdateIdpGroupMappingParams{
		HTTPClient: client,
	}
}

/*UpdateIdpGroupMappingParams contains all the parameters to send to the API endpoint
for the update idp group mapping operation typically these are written to a http.Request
*/
type UpdateIdpGroupMappingParams struct {

	/*IdentityProviderID
	  The OCID of the identity provider.

	*/
	IdentityProviderID string
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*MappingID
	  The OCID of the group mapping.

	*/
	MappingID string
	/*UpdateIdpGroupMappingDetails
	  Request object for updating an identity provider group mapping

	*/
	UpdateIdpGroupMappingDetails *models.UpdateIdpGroupMappingDetails

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithTimeout(timeout time.Duration) *UpdateIdpGroupMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithContext(ctx context.Context) *UpdateIdpGroupMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithHTTPClient(client *http.Client) *UpdateIdpGroupMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityProviderID adds the identityProviderID to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithIdentityProviderID(identityProviderID string) *UpdateIdpGroupMappingParams {
	o.SetIdentityProviderID(identityProviderID)
	return o
}

// SetIdentityProviderID adds the identityProviderId to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetIdentityProviderID(identityProviderID string) {
	o.IdentityProviderID = identityProviderID
}

// WithIfMatch adds the ifMatch to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithIfMatch(ifMatch *string) *UpdateIdpGroupMappingParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithMappingID adds the mappingID to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithMappingID(mappingID string) *UpdateIdpGroupMappingParams {
	o.SetMappingID(mappingID)
	return o
}

// SetMappingID adds the mappingId to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetMappingID(mappingID string) {
	o.MappingID = mappingID
}

// WithUpdateIdpGroupMappingDetails adds the updateIdpGroupMappingDetails to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) WithUpdateIdpGroupMappingDetails(updateIdpGroupMappingDetails *models.UpdateIdpGroupMappingDetails) *UpdateIdpGroupMappingParams {
	o.SetUpdateIdpGroupMappingDetails(updateIdpGroupMappingDetails)
	return o
}

// SetUpdateIdpGroupMappingDetails adds the updateIdpGroupMappingDetails to the update idp group mapping params
func (o *UpdateIdpGroupMappingParams) SetUpdateIdpGroupMappingDetails(updateIdpGroupMappingDetails *models.UpdateIdpGroupMappingDetails) {
	o.UpdateIdpGroupMappingDetails = updateIdpGroupMappingDetails
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIdpGroupMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identityProviderId
	if err := r.SetPathParam("identityProviderId", o.IdentityProviderID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param mappingId
	if err := r.SetPathParam("mappingId", o.MappingID); err != nil {
		return err
	}

	if o.UpdateIdpGroupMappingDetails == nil {
		o.UpdateIdpGroupMappingDetails = new(models.UpdateIdpGroupMappingDetails)
	}

	if err := r.SetBodyParam(o.UpdateIdpGroupMappingDetails); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
