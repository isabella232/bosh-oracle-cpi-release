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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListCrossConnectLocationsParams creates a new ListCrossConnectLocationsParams object
// with the default values initialized.
func NewListCrossConnectLocationsParams() *ListCrossConnectLocationsParams {
	var ()
	return &ListCrossConnectLocationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCrossConnectLocationsParamsWithTimeout creates a new ListCrossConnectLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCrossConnectLocationsParamsWithTimeout(timeout time.Duration) *ListCrossConnectLocationsParams {
	var ()
	return &ListCrossConnectLocationsParams{

		timeout: timeout,
	}
}

// NewListCrossConnectLocationsParamsWithContext creates a new ListCrossConnectLocationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCrossConnectLocationsParamsWithContext(ctx context.Context) *ListCrossConnectLocationsParams {
	var ()
	return &ListCrossConnectLocationsParams{

		Context: ctx,
	}
}

// NewListCrossConnectLocationsParamsWithHTTPClient creates a new ListCrossConnectLocationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCrossConnectLocationsParamsWithHTTPClient(client *http.Client) *ListCrossConnectLocationsParams {
	var ()
	return &ListCrossConnectLocationsParams{
		HTTPClient: client,
	}
}

/*ListCrossConnectLocationsParams contains all the parameters to send to the API endpoint
for the list cross connect locations operation typically these are written to a http.Request
*/
type ListCrossConnectLocationsParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithTimeout(timeout time.Duration) *ListCrossConnectLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithContext(ctx context.Context) *ListCrossConnectLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithHTTPClient(client *http.Client) *ListCrossConnectLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithCompartmentID(compartmentID string) *ListCrossConnectLocationsParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithLimit adds the limit to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithLimit(limit *int64) *ListCrossConnectLocationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) WithPage(page *string) *ListCrossConnectLocationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list cross connect locations params
func (o *ListCrossConnectLocationsParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListCrossConnectLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
