package handlers

import (
	"github.com/Car-parking-management-systems"
	"github.com/Car-parking-management-systems/gen/models"
	"github.com/Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetCarByID handles request for retrieving a car by ID
func GetCarByID(rt *runtime.Runtime) operations.GetCarHandler {
	return &getCarByID{rt: rt}
}

type getCarByID struct {
	rt *runtime.Runtime
}

// Handle the get car by ID request
func (d *getCarByID) Handle(params operations.GetCarParams) middleware.Responder {
	car, err := d.rt.Service().GetCarByID(params.ID)
	if err != nil {
		log().Errorf("failed to get car by ID: %s", err)
		return operations.NewGetCarNotFound()
	}

	payload := &models.Car{
		ID:           car.ID,
		LicensePlate: car.LicensePlate,
		Model:        car.Model,
		Brand:        car.Brand,
		Color:        car.Color,
		PlateNumber:  car.PlateNumber,
	}

	return operations.NewGetCarOK().WithPayload(payload)
}
