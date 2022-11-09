package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_test/models"
)

type postUrl struct {
	Url string `form:"url"`
}

func main() {

	err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	app.GET("/url/get", func(c *gin.Context) {
		id := c.Query("id")
		intId, err_change := strconv.Atoi(id)
		data, err_get := models.GetUrl(intId)
		if id == "" || err_change != nil || err_get != nil {
			c.JSON(http.StatusOK, gin.H{
				"id":  -1,
				"url": "https://example.com",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"id":  data.Id,
			"url": data.Url,
		})
	})
	app.POST("/url/create", func(c *gin.Context) {
		var u postUrl
		if c.ShouldBind(&u) != nil {
			c.JSON(http.StatusOK, gin.H{
				"id":  -1,
				"url": "https://example.com",
			})
		}
		add(u.Url)
		url, _ := models.GetLastUrl()
		c.JSON(http.StatusOK, gin.H{
			"id":  url.Id,
			"url": url.Url,
		})

	})
	app.Run(":4000")
}

func add(url string) bool {
	var json models.Url
	json.Url = url
	success, _ := models.AddUrl(json)
	if success {
		return true
	} else {
		return false
	}
}
