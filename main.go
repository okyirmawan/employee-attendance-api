package main

import (
	"github.com/okyirmawan/employee-attendance-api/docs"
	routes2 "github.com/okyirmawan/employee-attendance-api/routes"
	"os"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a employee attendance API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + getPort()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	os.Setenv("TZ", "Asia/Jakarta")

	routes := routes2.Init()
	routes.Logger.Fatal(routes.Start(":" + getPort()))
}

func getPort() string {
	// Get the PORT environment variable or use a default value (e.g., "8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}
