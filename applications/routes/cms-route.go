package routes

import (
	controllers "company-config-go/delivery/rest"

	"github.com/gin-gonic/gin"
)

func CmsRoute(router *gin.Engine) {
	//All routes related to cms comes here

	router.GET("/cms/:id", controllers.GetOne()) //add this
	router.GET("/cms", controllers.GetAll())
}
