package main

import (
	"app/authentication"
	"app/settings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var connection string = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", settings.DBHOST, settings.DBPORT, settings.DBUSER, settings.DBNAME, settings.DBPASSWORD)
	db, err := gorm.Open("postgres", connection)
	authentication.DB = db
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/animal", authentication.ListAnimalView)
	r.GET("/animal/:id/", authentication.GetAnimalView)
	r.POST("/animal", authentication.CreateAnimalView)
	r.PUT("/animal/:id/", authentication.UpdateAnimalView)
	r.DELETE("/animal/:id/", authentication.DeleteAnimalView)

	r.Run()
}
