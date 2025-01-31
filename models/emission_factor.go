package models

type EmissionFactor struct {
	ID                    uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ActivityType          string  `grom:"NOT NULL" json:"activity_type"`
	EmissionFactor        float32 `gorm:"NOT NULL" json:"factor_value"`
	EnergyConsumption     float32 `gorm:"NOT NULL" json:"energy_consumption"`
	EnergyConsumptionUnit string  `gorm:"NOT NULL" json:"energy_consumption_unit"`
	VehicleType           *string `gorm:"DEFAULT:NULL" json:"vehicle_type"`
	FuelType              *string `gorm:"DEFAULT:NULL" json:"fuel_type"`
}
