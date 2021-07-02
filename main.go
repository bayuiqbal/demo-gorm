package main

import (
	"github.com/bayuiqballl/demo/models"
	"github.com/bayuiqballl/demo/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.People{})

	r := routes.SetupRoutes(db)
	r.Run()
}
