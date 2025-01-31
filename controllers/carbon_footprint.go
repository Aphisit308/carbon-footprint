package controllers

import (
	"carbon-footprint/config"
	"carbon-footprint/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/carbon/footprint/calculate", Calculator)

}
func Calculator(c *fiber.Ctx) error {
	req := new(models.CarbonFootprintCalculatorRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	var response models.CarbonFootprintCalculatorResponse
	var emissionFactor models.EmissionFactor
	if err := config.DB.
		Table("emission_factors").
		Where(`
			emission_factors.activity_type = $1 AND
			emission_factors.vehicle_type = $2 AND 
			emission_factors.fuel_type = $3
		`, req.ActivityType, req.VehicleType, req.FuelType).
		First(&emissionFactor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	response.ActivityType = req.ActivityType
	response.CarbonEmissionKg = float64(emissionFactor.EmissionFactor) * req.DistanceKm

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "calculate success",
		"data":    response,
	})
}
