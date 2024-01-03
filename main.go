package main

import (
	"github.com/okyirmawan/employee-attendance-api/docs"
	routes2 "github.com/okyirmawan/employee-attendance-api/routes"
)

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a employee attendance API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	routes := routes2.Init()
	routes.Logger.Fatal(routes.Start(":8080"))
}
