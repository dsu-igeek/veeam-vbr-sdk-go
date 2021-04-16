// Code generated by go-swagger; DO NOT EDIT.

package repositories

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

// NewEnableScaleOutExtentMaintenanceModeParams creates a new EnableScaleOutExtentMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableScaleOutExtentMaintenanceModeParams() *EnableScaleOutExtentMaintenanceModeParams {
	return &EnableScaleOutExtentMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableScaleOutExtentMaintenanceModeParamsWithTimeout creates a new EnableScaleOutExtentMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewEnableScaleOutExtentMaintenanceModeParamsWithTimeout(timeout time.Duration) *EnableScaleOutExtentMaintenanceModeParams {
	return &EnableScaleOutExtentMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewEnableScaleOutExtentMaintenanceModeParamsWithContext creates a new EnableScaleOutExtentMaintenanceModeParams object
// with the ability to set a context for a request.
func NewEnableScaleOutExtentMaintenanceModeParamsWithContext(ctx context.Context) *EnableScaleOutExtentMaintenanceModeParams {
	return &EnableScaleOutExtentMaintenanceModeParams{
		Context: ctx,
	}
}

// NewEnableScaleOutExtentMaintenanceModeParamsWithHTTPClient creates a new EnableScaleOutExtentMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableScaleOutExtentMaintenanceModeParamsWithHTTPClient(client *http.Client) *EnableScaleOutExtentMaintenanceModeParams {
	return &EnableScaleOutExtentMaintenanceModeParams{
		HTTPClient: client,
	}
}

/* EnableScaleOutExtentMaintenanceModeParams contains all the parameters to send to the API endpoint
   for the enable scale out extent maintenance mode operation.

   Typically these are written to a http.Request.
*/
type EnableScaleOutExtentMaintenanceModeParams struct {

	// Body.
	Body *models.ScaleOutExtentMaintenanceSpec

	/* ID.

	   ID of the scale-out backup repository extent.

	   Format: uuid
	*/
	ID strfmt.UUID

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

// WithDefaults hydrates default values in the enable scale out extent maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableScaleOutExtentMaintenanceModeParams) WithDefaults() *EnableScaleOutExtentMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable scale out extent maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableScaleOutExtentMaintenanceModeParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := EnableScaleOutExtentMaintenanceModeParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithTimeout(timeout time.Duration) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithContext(ctx context.Context) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithHTTPClient(client *http.Client) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithBody(body *models.ScaleOutExtentMaintenanceSpec) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetBody(body *models.ScaleOutExtentMaintenanceSpec) {
	o.Body = body
}

// WithID adds the id to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithID(id strfmt.UUID) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithXAPIVersion adds the xAPIVersion to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) WithXAPIVersion(xAPIVersion string) *EnableScaleOutExtentMaintenanceModeParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the enable scale out extent maintenance mode params
func (o *EnableScaleOutExtentMaintenanceModeParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *EnableScaleOutExtentMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
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
