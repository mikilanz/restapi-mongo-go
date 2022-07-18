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
// @Param        x-apigateway-api-userinfo   header      string  true  "x-apigateway-api-userinfo"
// @Param        companyCode   path      string  true  "Company Code"
// @Success      200  {object}   schema.CmsConfigs
// @Failure      400  {object}  dto.DefaultHttpResponse
// @Failure      404  {object}  dto.DefaultHttpResponse
// @Failure      500  {object}  dto.DefaultHttpResponse
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
// @Param        x-apigateway-api-userinfo   header      string  true  "x-apigateway-api-userinfo"
// @Param   size     query     int     false  "size"
// @Param   page     query     int     false  "page"
// @Param   sort     query     string     false  "sort"     example("id,DESC")
// @Success      200  {array}   schema.CmsConfigs
// @Failure      400  {object}  dto.DefaultHttpResponse
// @Failure      404  {object}  dto.DefaultHttpResponse
// @Failure      500  {object}  dto.DefaultHttpResponse
// @Security BearerAuth
// @Router /cms [get]
func GetAll() gin.HandlerFunc {
	var c *gin.Context
	return service.GetAll(c)
}
