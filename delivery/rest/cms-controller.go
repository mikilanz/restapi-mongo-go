package controllers

import (
	//other import goes here

	service "company-config-go/usecase/crud"

	"github.com/gin-gonic/gin" //add this
)

// var validate = validator.New()

// CMSGetOne
// @Summary get one cms config
// @Schemes
// @Description get one cms config
// @Tags CMS Company
// @Accept json
// @Produce json
// @Param        companyCode   path      string  true  "Company Code"
// @Security BearerAuth
// @Router /cms/{companyCode} [get]
func GetOne() gin.HandlerFunc {
	var c *gin.Context
	return service.GetOne(c)
}

// CMSGetAll
// @Summary get all cms config
// @Schemes
// @Description get all cms config
// @Tags CMS Company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /cms [get]
func GetAll() gin.HandlerFunc {
	var c *gin.Context
	return service.GetAll(c)
}
