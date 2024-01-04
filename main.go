package main

import (
	"github.com/joho/godotenv"
	"github.com/okyirmawan/employee-attendance-api/docs"
	routes2 "github.com/okyirmawan/employee-attendance-api/routes"
	"github.com/okyirmawan/employee-attendance-api/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			logrus.Fatal("Error loading .env file")
		}
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Employee Attendance API"
	docs.SwaggerInfo.Description = "This is a employee attendance API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	routes := routes2.Init()
	routes.Logger.Fatal(routes.Start(":" + utils.Getenv("PORT", "8080")))
}
