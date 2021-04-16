// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// NewCreateProxyParams creates a new CreateProxyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProxyParams() *CreateProxyParams {
	return &CreateProxyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProxyParamsWithTimeout creates a new CreateProxyParams object
// with the ability to set a timeout on a request.
func NewCreateProxyParamsWithTimeout(timeout time.Duration) *CreateProxyParams {
	return &CreateProxyParams{
		timeout: timeout,
	}
}

// NewCreateProxyParamsWithContext creates a new CreateProxyParams object
// with the ability to set a context for a request.
func NewCreateProxyParamsWithContext(ctx context.Context) *CreateProxyParams {
	return &CreateProxyParams{
		Context: ctx,
	}
}

// NewCreateProxyParamsWithHTTPClient creates a new CreateProxyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProxyParamsWithHTTPClient(client *http.Client) *CreateProxyParams {
	return &CreateProxyParams{
		HTTPClient: client,
	}
}

/* CreateProxyParams contains all the parameters to send to the API endpoint
   for the create proxy operation.

   Typically these are written to a http.Request.
*/
type CreateProxyParams struct {

	// Body.
	Body *models.ProxySpec

	/* XAPIVersion.

	     Version and revision of the client REST API. Must be in the following
	format: *\<version\>-\<revision\>*.


	     Default: "1.0-rev1"
	*/
	XAPIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProxyParams) WithDefaults() *CreateProxyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create proxy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProxyParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := CreateProxyParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create proxy params
func (o *CreateProxyParams) WithTimeout(timeout time.Duration) *CreateProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create proxy params
func (o *CreateProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create proxy params
func (o *CreateProxyParams) WithContext(ctx context.Context) *CreateProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create proxy params
func (o *CreateProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create proxy params
func (o *CreateProxyParams) WithHTTPClient(client *http.Client) *CreateProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create proxy params
func (o *CreateProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create proxy params
func (o *CreateProxyParams) WithBody(body *models.ProxySpec) *CreateProxyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create proxy params
func (o *CreateProxyParams) SetBody(body *models.ProxySpec) {
	o.Body = body
}

// WithXAPIVersion adds the xAPIVersion to the create proxy params
func (o *CreateProxyParams) WithXAPIVersion(xAPIVersion string) *CreateProxyParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the create proxy params
func (o *CreateProxyParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param x-api-version
	if err := r.SetHeaderParam("x-api-version", o.XAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
