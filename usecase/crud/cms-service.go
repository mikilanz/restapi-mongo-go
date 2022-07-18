package service

import (
	//other import goes here
	"company-config-go/configs"
	"company-config-go/dataproviders/schema"
	"company-config-go/delivery/dto"
	"company-config-go/usecase/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson" //add this
	"go.mongodb.org/mongo-driver/mongo"
)

var cmsConfigCollection *mongo.Collection = configs.GetCollection(configs.DB, "cmsconfigs")

func GetOne(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		companyCode := c.Param("companyCode")

		var user schema.CmsConfigs
		defer cancel()

		// objId, _ := primitive.ObjectIDFromHex(companyCode)

		err := cmsConfigCollection.FindOne(ctx, bson.M{"companyCode": companyCode}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func GetAll(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var result []schema.CmsConfigs
		defer cancel()

		size := c.Query("size")
		if size == "" {
			size = "10"
		}

		page := c.Query("page")
		if page == "" {
			page = "0"
		}

		sort := c.Query("sort")
		if sort == "" {
			sort = "id,DESC"
		}

		sizeInt, _ := strconv.Atoi(size)
		pageInt, _ := strconv.Atoi(page)
		sortField := strings.Split(sort, ",")[0]
		sortType := strings.Split(sort, ",")[1]

		lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "siteconfigs"},
			{Key: "localField", Value: "companyCode"},
			{Key: "foreignField", Value: "companyCode"},
			// {Key: "let", Value: bson.D{{Key: "companyCode", Value: "$companyCode"}}},
			// {Key: "pipeline", Value: bson.A{bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$companyCode", "$$companyCode"}}}}}}}}}}}},
			{Key: "as", Value: "config"},
		}}}

		// lookupStage := bson.M{"$lookup": bson.M{ "from": "siteconfigs", "localField": "companyCode", "foreignField": "companyCode", "as": "config"}}
		unwindStage := bson.D{{Key: "$unwind", Value: "$config"}}

		sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: sortField, Value: utils.ConvertSortToMongo(strings.ToLower(sortType))}}}}

		limitStage := bson.D{{Key: "$limit", Value: sizeInt + pageInt}}
		skipStage := bson.D{{Key: "$skip", Value: pageInt}}

		// cursor, err := cmsConfigCollection.Find(ctx, bson.M{})
		cursor, err := cmsConfigCollection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, sortStage, limitStage, skipStage})

		cursor.All(ctx, &result)

		counts, _ := cmsConfigCollection.CountDocuments(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}})
			return
		}

		// for cursor.Next(ctx) {
		// 	var singleResult schema.CmsConfigs
		// 	fmt.Println(singleResult)
		// 	if err = cursor.Decode(&singleResult); err != nil {
		// 		c.JSON(http.StatusInternalServerError, dto.DefaultHttpResponse{StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}})
		// 		return
		// 	}

		// 	result = append(result, singleResult)
		// }

		//reading from the db in an optimal way
		defer cursor.Close(ctx)

		c.Header("x-total-count", fmt.Sprint(counts))
		c.JSON(http.StatusOK, result)
	}
}
