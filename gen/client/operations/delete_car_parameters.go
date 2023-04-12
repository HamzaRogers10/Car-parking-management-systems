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

// NewDeleteCarParams creates a new DeleteCarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCarParams() *DeleteCarParams {
	return &DeleteCarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCarParamsWithTimeout creates a new DeleteCarParams object
// with the ability to set a timeout on a request.
func NewDeleteCarParamsWithTimeout(timeout time.Duration) *DeleteCarParams {
	return &DeleteCarParams{
		timeout: timeout,
	}
}

// NewDeleteCarParamsWithContext creates a new DeleteCarParams object
// with the ability to set a context for a request.
func NewDeleteCarParamsWithContext(ctx context.Context) *DeleteCarParams {
	return &DeleteCarParams{
		Context: ctx,
	}
}

// NewDeleteCarParamsWithHTTPClient creates a new DeleteCarParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCarParamsWithHTTPClient(client *http.Client) *DeleteCarParams {
	return &DeleteCarParams{
		HTTPClient: client,
	}
}

/*
DeleteCarParams contains all the parameters to send to the API endpoint

	for the delete car operation.

	Typically these are written to a http.Request.
*/
type DeleteCarParams struct {

	/* ID.

	   ID of the car
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCarParams) WithDefaults() *DeleteCarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete car params
func (o *DeleteCarParams) WithTimeout(timeout time.Duration) *DeleteCarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete car params
func (o *DeleteCarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete car params
func (o *DeleteCarParams) WithContext(ctx context.Context) *DeleteCarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete car params
func (o *DeleteCarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete car params
func (o *DeleteCarParams) WithHTTPClient(client *http.Client) *DeleteCarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete car params
func (o *DeleteCarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete car params
func (o *DeleteCarParams) WithID(id string) *DeleteCarParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete car params
func (o *DeleteCarParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
