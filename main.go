package main

import (
	"go_crud/db"
	"go_crud/routes"
)

func main() {
	db.DBConfig()
	routes.InitRoutes()
}
