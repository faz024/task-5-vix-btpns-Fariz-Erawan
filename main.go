package main

import (
	"os"
	"task-5-vix-btpns-fariz/database"
	"task-5-vix-btpns-fariz/models"
	"task-5-vix-btpns-fariz/router"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{})

	r := router.InitRoutes(db)
	r.Run(":" + os.Getenv("PORT"))
}
