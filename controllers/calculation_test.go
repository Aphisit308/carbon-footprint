package controllers_test

import (
	"carbon-footprint/config"
	"carbon-footprint/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CalculatorTest(t *testing.T) {

	req := models.CarbonFootprintCalculatorRequest{
		ActivityType: "transportation",
		DistanceKm:   100,
		VehicleType:  "car",
		FuelType:     "gasoline",
	}

	var response models.CarbonFootprintCalculatorResponse
	var emissionFactor models.EmissionFactor
	err := config.DB.
		Table("emission_factors").
		Where(`
			emission_factors.activity_type = $1 AND
			emission_factors.vehicle_type = $2 AND 
			emission_factors.fuel_type = $3
		`, req.ActivityType, req.VehicleType, req.FuelType).
		First(&emissionFactor).Error
	assert.NoError(t, err)

	response.ActivityType = req.ActivityType
	response.CarbonEmissionKg = float64(emissionFactor.EmissionFactor) * req.DistanceKm
	assert.Equal(t, 3000, response.CarbonEmissionKg)
	assert.Equal(t, "transportation", response.ActivityType)

}
