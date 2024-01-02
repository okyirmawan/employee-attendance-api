package main

import routes2 "github.com/okyirmawan/employee-attendance-api/routes"

func main() {
	routes := routes2.Init()

	routes.Logger.Fatal(routes.Start(":8888"))
}
