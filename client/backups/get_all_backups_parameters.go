// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// NewGetAllBackupsParams creates a new GetAllBackupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllBackupsParams() *GetAllBackupsParams {
	return &GetAllBackupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllBackupsParamsWithTimeout creates a new GetAllBackupsParams object
// with the ability to set a timeout on a request.
func NewGetAllBackupsParamsWithTimeout(timeout time.Duration) *GetAllBackupsParams {
	return &GetAllBackupsParams{
		timeout: timeout,
	}
}

// NewGetAllBackupsParamsWithContext creates a new GetAllBackupsParams object
// with the ability to set a context for a request.
func NewGetAllBackupsParamsWithContext(ctx context.Context) *GetAllBackupsParams {
	return &GetAllBackupsParams{
		Context: ctx,
	}
}

// NewGetAllBackupsParamsWithHTTPClient creates a new GetAllBackupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllBackupsParamsWithHTTPClient(client *http.Client) *GetAllBackupsParams {
	return &GetAllBackupsParams{
		HTTPClient: client,
	}
}

/* GetAllBackupsParams contains all the parameters to send to the API endpoint
   for the get all backups operation.

   Typically these are written to a http.Request.
*/
type GetAllBackupsParams struct {

	/* CreatedAfterFilter.

	   Returns backups that are created after the specified date and time.

	   Format: date-time
	*/
	CreatedAfterFilter *strfmt.DateTime

	/* CreatedBeforeFilter.

	   Returns backups that are created before the specified date and time.

	   Format: date-time
	*/
	CreatedBeforeFilter *strfmt.DateTime

	/* JobIDFilter.

	   Filters backups by ID of the parent job.

	   Format: uuid
	*/
	JobIDFilter *strfmt.UUID

	/* Limit.

	   Maximum number of backups to return.

	   Format: int32
	*/
	Limit *int32

	/* NameFilter.

	   Filters backups by the `nameFilter` pattern. The pattern can match any backup parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	*/
	NameFilter *string

	/* OrderAsc.

	   Sorts backups in the ascending order by the `orderColumn` parameter.
	*/
	OrderAsc *bool

	/* OrderColumn.

	   Sorts backups by one of the backup parameters.
	*/
	OrderColumn *string

	/* PlatformIDFilter.

	   Filters backups by ID of the backup object platform.

	   Format: uuid
	*/
	PlatformIDFilter *strfmt.UUID

	/* PolicyTagFilter.

	   Filters backups by retention policy tag.
	*/
	PolicyTagFilter *string

	/* Skip.

	   Number of backups to skip.

	   Format: int32
	*/
	Skip *int32

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

// WithDefaults hydrates default values in the get all backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBackupsParams) WithDefaults() *GetAllBackupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBackupsParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := GetAllBackupsParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all backups params
func (o *GetAllBackupsParams) WithTimeout(timeout time.Duration) *GetAllBackupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all backups params
func (o *GetAllBackupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all backups params
func (o *GetAllBackupsParams) WithContext(ctx context.Context) *GetAllBackupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all backups params
func (o *GetAllBackupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all backups params
func (o *GetAllBackupsParams) WithHTTPClient(client *http.Client) *GetAllBackupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all backups params
func (o *GetAllBackupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAfterFilter adds the createdAfterFilter to the get all backups params
func (o *GetAllBackupsParams) WithCreatedAfterFilter(createdAfterFilter *strfmt.DateTime) *GetAllBackupsParams {
	o.SetCreatedAfterFilter(createdAfterFilter)
	return o
}

// SetCreatedAfterFilter adds the createdAfterFilter to the get all backups params
func (o *GetAllBackupsParams) SetCreatedAfterFilter(createdAfterFilter *strfmt.DateTime) {
	o.CreatedAfterFilter = createdAfterFilter
}

// WithCreatedBeforeFilter adds the createdBeforeFilter to the get all backups params
func (o *GetAllBackupsParams) WithCreatedBeforeFilter(createdBeforeFilter *strfmt.DateTime) *GetAllBackupsParams {
	o.SetCreatedBeforeFilter(createdBeforeFilter)
	return o
}

// SetCreatedBeforeFilter adds the createdBeforeFilter to the get all backups params
func (o *GetAllBackupsParams) SetCreatedBeforeFilter(createdBeforeFilter *strfmt.DateTime) {
	o.CreatedBeforeFilter = createdBeforeFilter
}

// WithJobIDFilter adds the jobIDFilter to the get all backups params
func (o *GetAllBackupsParams) WithJobIDFilter(jobIDFilter *strfmt.UUID) *GetAllBackupsParams {
	o.SetJobIDFilter(jobIDFilter)
	return o
}

// SetJobIDFilter adds the jobIdFilter to the get all backups params
func (o *GetAllBackupsParams) SetJobIDFilter(jobIDFilter *strfmt.UUID) {
	o.JobIDFilter = jobIDFilter
}

// WithLimit adds the limit to the get all backups params
func (o *GetAllBackupsParams) WithLimit(limit *int32) *GetAllBackupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all backups params
func (o *GetAllBackupsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNameFilter adds the nameFilter to the get all backups params
func (o *GetAllBackupsParams) WithNameFilter(nameFilter *string) *GetAllBackupsParams {
	o.SetNameFilter(nameFilter)
	return o
}

// SetNameFilter adds the nameFilter to the get all backups params
func (o *GetAllBackupsParams) SetNameFilter(nameFilter *string) {
	o.NameFilter = nameFilter
}

// WithOrderAsc adds the orderAsc to the get all backups params
func (o *GetAllBackupsParams) WithOrderAsc(orderAsc *bool) *GetAllBackupsParams {
	o.SetOrderAsc(orderAsc)
	return o
}

// SetOrderAsc adds the orderAsc to the get all backups params
func (o *GetAllBackupsParams) SetOrderAsc(orderAsc *bool) {
	o.OrderAsc = orderAsc
}

// WithOrderColumn adds the orderColumn to the get all backups params
func (o *GetAllBackupsParams) WithOrderColumn(orderColumn *string) *GetAllBackupsParams {
	o.SetOrderColumn(orderColumn)
	return o
}

// SetOrderColumn adds the orderColumn to the get all backups params
func (o *GetAllBackupsParams) SetOrderColumn(orderColumn *string) {
	o.OrderColumn = orderColumn
}

// WithPlatformIDFilter adds the platformIDFilter to the get all backups params
func (o *GetAllBackupsParams) WithPlatformIDFilter(platformIDFilter *strfmt.UUID) *GetAllBackupsParams {
	o.SetPlatformIDFilter(platformIDFilter)
	return o
}

// SetPlatformIDFilter adds the platformIdFilter to the get all backups params
func (o *GetAllBackupsParams) SetPlatformIDFilter(platformIDFilter *strfmt.UUID) {
	o.PlatformIDFilter = platformIDFilter
}

// WithPolicyTagFilter adds the policyTagFilter to the get all backups params
func (o *GetAllBackupsParams) WithPolicyTagFilter(policyTagFilter *string) *GetAllBackupsParams {
	o.SetPolicyTagFilter(policyTagFilter)
	return o
}

// SetPolicyTagFilter adds the policyTagFilter to the get all backups params
func (o *GetAllBackupsParams) SetPolicyTagFilter(policyTagFilter *string) {
	o.PolicyTagFilter = policyTagFilter
}

// WithSkip adds the skip to the get all backups params
func (o *GetAllBackupsParams) WithSkip(skip *int32) *GetAllBackupsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get all backups params
func (o *GetAllBackupsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithXAPIVersion adds the xAPIVersion to the get all backups params
func (o *GetAllBackupsParams) WithXAPIVersion(xAPIVersion string) *GetAllBackupsParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the get all backups params
func (o *GetAllBackupsParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllBackupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedAfterFilter != nil {

		// query param createdAfterFilter
		var qrCreatedAfterFilter strfmt.DateTime

		if o.CreatedAfterFilter != nil {
			qrCreatedAfterFilter = *o.CreatedAfterFilter
		}
		qCreatedAfterFilter := qrCreatedAfterFilter.String()
		if qCreatedAfterFilter != "" {

			if err := r.SetQueryParam("createdAfterFilter", qCreatedAfterFilter); err != nil {
				return err
			}
		}
	}

	if o.CreatedBeforeFilter != nil {

		// query param createdBeforeFilter
		var qrCreatedBeforeFilter strfmt.DateTime

		if o.CreatedBeforeFilter != nil {
			qrCreatedBeforeFilter = *o.CreatedBeforeFilter
		}
		qCreatedBeforeFilter := qrCreatedBeforeFilter.String()
		if qCreatedBeforeFilter != "" {

			if err := r.SetQueryParam("createdBeforeFilter", qCreatedBeforeFilter); err != nil {
				return err
			}
		}
	}

	if o.JobIDFilter != nil {

		// query param jobIdFilter
		var qrJobIDFilter strfmt.UUID

		if o.JobIDFilter != nil {
			qrJobIDFilter = *o.JobIDFilter
		}
		qJobIDFilter := qrJobIDFilter.String()
		if qJobIDFilter != "" {

			if err := r.SetQueryParam("jobIdFilter", qJobIDFilter); err != nil {
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

	if o.PlatformIDFilter != nil {

		// query param platformIdFilter
		var qrPlatformIDFilter strfmt.UUID

		if o.PlatformIDFilter != nil {
			qrPlatformIDFilter = *o.PlatformIDFilter
		}
		qPlatformIDFilter := qrPlatformIDFilter.String()
		if qPlatformIDFilter != "" {

			if err := r.SetQueryParam("platformIdFilter", qPlatformIDFilter); err != nil {
				return err
			}
		}
	}

	if o.PolicyTagFilter != nil {

		// query param policyTagFilter
		var qrPolicyTagFilter string

		if o.PolicyTagFilter != nil {
			qrPolicyTagFilter = *o.PolicyTagFilter
		}
		qPolicyTagFilter := qrPolicyTagFilter
		if qPolicyTagFilter != "" {

			if err := r.SetQueryParam("policyTagFilter", qPolicyTagFilter); err != nil {
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

	// header param x-api-version
	if err := r.SetHeaderParam("x-api-version", o.XAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}