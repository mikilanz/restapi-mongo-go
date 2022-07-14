package main

import (
	"company-config-go/applications/routes"
	"company-config-go/configs" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.CmsRoute(router)

	router.Run("localhost:4002")
}
