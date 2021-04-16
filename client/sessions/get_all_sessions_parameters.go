// Code generated by go-swagger; DO NOT EDIT.

package sessions

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

// NewGetAllSessionsParams creates a new GetAllSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllSessionsParams() *GetAllSessionsParams {
	return &GetAllSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSessionsParamsWithTimeout creates a new GetAllSessionsParams object
// with the ability to set a timeout on a request.
func NewGetAllSessionsParamsWithTimeout(timeout time.Duration) *GetAllSessionsParams {
	return &GetAllSessionsParams{
		timeout: timeout,
	}
}

// NewGetAllSessionsParamsWithContext creates a new GetAllSessionsParams object
// with the ability to set a context for a request.
func NewGetAllSessionsParamsWithContext(ctx context.Context) *GetAllSessionsParams {
	return &GetAllSessionsParams{
		Context: ctx,
	}
}

// NewGetAllSessionsParamsWithHTTPClient creates a new GetAllSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllSessionsParamsWithHTTPClient(client *http.Client) *GetAllSessionsParams {
	return &GetAllSessionsParams{
		HTTPClient: client,
	}
}

/* GetAllSessionsParams contains all the parameters to send to the API endpoint
   for the get all sessions operation.

   Typically these are written to a http.Request.
*/
type GetAllSessionsParams struct {

	/* CreatedAfterFilter.

	   Returns sessions that are created after the specified date and time.

	   Format: date-time
	*/
	CreatedAfterFilter *strfmt.DateTime

	/* CreatedBeforeFilter.

	   Returns sessions that are created before the specified date and time.

	   Format: date-time
	*/
	CreatedBeforeFilter *strfmt.DateTime

	/* EndedAfterFilter.

	   Returns sessions that are finished after the specified date and time.

	   Format: date-time
	*/
	EndedAfterFilter *strfmt.DateTime

	/* EndedBeforeFilter.

	   Returns sessions that are finished before the specified date and time.

	   Format: date-time
	*/
	EndedBeforeFilter *strfmt.DateTime

	/* JobIDFilter.

	   Filters sessions by job ID.

	   Format: uuid
	*/
	JobIDFilter *strfmt.UUID

	/* Limit.

	   Maximum number of sessions to return.

	   Format: int32
	*/
	Limit *int32

	/* NameFilter.

	   Filters sessions by the `nameFilter` pattern. The pattern can match any session parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	*/
	NameFilter *string

	/* OrderAsc.

	   Sorts sessions in the ascending order by the `orderColumn` parameter.
	*/
	OrderAsc *bool

	/* OrderColumn.

	   Sorts sessions by one of the session parameters.
	*/
	OrderColumn *string

	/* ResultFilter.

	   Filters sessions by session result.
	*/
	ResultFilter *string

	/* Skip.

	   Number of sessions to skip.

	   Format: int32
	*/
	Skip *int32

	/* StateFilter.

	   Filters sessions by session state.
	*/
	StateFilter *string

	/* TypeFilter.

	   Filters sessions by session type.
	*/
	TypeFilter *string

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

// WithDefaults hydrates default values in the get all sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSessionsParams) WithDefaults() *GetAllSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSessionsParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := GetAllSessionsParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all sessions params
func (o *GetAllSessionsParams) WithTimeout(timeout time.Duration) *GetAllSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all sessions params
func (o *GetAllSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all sessions params
func (o *GetAllSessionsParams) WithContext(ctx context.Context) *GetAllSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all sessions params
func (o *GetAllSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all sessions params
func (o *GetAllSessionsParams) WithHTTPClient(client *http.Client) *GetAllSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all sessions params
func (o *GetAllSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAfterFilter adds the createdAfterFilter to the get all sessions params
func (o *GetAllSessionsParams) WithCreatedAfterFilter(createdAfterFilter *strfmt.DateTime) *GetAllSessionsParams {
	o.SetCreatedAfterFilter(createdAfterFilter)
	return o
}

// SetCreatedAfterFilter adds the createdAfterFilter to the get all sessions params
func (o *GetAllSessionsParams) SetCreatedAfterFilter(createdAfterFilter *strfmt.DateTime) {
	o.CreatedAfterFilter = createdAfterFilter
}

// WithCreatedBeforeFilter adds the createdBeforeFilter to the get all sessions params
func (o *GetAllSessionsParams) WithCreatedBeforeFilter(createdBeforeFilter *strfmt.DateTime) *GetAllSessionsParams {
	o.SetCreatedBeforeFilter(createdBeforeFilter)
	return o
}

// SetCreatedBeforeFilter adds the createdBeforeFilter to the get all sessions params
func (o *GetAllSessionsParams) SetCreatedBeforeFilter(createdBeforeFilter *strfmt.DateTime) {
	o.CreatedBeforeFilter = createdBeforeFilter
}

// WithEndedAfterFilter adds the endedAfterFilter to the get all sessions params
func (o *GetAllSessionsParams) WithEndedAfterFilter(endedAfterFilter *strfmt.DateTime) *GetAllSessionsParams {
	o.SetEndedAfterFilter(endedAfterFilter)
	return o
}

// SetEndedAfterFilter adds the endedAfterFilter to the get all sessions params
func (o *GetAllSessionsParams) SetEndedAfterFilter(endedAfterFilter *strfmt.DateTime) {
	o.EndedAfterFilter = endedAfterFilter
}

// WithEndedBeforeFilter adds the endedBeforeFilter to the get all sessions params
func (o *GetAllSessionsParams) WithEndedBeforeFilter(endedBeforeFilter *strfmt.DateTime) *GetAllSessionsParams {
	o.SetEndedBeforeFilter(endedBeforeFilter)
	return o
}

// SetEndedBeforeFilter adds the endedBeforeFilter to the get all sessions params
func (o *GetAllSessionsParams) SetEndedBeforeFilter(endedBeforeFilter *strfmt.DateTime) {
	o.EndedBeforeFilter = endedBeforeFilter
}

// WithJobIDFilter adds the jobIDFilter to the get all sessions params
func (o *GetAllSessionsParams) WithJobIDFilter(jobIDFilter *strfmt.UUID) *GetAllSessionsParams {
	o.SetJobIDFilter(jobIDFilter)
	return o
}

// SetJobIDFilter adds the jobIdFilter to the get all sessions params
func (o *GetAllSessionsParams) SetJobIDFilter(jobIDFilter *strfmt.UUID) {
	o.JobIDFilter = jobIDFilter
}

// WithLimit adds the limit to the get all sessions params
func (o *GetAllSessionsParams) WithLimit(limit *int32) *GetAllSessionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all sessions params
func (o *GetAllSessionsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNameFilter adds the nameFilter to the get all sessions params
func (o *GetAllSessionsParams) WithNameFilter(nameFilter *string) *GetAllSessionsParams {
	o.SetNameFilter(nameFilter)
	return o
}

// SetNameFilter adds the nameFilter to the get all sessions params
func (o *GetAllSessionsParams) SetNameFilter(nameFilter *string) {
	o.NameFilter = nameFilter
}

// WithOrderAsc adds the orderAsc to the get all sessions params
func (o *GetAllSessionsParams) WithOrderAsc(orderAsc *bool) *GetAllSessionsParams {
	o.SetOrderAsc(orderAsc)
	return o
}

// SetOrderAsc adds the orderAsc to the get all sessions params
func (o *GetAllSessionsParams) SetOrderAsc(orderAsc *bool) {
	o.OrderAsc = orderAsc
}

// WithOrderColumn adds the orderColumn to the get all sessions params
func (o *GetAllSessionsParams) WithOrderColumn(orderColumn *string) *GetAllSessionsParams {
	o.SetOrderColumn(orderColumn)
	return o
}

// SetOrderColumn adds the orderColumn to the get all sessions params
func (o *GetAllSessionsParams) SetOrderColumn(orderColumn *string) {
	o.OrderColumn = orderColumn
}

// WithResultFilter adds the resultFilter to the get all sessions params
func (o *GetAllSessionsParams) WithResultFilter(resultFilter *string) *GetAllSessionsParams {
	o.SetResultFilter(resultFilter)
	return o
}

// SetResultFilter adds the resultFilter to the get all sessions params
func (o *GetAllSessionsParams) SetResultFilter(resultFilter *string) {
	o.ResultFilter = resultFilter
}

// WithSkip adds the skip to the get all sessions params
func (o *GetAllSessionsParams) WithSkip(skip *int32) *GetAllSessionsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get all sessions params
func (o *GetAllSessionsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithStateFilter adds the stateFilter to the get all sessions params
func (o *GetAllSessionsParams) WithStateFilter(stateFilter *string) *GetAllSessionsParams {
	o.SetStateFilter(stateFilter)
	return o
}

// SetStateFilter adds the stateFilter to the get all sessions params
func (o *GetAllSessionsParams) SetStateFilter(stateFilter *string) {
	o.StateFilter = stateFilter
}

// WithTypeFilter adds the typeFilter to the get all sessions params
func (o *GetAllSessionsParams) WithTypeFilter(typeFilter *string) *GetAllSessionsParams {
	o.SetTypeFilter(typeFilter)
	return o
}

// SetTypeFilter adds the typeFilter to the get all sessions params
func (o *GetAllSessionsParams) SetTypeFilter(typeFilter *string) {
	o.TypeFilter = typeFilter
}

// WithXAPIVersion adds the xAPIVersion to the get all sessions params
func (o *GetAllSessionsParams) WithXAPIVersion(xAPIVersion string) *GetAllSessionsParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the get all sessions params
func (o *GetAllSessionsParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EndedAfterFilter != nil {

		// query param endedAfterFilter
		var qrEndedAfterFilter strfmt.DateTime

		if o.EndedAfterFilter != nil {
			qrEndedAfterFilter = *o.EndedAfterFilter
		}
		qEndedAfterFilter := qrEndedAfterFilter.String()
		if qEndedAfterFilter != "" {

			if err := r.SetQueryParam("endedAfterFilter", qEndedAfterFilter); err != nil {
				return err
			}
		}
	}

	if o.EndedBeforeFilter != nil {

		// query param endedBeforeFilter
		var qrEndedBeforeFilter strfmt.DateTime

		if o.EndedBeforeFilter != nil {
			qrEndedBeforeFilter = *o.EndedBeforeFilter
		}
		qEndedBeforeFilter := qrEndedBeforeFilter.String()
		if qEndedBeforeFilter != "" {

			if err := r.SetQueryParam("endedBeforeFilter", qEndedBeforeFilter); err != nil {
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

	if o.ResultFilter != nil {

		// query param resultFilter
		var qrResultFilter string

		if o.ResultFilter != nil {
			qrResultFilter = *o.ResultFilter
		}
		qResultFilter := qrResultFilter
		if qResultFilter != "" {

			if err := r.SetQueryParam("resultFilter", qResultFilter); err != nil {
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

	if o.StateFilter != nil {

		// query param stateFilter
		var qrStateFilter string

		if o.StateFilter != nil {
			qrStateFilter = *o.StateFilter
		}
		qStateFilter := qrStateFilter
		if qStateFilter != "" {

			if err := r.SetQueryParam("stateFilter", qStateFilter); err != nil {
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

	// header param x-api-version
	if err := r.SetHeaderParam("x-api-version", o.XAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
