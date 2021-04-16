// Code generated by go-swagger; DO NOT EDIT.

package login

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

// NewCreateAuthorizationCodeParams creates a new CreateAuthorizationCodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAuthorizationCodeParams() *CreateAuthorizationCodeParams {
	return &CreateAuthorizationCodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthorizationCodeParamsWithTimeout creates a new CreateAuthorizationCodeParams object
// with the ability to set a timeout on a request.
func NewCreateAuthorizationCodeParamsWithTimeout(timeout time.Duration) *CreateAuthorizationCodeParams {
	return &CreateAuthorizationCodeParams{
		timeout: timeout,
	}
}

// NewCreateAuthorizationCodeParamsWithContext creates a new CreateAuthorizationCodeParams object
// with the ability to set a context for a request.
func NewCreateAuthorizationCodeParamsWithContext(ctx context.Context) *CreateAuthorizationCodeParams {
	return &CreateAuthorizationCodeParams{
		Context: ctx,
	}
}

// NewCreateAuthorizationCodeParamsWithHTTPClient creates a new CreateAuthorizationCodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAuthorizationCodeParamsWithHTTPClient(client *http.Client) *CreateAuthorizationCodeParams {
	return &CreateAuthorizationCodeParams{
		HTTPClient: client,
	}
}

/* CreateAuthorizationCodeParams contains all the parameters to send to the API endpoint
   for the create authorization code operation.

   Typically these are written to a http.Request.
*/
type CreateAuthorizationCodeParams struct {

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

// WithDefaults hydrates default values in the create authorization code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthorizationCodeParams) WithDefaults() *CreateAuthorizationCodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create authorization code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthorizationCodeParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := CreateAuthorizationCodeParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create authorization code params
func (o *CreateAuthorizationCodeParams) WithTimeout(timeout time.Duration) *CreateAuthorizationCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create authorization code params
func (o *CreateAuthorizationCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create authorization code params
func (o *CreateAuthorizationCodeParams) WithContext(ctx context.Context) *CreateAuthorizationCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create authorization code params
func (o *CreateAuthorizationCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create authorization code params
func (o *CreateAuthorizationCodeParams) WithHTTPClient(client *http.Client) *CreateAuthorizationCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create authorization code params
func (o *CreateAuthorizationCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAPIVersion adds the xAPIVersion to the create authorization code params
func (o *CreateAuthorizationCodeParams) WithXAPIVersion(xAPIVersion string) *CreateAuthorizationCodeParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the create authorization code params
func (o *CreateAuthorizationCodeParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthorizationCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
