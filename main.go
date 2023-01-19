package main

import (
	databases "resto-backend-go-1/database"
	"resto-backend-go-1/routes"
)

func main() {
	databases.StartDB()
	routes.Routes()
}
