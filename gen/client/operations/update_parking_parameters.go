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

	"Car-parking-management-systems/gen/models"
)

// NewUpdateParkingParams creates a new UpdateParkingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateParkingParams() *UpdateParkingParams {
	return &UpdateParkingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParkingParamsWithTimeout creates a new UpdateParkingParams object
// with the ability to set a timeout on a request.
func NewUpdateParkingParamsWithTimeout(timeout time.Duration) *UpdateParkingParams {
	return &UpdateParkingParams{
		timeout: timeout,
	}
}

// NewUpdateParkingParamsWithContext creates a new UpdateParkingParams object
// with the ability to set a context for a request.
func NewUpdateParkingParamsWithContext(ctx context.Context) *UpdateParkingParams {
	return &UpdateParkingParams{
		Context: ctx,
	}
}

// NewUpdateParkingParamsWithHTTPClient creates a new UpdateParkingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateParkingParamsWithHTTPClient(client *http.Client) *UpdateParkingParams {
	return &UpdateParkingParams{
		HTTPClient: client,
	}
}

/*
UpdateParkingParams contains all the parameters to send to the API endpoint

	for the update parking operation.

	Typically these are written to a http.Request.
*/
type UpdateParkingParams struct {

	/* ID.

	   ID of the parking
	*/
	ID string

	/* Parking.

	   parking's details
	*/
	Parking *models.Parking

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update parking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParkingParams) WithDefaults() *UpdateParkingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update parking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParkingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update parking params
func (o *UpdateParkingParams) WithTimeout(timeout time.Duration) *UpdateParkingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update parking params
func (o *UpdateParkingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update parking params
func (o *UpdateParkingParams) WithContext(ctx context.Context) *UpdateParkingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update parking params
func (o *UpdateParkingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update parking params
func (o *UpdateParkingParams) WithHTTPClient(client *http.Client) *UpdateParkingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update parking params
func (o *UpdateParkingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update parking params
func (o *UpdateParkingParams) WithID(id string) *UpdateParkingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update parking params
func (o *UpdateParkingParams) SetID(id string) {
	o.ID = id
}

// WithParking adds the parking to the update parking params
func (o *UpdateParkingParams) WithParking(parking *models.Parking) *UpdateParkingParams {
	o.SetParking(parking)
	return o
}

// SetParking adds the parking to the update parking params
func (o *UpdateParkingParams) SetParking(parking *models.Parking) {
	o.Parking = parking
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParkingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Parking != nil {
		if err := r.SetBodyParam(o.Parking); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}