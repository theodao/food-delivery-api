package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column: name;"`
	Addr string `json:"address" gorm:"column: addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
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
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", func(c *gin.Context) {
			var data Restaurant
			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			if err := db.Create(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
			}

			var data Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, data)
		})
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
