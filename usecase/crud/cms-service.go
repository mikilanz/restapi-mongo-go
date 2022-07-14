package service

import (
	//other import goes here
	"company-config-go/configs"
	"company-config-go/dataproviders/schema"
	"company-config-go/delivery/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson" //add this
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cmsConfigCollection *mongo.Collection = configs.GetCollection(configs.DB, "cmsconfigs")

func GetOne(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user schema.CmsConfigs
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := cmsConfigCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, dto.DefaultHttpResponse{StatusCode: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func GetAll(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []schema.CmsConfigs
		defer cancel()

		results, err := cmsConfigCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser schema.CmsConfigs
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			dto.DefaultHttpResponse{StatusCode: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}
