package main

import (
	"net/http"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PostData struct {
	inputContent string `form:"inputContent" binding:"required"`
	publishMonth uint `form:"publishMonth" binding:"required"`
	publishDays uint `form:"publishDays" binding:"required"`
	publishHours uint `form:"publishHours" binding:"requirer"`
}
type DataList struct {
	gorm.Model
	Data string
	Time uint
	Hash string
	Password string
}

func main(){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.autoMigrate(&DataList{})
	db.create

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
