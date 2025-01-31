package config

import (
	"carbon-footprint/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var DB *gorm.DB

func SetupDB() {
	LoadConfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USERNAME"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DATABASE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connect database")
	}
	migrateDB()
}

func migrateDB() {
	DB.AutoMigrate(&models.EmissionFactor{})
	DB.AutoMigrate(&models.EmissionData{})

	vehicleType := "car"
	fuelType := "gasoline"
	DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&models.EmissionFactor{
		ID:                    1,
		EmissionFactor:        30,
		ActivityType:          "",
		EnergyConsumption:     20,
		EnergyConsumptionUnit: "km",
		VehicleType:           &vehicleType,
		FuelType:              &fuelType,
	})
}
