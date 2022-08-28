package main

import (
	"go-alqolam/helper"
	"go-alqolam/routes"
	"go-alqolam/user"
)

func main() {
	db := helper.SetupDB()
	// Migrate Table From Entity
	db.AutoMigrate(&user.User{})

	router := routes.SetupRoutes(db)
	router.Run()
}
