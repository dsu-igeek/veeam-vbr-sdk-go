// Code generated by go-swagger; DO NOT EDIT.

package configuration_backup

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
)

// NewGetConfigBackupOptionsParams creates a new GetConfigBackupOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigBackupOptionsParams() *GetConfigBackupOptionsParams {
	return &GetConfigBackupOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigBackupOptionsParamsWithTimeout creates a new GetConfigBackupOptionsParams object
// with the ability to set a timeout on a request.
func NewGetConfigBackupOptionsParamsWithTimeout(timeout time.Duration) *GetConfigBackupOptionsParams {
	return &GetConfigBackupOptionsParams{
		timeout: timeout,
	}
}

// NewGetConfigBackupOptionsParamsWithContext creates a new GetConfigBackupOptionsParams object
// with the ability to set a context for a request.
func NewGetConfigBackupOptionsParamsWithContext(ctx context.Context) *GetConfigBackupOptionsParams {
	return &GetConfigBackupOptionsParams{
		Context: ctx,
	}
}

// NewGetConfigBackupOptionsParamsWithHTTPClient creates a new GetConfigBackupOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigBackupOptionsParamsWithHTTPClient(client *http.Client) *GetConfigBackupOptionsParams {
	return &GetConfigBackupOptionsParams{
		HTTPClient: client,
	}
}

/* GetConfigBackupOptionsParams contains all the parameters to send to the API endpoint
   for the get config backup options operation.

   Typically these are written to a http.Request.
*/
type GetConfigBackupOptionsParams struct {

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

// WithDefaults hydrates default values in the get config backup options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigBackupOptionsParams) WithDefaults() *GetConfigBackupOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get config backup options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigBackupOptionsParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := GetConfigBackupOptionsParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get config backup options params
func (o *GetConfigBackupOptionsParams) WithTimeout(timeout time.Duration) *GetConfigBackupOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config backup options params
func (o *GetConfigBackupOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config backup options params
func (o *GetConfigBackupOptionsParams) WithContext(ctx context.Context) *GetConfigBackupOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config backup options params
func (o *GetConfigBackupOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get config backup options params
func (o *GetConfigBackupOptionsParams) WithHTTPClient(client *http.Client) *GetConfigBackupOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config backup options params
func (o *GetConfigBackupOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAPIVersion adds the xAPIVersion to the get config backup options params
func (o *GetConfigBackupOptionsParams) WithXAPIVersion(xAPIVersion string) *GetConfigBackupOptionsParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the get config backup options params
func (o *GetConfigBackupOptionsParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigBackupOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param x-api-version
	if err := r.SetHeaderParam("x-api-version", o.XAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
