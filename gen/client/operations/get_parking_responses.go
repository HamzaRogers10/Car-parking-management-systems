// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Car-parking-management-systems/gen/models"
)

// GetParkingReader is a Reader for the GetParking structure.
type GetParkingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParkingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParkingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetParkingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetParkingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetParkingOK creates a GetParkingOK with default headers values
func NewGetParkingOK() *GetParkingOK {
	return &GetParkingOK{}
}

/*
GetParkingOK describes a response with status code 200, with default header values.

parking response
*/
type GetParkingOK struct {
	Payload *models.Parking
}

// IsSuccess returns true when this get parking o k response has a 2xx status code
func (o *GetParkingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get parking o k response has a 3xx status code
func (o *GetParkingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parking o k response has a 4xx status code
func (o *GetParkingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get parking o k response has a 5xx status code
func (o *GetParkingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get parking o k response a status code equal to that given
func (o *GetParkingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetParkingOK) Error() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingOK  %+v", 200, o.Payload)
}

func (o *GetParkingOK) String() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingOK  %+v", 200, o.Payload)
}

func (o *GetParkingOK) GetPayload() *models.Parking {
	return o.Payload
}

func (o *GetParkingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Parking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParkingNotFound creates a GetParkingNotFound with default headers values
func NewGetParkingNotFound() *GetParkingNotFound {
	return &GetParkingNotFound{}
}

/*
GetParkingNotFound describes a response with status code 404, with default header values.

parking not found
*/
type GetParkingNotFound struct {
}

// IsSuccess returns true when this get parking not found response has a 2xx status code
func (o *GetParkingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parking not found response has a 3xx status code
func (o *GetParkingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parking not found response has a 4xx status code
func (o *GetParkingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parking not found response has a 5xx status code
func (o *GetParkingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get parking not found response a status code equal to that given
func (o *GetParkingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetParkingNotFound) Error() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingNotFound ", 404)
}

func (o *GetParkingNotFound) String() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingNotFound ", 404)
}

func (o *GetParkingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParkingInternalServerError creates a GetParkingInternalServerError with default headers values
func NewGetParkingInternalServerError() *GetParkingInternalServerError {
	return &GetParkingInternalServerError{}
}

/*
GetParkingInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetParkingInternalServerError struct {
}

// IsSuccess returns true when this get parking internal server error response has a 2xx status code
func (o *GetParkingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parking internal server error response has a 3xx status code
func (o *GetParkingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parking internal server error response has a 4xx status code
func (o *GetParkingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get parking internal server error response has a 5xx status code
func (o *GetParkingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get parking internal server error response a status code equal to that given
func (o *GetParkingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetParkingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingInternalServerError ", 500)
}

func (o *GetParkingInternalServerError) String() string {
	return fmt.Sprintf("[GET /parkings/{id}][%d] getParkingInternalServerError ", 500)
}

func (o *GetParkingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}