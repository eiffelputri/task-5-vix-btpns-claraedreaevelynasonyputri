package main

import (
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/database"
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	r := gin.Default()

	router.InitializeRouter(r, db) // Initialize routes using the router configuration

	r.Run(":8080")
}
