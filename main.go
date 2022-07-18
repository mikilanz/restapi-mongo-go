package main

import (
	"company-config-go/applications/routes"
	"company-config-go/configs" //add this
	"company-config-go/docs"
	"company-config-go/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Company Config Engine
// @version         1.0
// @description     This is a sample result for Company Config Engine.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      companycms.nukeapp.dev

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	//run database
	// configs.ConnectDB()

	//routes

	router.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		docs.SwaggerInfo.Version = configs.GetEnvKey("APP_VERSION")
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerfiles.Handler)(context)
	})

	router.Use(middleware.AccessMiddleware())
	routes.CmsRoute(router)

	appPort := fmt.Sprintf("0.0.0.0:%v", configs.GetEnvKey("APP_PORT"))

	router.Run(appPort)
}
