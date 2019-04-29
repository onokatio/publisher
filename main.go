package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()

	router.GET("/",index)
	router.Static("/assets/","./assets")
	router.GET("/ping/:id",ping)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func index(c *gin.Context) {
}

func ping(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "message: pong" + id)
}
