package main

import (
	"log"
	"net/http"
	"os"
	"rest/component"
	"rest/middleware"
	"rest/modules/restaurant/restauranttransport/ginrestaurant"
	"rest/modules/user/usertransport/ginuser"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SYSTEM_SECRET")
	atExp, _ := strconv.Atoi(os.Getenv("AT_EXP"))
	rtExp, _ := strconv.Atoi(os.Getenv("RT_EXP"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	
	if err := runService(db, secretKey, atExp, rtExp); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string, atExp, rtExp int) error {
	appCtx := component.NewAppContext(db, secretKey, atExp, rtExp)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H {
			"message" : "pong",
		})
	})

	//versioning
	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))
	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx)) 
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
} 