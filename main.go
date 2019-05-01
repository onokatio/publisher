package main

import (
	"time"
	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PostData struct {
	inputContent string `form:"inputContent" binding:"required"`
	publishMonth int `form:"publishMonth" binding:"required"`
	publishDays int `form:"publishDays" binding:"required"`
	publishHours int `form:"publishHours" binding:"requirer"`
	publishMinutes int `form:"publishMinutes" binding:"requirer"`
	password string `form:"password" binding:"requirer"`
}
type ContentList struct {
	gorm.Model
	content string
	contentHash string
	publishTime int64
	passwordHash string
}

func main(){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&ContentList{})

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
	var form PostData
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data ContentList

	contentHash := sha256.Sum256([]byte(form.inputContent))

	publishTime := time.Now().Add(time.Hour * time.Duration(form.publishHours))
	publishTime = publishTime.Add(time.Minute * time.Duration(form.publishMinutes))
	publishTime = publishTime.AddDate(0, form.publishMonth, form.publishDays)

	passwordHash := sha256.Sum256([]byte(form.password))

	data.content = form.inputContent
	data.publishTime = publishTime.Unix()
	data.contentHash = hex.EncodeToString(contentHash[:])
	data.passwordHash = hex.EncodeToString(passwordHash[:])
}

func getContent(c *gin.Context) {
	//contentiD := c.Param("contentiD")
}
