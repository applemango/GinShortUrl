package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_test/models"
)

type postUrl struct {
	Url string `form:"url"`
}

func checkErr(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

func main() {

	err := models.ConnectDB()
	checkErr(err)
	/*add()
	d, _ := models.GetUrl(1)
	fmt.Println(d.Url)*/

	d, err := models.GetLastUrl()
	fmt.Fprintf(os.Stderr, "err:%v\n", err)
	fmt.Printf("data: %s\n", d.Url)

	app := gin.Default()
	/*ua := ""
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
	app.Static("/img", "A:/contents_zip/el28whAtuu/")*/
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

		/*if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"id":  -1,
				"url": "https://example.com",
			})
		}*/
		//var out string
		//gob.NewDecoder(c.Request.Body).Decode(&out)
		//fmt.Println(out)
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
	//success, err := models.AddUrl(json)
	success, _ := models.AddUrl(json)
	if success {
		return true
	} else {
		return false
		//fmt.Fprintf(os.Stderr, "err :%v\n", err)
	}
}

/*
func getUrl(id int) {
	url, err := models.GetUrl(id)
}*/
