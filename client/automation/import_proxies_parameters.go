// Code generated by go-swagger; DO NOT EDIT.

package automation

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

// NewImportProxiesParams creates a new ImportProxiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportProxiesParams() *ImportProxiesParams {
	return &ImportProxiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportProxiesParamsWithTimeout creates a new ImportProxiesParams object
// with the ability to set a timeout on a request.
func NewImportProxiesParamsWithTimeout(timeout time.Duration) *ImportProxiesParams {
	return &ImportProxiesParams{
		timeout: timeout,
	}
}

// NewImportProxiesParamsWithContext creates a new ImportProxiesParams object
// with the ability to set a context for a request.
func NewImportProxiesParamsWithContext(ctx context.Context) *ImportProxiesParams {
	return &ImportProxiesParams{
		Context: ctx,
	}
}

// NewImportProxiesParamsWithHTTPClient creates a new ImportProxiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportProxiesParamsWithHTTPClient(client *http.Client) *ImportProxiesParams {
	return &ImportProxiesParams{
		HTTPClient: client,
	}
}

/* ImportProxiesParams contains all the parameters to send to the API endpoint
   for the import proxies operation.

   Typically these are written to a http.Request.
*/
type ImportProxiesParams struct {

	// Body.
	Body *models.ProxyImportSpecCollection

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

// WithDefaults hydrates default values in the import proxies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportProxiesParams) WithDefaults() *ImportProxiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import proxies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportProxiesParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := ImportProxiesParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the import proxies params
func (o *ImportProxiesParams) WithTimeout(timeout time.Duration) *ImportProxiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import proxies params
func (o *ImportProxiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import proxies params
func (o *ImportProxiesParams) WithContext(ctx context.Context) *ImportProxiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import proxies params
func (o *ImportProxiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import proxies params
func (o *ImportProxiesParams) WithHTTPClient(client *http.Client) *ImportProxiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import proxies params
func (o *ImportProxiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import proxies params
func (o *ImportProxiesParams) WithBody(body *models.ProxyImportSpecCollection) *ImportProxiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import proxies params
func (o *ImportProxiesParams) SetBody(body *models.ProxyImportSpecCollection) {
	o.Body = body
}

// WithXAPIVersion adds the xAPIVersion to the import proxies params
func (o *ImportProxiesParams) WithXAPIVersion(xAPIVersion string) *ImportProxiesParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the import proxies params
func (o *ImportProxiesParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *ImportProxiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
