package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"crypto/sha256"
)

func main(){
	router := gin.Default()

	router.Static("/assets/","./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/",getIndex)
	router.POST("/",postIndex)

	router.GET("/content/:contentId",getContent)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func postIndex(c *gin.Context) {
	inputContent := c.PostForm("inputContent")
	c.HTML(http.StatusOK, "complete.html", gin.H{
		"inputContent" : inputContent,
		"hash" : sha256.Sum256([]byte(inputContent)),
		"publishTime" : 0,
		"password" : "password",
	})
}

func getContent(c *gin.Context) {
	//contentiD := c.Param("contentiD")
}
