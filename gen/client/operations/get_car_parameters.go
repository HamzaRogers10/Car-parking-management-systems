// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetCarParams creates a new GetCarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCarParams() *GetCarParams {
	return &GetCarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCarParamsWithTimeout creates a new GetCarParams object
// with the ability to set a timeout on a request.
func NewGetCarParamsWithTimeout(timeout time.Duration) *GetCarParams {
	return &GetCarParams{
		timeout: timeout,
	}
}

// NewGetCarParamsWithContext creates a new GetCarParams object
// with the ability to set a context for a request.
func NewGetCarParamsWithContext(ctx context.Context) *GetCarParams {
	return &GetCarParams{
		Context: ctx,
	}
}

// NewGetCarParamsWithHTTPClient creates a new GetCarParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCarParamsWithHTTPClient(client *http.Client) *GetCarParams {
	return &GetCarParams{
		HTTPClient: client,
	}
}

/*
GetCarParams contains all the parameters to send to the API endpoint

	for the get car operation.

	Typically these are written to a http.Request.
*/
type GetCarParams struct {

	/* ID.

	   ID of the car
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCarParams) WithDefaults() *GetCarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get car params
func (o *GetCarParams) WithTimeout(timeout time.Duration) *GetCarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get car params
func (o *GetCarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get car params
func (o *GetCarParams) WithContext(ctx context.Context) *GetCarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get car params
func (o *GetCarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get car params
func (o *GetCarParams) WithHTTPClient(client *http.Client) *GetCarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get car params
func (o *GetCarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get car params
func (o *GetCarParams) WithID(id string) *GetCarParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get car params
func (o *GetCarParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}