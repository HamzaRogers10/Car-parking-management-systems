package models_test

import (
	"Car-parking-management-systems/models"
	"github.com/go-openapi/strfmt"
	"testing"
	"time"
)

// CalculateTotalAmount calculates the total parking fee based on the check-in time, check-out time, and car type.
// It returns the total amount and any error that occurred during the calculation.
func TestCalculateTotalAmount(t *testing.T) {
	car := &models.Car{PlateNumber: "ABC123", Color: "Blue"}
	entryTime := time.Date(2022, time.March, 27, 14, 0, 0, 0, time.UTC)
	exitTime := time.Date(2022, time.March, 27, 18, 0, 0, 0, time.UTC)
	hourlyRate := 10.0
	expectedTotalAmount := 40.0

	parking := &models.Parking{Car: *car, EntryTime: time.Time(strfmt.DateTime(entryTime)), ExitTime: exitTime, HourlyRate: hourlyRate}
	parking.CalculateTotalAmount()

	if parking.TotalAmount != expectedTotalAmount {
		t.Errorf("Expected total amount to be %f, but got %f", expectedTotalAmount, parking.TotalAmount)
	}
}
