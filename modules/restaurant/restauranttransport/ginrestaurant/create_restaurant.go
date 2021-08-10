package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-api-training/component"
	"go-api-training/modules/restaurant/restaurantbiz"
	"go-api-training/modules/restaurant/restaurantmodel"
	"go-api-training/modules/restaurant/restaurantstore"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.CreateRestaurantStore(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
