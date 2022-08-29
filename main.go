package main

import (
	"go-alqolam/helper"
	incomecategories "go-alqolam/income_categories"
	"go-alqolam/routes"
	"go-alqolam/user"
)

func main() {
	db := helper.SetupDB()
	// Migrate Table From Entity
	db.AutoMigrate(&user.User{}, &incomecategories.IncomeCategory{})

	router := routes.SetupRoutes(db)
	router.Run()
}
