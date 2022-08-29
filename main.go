package main

import (
	"go-alqolam/helper"
	incomecategories "go-alqolam/income_categories"
	"go-alqolam/member"
	"go-alqolam/region"
	"go-alqolam/routes"
	"go-alqolam/user"
)

func main() {
	db := helper.SetupDB()
	// Migrate Table From Entity
	db.AutoMigrate(&user.User{}, &incomecategories.IncomeCategory{}, &region.Region{}, &member.Member{})

	router := routes.SetupRoutes(db)
	router.Run()
}
