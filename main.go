package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_test/models"
)

func checkErr(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

func main() {

	err := models.ConnectDB()
	checkErr(err)

	app := gin.Default()
	ua := ""
	app.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
		})
	})
	app.Static("/img", "A:/contents_zip/el28whAtuu/")
	app.Run(":4000")
}

/*
func getUrl(id int) {
	url, err := models.GetUrl(id)
}*/
