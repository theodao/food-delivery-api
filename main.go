package main

import (
	"github.com/gin-gonic/gin"
	"go-api-training/component"
	"go-api-training/modules/restaurant/restauranttransport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messsage": "pong",
		})
	})

	// CRUD
	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		//restaurants.GET("/:id", func(c *gin.Context) {
		//	id, err := strconv.Atoi(c.Param("id"))
		//
		//	if err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//	}
		//
		//	var data Restaurant
		//
		//	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, data)
		//})

		//restaurants.GET("", func(c *gin.Context) {
		//	var data []Restaurant
		//
		//	type Filter struct {
		//		CityId int `json:"city_id" form:"city_id"`
		//	}
		//
		//	var filter Filter
		//	c.ShouldBind(&filter)
		//	newDb := db
		//
		//	if filter.CityId > 0 {
		//		newDb = db.Where("city_id = ?", filter.CityId)
		//	}
		//
		//	if err := newDb.Find(&data).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, data)
		//})

		//restaurants.PUT("/:id", func(c *gin.Context) {
		//	id, err := strconv.Atoi(c.Param("id"))
		//
		//	if err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//	}
		//
		//	var data RestaurantUpdate
		//	if err := c.ShouldBind(&data); err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, gin.H{"ok": 1})
		//})
	}

	return r.Run()
}

func main() {
	dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to run ")
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}
