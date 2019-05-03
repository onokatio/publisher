package main

import (
	"time"
	"fmt"
	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PostData struct {
	InputContent string `form:"InputContent"`
	PublishMonth int    `form:"PublishMonth"`
	PublishDays int     `form:"PublishDays"`
	PublishHours int    `form:"PublishHours"`
	PublishMinutes int  `form:"PublishMinutes"`
	Password string     `form:"Password"`
}
type ContentList struct {
	gorm.Model
	Text string
	ContentHash string `gorm:"unique_index"`
	PublishTime time.Time
	PasswordHash string
	Uuid uuid.UUID
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

func PostData2ContentList(form PostData) ContentList {
	var content ContentList

	contentHash := sha256.Sum256([]byte(form.InputContent))

	PublishTime := time.Now().UTC().Add(time.Hour * time.Duration(form.PublishHours))
	PublishTime = PublishTime.Add(time.Minute * time.Duration(form.PublishMinutes))
	PublishTime = PublishTime.AddDate(0, form.PublishMonth, form.PublishDays)

	passwordHash := sha256.Sum256([]byte(form.Password))

	content.Text = form.InputContent
	content.ContentHash = hex.EncodeToString(contentHash[:])
	content.PublishTime = PublishTime
	content.PasswordHash = hex.EncodeToString(passwordHash[:])

	return content
}

func postIndex(c *gin.Context) {

	var form PostData

	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content := PostData2ContentList(form)

	id := uuid.New()
	fmt.Println(id)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	//db.Create(&content)

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err.Error())
	}

	c.HTML(http.StatusOK, "complete.html", gin.H{
		"Content": content.Text,
		"Hash": content.ContentHash,
		"PublishTime": content.PublishTime.In(location).String(),
	})
}

func getContent(c *gin.Context) {
	contentId := c.Param("contentId")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var content ContentList
	db.First(&content, "contentHash = ?", contentId)
}
