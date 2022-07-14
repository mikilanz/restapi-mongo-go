package controllers

import (
	//other import goes here

	service "company-config-go/usecase/crud"

	"github.com/gin-gonic/gin" //add this
)

// var validate = validator.New()

func GetOne() gin.HandlerFunc {
	var c *gin.Context
	return service.GetOne(c)
}

func GetAll() gin.HandlerFunc {
	var c *gin.Context
	return service.GetAll(c)
}
