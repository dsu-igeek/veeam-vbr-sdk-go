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

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// NewUpdateConfigBackupOptionsParams creates a new UpdateConfigBackupOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConfigBackupOptionsParams() *UpdateConfigBackupOptionsParams {
	return &UpdateConfigBackupOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigBackupOptionsParamsWithTimeout creates a new UpdateConfigBackupOptionsParams object
// with the ability to set a timeout on a request.
func NewUpdateConfigBackupOptionsParamsWithTimeout(timeout time.Duration) *UpdateConfigBackupOptionsParams {
	return &UpdateConfigBackupOptionsParams{
		timeout: timeout,
	}
}

// NewUpdateConfigBackupOptionsParamsWithContext creates a new UpdateConfigBackupOptionsParams object
// with the ability to set a context for a request.
func NewUpdateConfigBackupOptionsParamsWithContext(ctx context.Context) *UpdateConfigBackupOptionsParams {
	return &UpdateConfigBackupOptionsParams{
		Context: ctx,
	}
}

// NewUpdateConfigBackupOptionsParamsWithHTTPClient creates a new UpdateConfigBackupOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConfigBackupOptionsParamsWithHTTPClient(client *http.Client) *UpdateConfigBackupOptionsParams {
	return &UpdateConfigBackupOptionsParams{
		HTTPClient: client,
	}
}

/* UpdateConfigBackupOptionsParams contains all the parameters to send to the API endpoint
   for the update config backup options operation.

   Typically these are written to a http.Request.
*/
type UpdateConfigBackupOptionsParams struct {

	// Body.
	Body *models.ConfigBackupModel

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

// WithDefaults hydrates default values in the update config backup options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigBackupOptionsParams) WithDefaults() *UpdateConfigBackupOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update config backup options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigBackupOptionsParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := UpdateConfigBackupOptionsParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) WithTimeout(timeout time.Duration) *UpdateConfigBackupOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) WithContext(ctx context.Context) *UpdateConfigBackupOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) WithHTTPClient(client *http.Client) *UpdateConfigBackupOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) WithBody(body *models.ConfigBackupModel) *UpdateConfigBackupOptionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) SetBody(body *models.ConfigBackupModel) {
	o.Body = body
}

// WithXAPIVersion adds the xAPIVersion to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) WithXAPIVersion(xAPIVersion string) *UpdateConfigBackupOptionsParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the update config backup options params
func (o *UpdateConfigBackupOptionsParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigBackupOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
