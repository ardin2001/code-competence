package main

import (
	route "echo_golang/routes"
	"echo_golang/utils"
)

func main() {
	utils.UserMigrate()
	e := route.Routers()
	e.Logger.Fatal(e.Start(":8000"))
}
