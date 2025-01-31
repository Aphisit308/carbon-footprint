package models


type CarbonFootprintCalculatorRequest struct {
	ActivityType string  `json:"activity_type"`
	DistanceKm   float64 `json:"distance_km"`
	VehicleType  string  `json:"vehicle_type"`
	FuelType     string  `json:"fuel_type"`
}

type CarbonFootprintCalculatorResponse struct {
	ActivityType     string  `json:"activity_type"`
	CarbonEmissionKg float64 `json:"carbon_emission_kg"`
}
