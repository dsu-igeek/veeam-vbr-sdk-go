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

// NewGetAllRepositoriesParams creates a new GetAllRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllRepositoriesParams() *GetAllRepositoriesParams {
	return &GetAllRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRepositoriesParamsWithTimeout creates a new GetAllRepositoriesParams object
// with the ability to set a timeout on a request.
func NewGetAllRepositoriesParamsWithTimeout(timeout time.Duration) *GetAllRepositoriesParams {
	return &GetAllRepositoriesParams{
		timeout: timeout,
	}
}

// NewGetAllRepositoriesParamsWithContext creates a new GetAllRepositoriesParams object
// with the ability to set a context for a request.
func NewGetAllRepositoriesParamsWithContext(ctx context.Context) *GetAllRepositoriesParams {
	return &GetAllRepositoriesParams{
		Context: ctx,
	}
}

// NewGetAllRepositoriesParamsWithHTTPClient creates a new GetAllRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllRepositoriesParamsWithHTTPClient(client *http.Client) *GetAllRepositoriesParams {
	return &GetAllRepositoriesParams{
		HTTPClient: client,
	}
}

/* GetAllRepositoriesParams contains all the parameters to send to the API endpoint
   for the get all repositories operation.

   Typically these are written to a http.Request.
*/
type GetAllRepositoriesParams struct {

	/* HostIDFilter.

	   Filters repositories by ID of the backup server.

	   Format: uuid
	*/
	HostIDFilter *strfmt.UUID

	/* Limit.

	   Maximum number of repositories to return.

	   Format: int32
	*/
	Limit *int32

	/* NameFilter.

	   Filters repositories by the `nameFilter` pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	*/
	NameFilter *string

	/* OrderAsc.

	   Sorts repositories in the ascending order by the `orderColumn` parameter.
	*/
	OrderAsc *bool

	/* OrderColumn.

	   Sorts repositories by one of the repository parameters.
	*/
	OrderColumn *string

	/* PathFilter.

	   Filters repositories by path to the folder where backup files are stored.
	*/
	PathFilter *string

	/* Skip.

	   Number of repositories to skip.

	   Format: int32
	*/
	Skip *int32

	/* TypeFilter.

	   Filters repositories by repository type.
	*/
	TypeFilter *string

	/* VmbAPIFilter.

	     Filters repositories by VM Backup API parameters converted to the base64 string.<br>
	To compose the base64 string:<br>
	<ol>
	  <li>Prepare VM Backup API parameters in the JSON format.</li>
	  <code>{<br>
	  "protocolVersion":"string",<br>
	  "assemblyVersion":"string",<br>
	  "productId":"string",<br>
	  "versionFlags":"string"<br>
	  }<br></code>
	  <li>Convert the JSON object into a string.</li>
	  <li>Encode the string with base64 encoding.</li>
	</ol>

	*/
	VmbAPIFilter *string

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

// WithDefaults hydrates default values in the get all repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRepositoriesParams) WithDefaults() *GetAllRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRepositoriesParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := GetAllRepositoriesParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all repositories params
func (o *GetAllRepositoriesParams) WithTimeout(timeout time.Duration) *GetAllRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all repositories params
func (o *GetAllRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all repositories params
func (o *GetAllRepositoriesParams) WithContext(ctx context.Context) *GetAllRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all repositories params
func (o *GetAllRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all repositories params
func (o *GetAllRepositoriesParams) WithHTTPClient(client *http.Client) *GetAllRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all repositories params
func (o *GetAllRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostIDFilter adds the hostIDFilter to the get all repositories params
func (o *GetAllRepositoriesParams) WithHostIDFilter(hostIDFilter *strfmt.UUID) *GetAllRepositoriesParams {
	o.SetHostIDFilter(hostIDFilter)
	return o
}

// SetHostIDFilter adds the hostIdFilter to the get all repositories params
func (o *GetAllRepositoriesParams) SetHostIDFilter(hostIDFilter *strfmt.UUID) {
	o.HostIDFilter = hostIDFilter
}

// WithLimit adds the limit to the get all repositories params
func (o *GetAllRepositoriesParams) WithLimit(limit *int32) *GetAllRepositoriesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all repositories params
func (o *GetAllRepositoriesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNameFilter adds the nameFilter to the get all repositories params
func (o *GetAllRepositoriesParams) WithNameFilter(nameFilter *string) *GetAllRepositoriesParams {
	o.SetNameFilter(nameFilter)
	return o
}

// SetNameFilter adds the nameFilter to the get all repositories params
func (o *GetAllRepositoriesParams) SetNameFilter(nameFilter *string) {
	o.NameFilter = nameFilter
}

// WithOrderAsc adds the orderAsc to the get all repositories params
func (o *GetAllRepositoriesParams) WithOrderAsc(orderAsc *bool) *GetAllRepositoriesParams {
	o.SetOrderAsc(orderAsc)
	return o
}

// SetOrderAsc adds the orderAsc to the get all repositories params
func (o *GetAllRepositoriesParams) SetOrderAsc(orderAsc *bool) {
	o.OrderAsc = orderAsc
}

// WithOrderColumn adds the orderColumn to the get all repositories params
func (o *GetAllRepositoriesParams) WithOrderColumn(orderColumn *string) *GetAllRepositoriesParams {
	o.SetOrderColumn(orderColumn)
	return o
}

// SetOrderColumn adds the orderColumn to the get all repositories params
func (o *GetAllRepositoriesParams) SetOrderColumn(orderColumn *string) {
	o.OrderColumn = orderColumn
}

// WithPathFilter adds the pathFilter to the get all repositories params
func (o *GetAllRepositoriesParams) WithPathFilter(pathFilter *string) *GetAllRepositoriesParams {
	o.SetPathFilter(pathFilter)
	return o
}

// SetPathFilter adds the pathFilter to the get all repositories params
func (o *GetAllRepositoriesParams) SetPathFilter(pathFilter *string) {
	o.PathFilter = pathFilter
}

// WithSkip adds the skip to the get all repositories params
func (o *GetAllRepositoriesParams) WithSkip(skip *int32) *GetAllRepositoriesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get all repositories params
func (o *GetAllRepositoriesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTypeFilter adds the typeFilter to the get all repositories params
func (o *GetAllRepositoriesParams) WithTypeFilter(typeFilter *string) *GetAllRepositoriesParams {
	o.SetTypeFilter(typeFilter)
	return o
}

// SetTypeFilter adds the typeFilter to the get all repositories params
func (o *GetAllRepositoriesParams) SetTypeFilter(typeFilter *string) {
	o.TypeFilter = typeFilter
}

// WithVmbAPIFilter adds the vmbAPIFilter to the get all repositories params
func (o *GetAllRepositoriesParams) WithVmbAPIFilter(vmbAPIFilter *string) *GetAllRepositoriesParams {
	o.SetVmbAPIFilter(vmbAPIFilter)
	return o
}

// SetVmbAPIFilter adds the vmbApiFilter to the get all repositories params
func (o *GetAllRepositoriesParams) SetVmbAPIFilter(vmbAPIFilter *string) {
	o.VmbAPIFilter = vmbAPIFilter
}

// WithXAPIVersion adds the xAPIVersion to the get all repositories params
func (o *GetAllRepositoriesParams) WithXAPIVersion(xAPIVersion string) *GetAllRepositoriesParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the get all repositories params
func (o *GetAllRepositoriesParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HostIDFilter != nil {

		// query param hostIdFilter
		var qrHostIDFilter strfmt.UUID

		if o.HostIDFilter != nil {
			qrHostIDFilter = *o.HostIDFilter
		}
		qHostIDFilter := qrHostIDFilter.String()
		if qHostIDFilter != "" {

			if err := r.SetQueryParam("hostIdFilter", qHostIDFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.NameFilter != nil {

		// query param nameFilter
		var qrNameFilter string

		if o.NameFilter != nil {
			qrNameFilter = *o.NameFilter
		}
		qNameFilter := qrNameFilter
		if qNameFilter != "" {

			if err := r.SetQueryParam("nameFilter", qNameFilter); err != nil {
				return err
			}
		}
	}

	if o.OrderAsc != nil {

		// query param orderAsc
		var qrOrderAsc bool

		if o.OrderAsc != nil {
			qrOrderAsc = *o.OrderAsc
		}
		qOrderAsc := swag.FormatBool(qrOrderAsc)
		if qOrderAsc != "" {

			if err := r.SetQueryParam("orderAsc", qOrderAsc); err != nil {
				return err
			}
		}
	}

	if o.OrderColumn != nil {

		// query param orderColumn
		var qrOrderColumn string

		if o.OrderColumn != nil {
			qrOrderColumn = *o.OrderColumn
		}
		qOrderColumn := qrOrderColumn
		if qOrderColumn != "" {

			if err := r.SetQueryParam("orderColumn", qOrderColumn); err != nil {
				return err
			}
		}
	}

	if o.PathFilter != nil {

		// query param pathFilter
		var qrPathFilter string

		if o.PathFilter != nil {
			qrPathFilter = *o.PathFilter
		}
		qPathFilter := qrPathFilter
		if qPathFilter != "" {

			if err := r.SetQueryParam("pathFilter", qPathFilter); err != nil {
				return err
			}
		}
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int32

		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {

			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}
	}

	if o.TypeFilter != nil {

		// query param typeFilter
		var qrTypeFilter string

		if o.TypeFilter != nil {
			qrTypeFilter = *o.TypeFilter
		}
		qTypeFilter := qrTypeFilter
		if qTypeFilter != "" {

			if err := r.SetQueryParam("typeFilter", qTypeFilter); err != nil {
				return err
			}
		}
	}

	if o.VmbAPIFilter != nil {

		// query param vmbApiFilter
		var qrVmbAPIFilter string

		if o.VmbAPIFilter != nil {
			qrVmbAPIFilter = *o.VmbAPIFilter
		}
		qVmbAPIFilter := qrVmbAPIFilter
		if qVmbAPIFilter != "" {

			if err := r.SetQueryParam("vmbApiFilter", qVmbAPIFilter); err != nil {
				return err
			}
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
