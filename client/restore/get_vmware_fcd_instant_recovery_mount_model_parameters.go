// Code generated by go-swagger; DO NOT EDIT.

package restore

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

// NewGetVmwareFcdInstantRecoveryMountModelParams creates a new GetVmwareFcdInstantRecoveryMountModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVmwareFcdInstantRecoveryMountModelParams() *GetVmwareFcdInstantRecoveryMountModelParams {
	return &GetVmwareFcdInstantRecoveryMountModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVmwareFcdInstantRecoveryMountModelParamsWithTimeout creates a new GetVmwareFcdInstantRecoveryMountModelParams object
// with the ability to set a timeout on a request.
func NewGetVmwareFcdInstantRecoveryMountModelParamsWithTimeout(timeout time.Duration) *GetVmwareFcdInstantRecoveryMountModelParams {
	return &GetVmwareFcdInstantRecoveryMountModelParams{
		timeout: timeout,
	}
}

// NewGetVmwareFcdInstantRecoveryMountModelParamsWithContext creates a new GetVmwareFcdInstantRecoveryMountModelParams object
// with the ability to set a context for a request.
func NewGetVmwareFcdInstantRecoveryMountModelParamsWithContext(ctx context.Context) *GetVmwareFcdInstantRecoveryMountModelParams {
	return &GetVmwareFcdInstantRecoveryMountModelParams{
		Context: ctx,
	}
}

// NewGetVmwareFcdInstantRecoveryMountModelParamsWithHTTPClient creates a new GetVmwareFcdInstantRecoveryMountModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVmwareFcdInstantRecoveryMountModelParamsWithHTTPClient(client *http.Client) *GetVmwareFcdInstantRecoveryMountModelParams {
	return &GetVmwareFcdInstantRecoveryMountModelParams{
		HTTPClient: client,
	}
}

/* GetVmwareFcdInstantRecoveryMountModelParams contains all the parameters to send to the API endpoint
   for the get vmware fcd instant recovery mount model operation.

   Typically these are written to a http.Request.
*/
type GetVmwareFcdInstantRecoveryMountModelParams struct {

	/* MountID.

	   Mount ID.

	   Format: uuid
	*/
	MountID strfmt.UUID

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

// WithDefaults hydrates default values in the get vmware fcd instant recovery mount model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithDefaults() *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vmware fcd instant recovery mount model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := GetVmwareFcdInstantRecoveryMountModelParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithTimeout(timeout time.Duration) *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithContext(ctx context.Context) *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithHTTPClient(client *http.Client) *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMountID adds the mountID to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithMountID(mountID strfmt.UUID) *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetMountID(mountID)
	return o
}

// SetMountID adds the mountId to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetMountID(mountID strfmt.UUID) {
	o.MountID = mountID
}

// WithXAPIVersion adds the xAPIVersion to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WithXAPIVersion(xAPIVersion string) *GetVmwareFcdInstantRecoveryMountModelParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the get vmware fcd instant recovery mount model params
func (o *GetVmwareFcdInstantRecoveryMountModelParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetVmwareFcdInstantRecoveryMountModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mountId
	if err := r.SetPathParam("mountId", o.MountID.String()); err != nil {
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
