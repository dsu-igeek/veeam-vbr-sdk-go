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
	"github.com/go-openapi/swag"
)

// NewDeleteRepositoryParams creates a new DeleteRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRepositoryParams() *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoryParamsWithTimeout creates a new DeleteRepositoryParams object
// with the ability to set a timeout on a request.
func NewDeleteRepositoryParamsWithTimeout(timeout time.Duration) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		timeout: timeout,
	}
}

// NewDeleteRepositoryParamsWithContext creates a new DeleteRepositoryParams object
// with the ability to set a context for a request.
func NewDeleteRepositoryParamsWithContext(ctx context.Context) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		Context: ctx,
	}
}

// NewDeleteRepositoryParamsWithHTTPClient creates a new DeleteRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRepositoryParamsWithHTTPClient(client *http.Client) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		HTTPClient: client,
	}
}

/* DeleteRepositoryParams contains all the parameters to send to the API endpoint
   for the delete repository operation.

   Typically these are written to a http.Request.
*/
type DeleteRepositoryParams struct {

	/* DeleteBackups.

	   If *true*, Veeam Backup & Replication will remove backup files.
	*/
	DeleteBackups *bool

	/* ID.

	   ID of the backup repository.

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

// WithDefaults hydrates default values in the delete repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRepositoryParams) WithDefaults() *DeleteRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRepositoryParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := DeleteRepositoryParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete repository params
func (o *DeleteRepositoryParams) WithTimeout(timeout time.Duration) *DeleteRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repository params
func (o *DeleteRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repository params
func (o *DeleteRepositoryParams) WithContext(ctx context.Context) *DeleteRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repository params
func (o *DeleteRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repository params
func (o *DeleteRepositoryParams) WithHTTPClient(client *http.Client) *DeleteRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repository params
func (o *DeleteRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteBackups adds the deleteBackups to the delete repository params
func (o *DeleteRepositoryParams) WithDeleteBackups(deleteBackups *bool) *DeleteRepositoryParams {
	o.SetDeleteBackups(deleteBackups)
	return o
}

// SetDeleteBackups adds the deleteBackups to the delete repository params
func (o *DeleteRepositoryParams) SetDeleteBackups(deleteBackups *bool) {
	o.DeleteBackups = deleteBackups
}

// WithID adds the id to the delete repository params
func (o *DeleteRepositoryParams) WithID(id strfmt.UUID) *DeleteRepositoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete repository params
func (o *DeleteRepositoryParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithXAPIVersion adds the xAPIVersion to the delete repository params
func (o *DeleteRepositoryParams) WithXAPIVersion(xAPIVersion string) *DeleteRepositoryParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the delete repository params
func (o *DeleteRepositoryParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteBackups != nil {

		// query param deleteBackups
		var qrDeleteBackups bool

		if o.DeleteBackups != nil {
			qrDeleteBackups = *o.DeleteBackups
		}
		qDeleteBackups := swag.FormatBool(qrDeleteBackups)
		if qDeleteBackups != "" {

			if err := r.SetQueryParam("deleteBackups", qDeleteBackups); err != nil {
				return err
			}
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