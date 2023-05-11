package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Post struct {
	gorm.Model
	Title string
	Text  string
}

func hello(c *gin.Context) {
	c.String(200, "Hello World!")
}

func getAll(c *gin.Context) {

	var posts []Post
	DB.Find(&posts)
	c.JSON(200, posts)
}

func getById(c *gin.Context) {
	id := c.Param("id")

	var post Post

	DB.First(&post, id)

	c.JSON(200, post)
}

func create(c *gin.Context) {
	var body struct {
		Title string
		Text  string
	}

	c.Bind(&body)

	post := Post{Title: body.Title, Text: body.Text}

	result := DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, post)
}

func main() {
	DB, err = gorm.Open(postgres.Open("postgres://asoggexd:9tqUok-Xb7IyYHHzsry-c1HdLPYYgsMd@kandula.db.elephantsql.com/asoggexd"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//DB.AutoMigrate(&Post{})

	r := gin.Default()

	r.GET("/", hello)
	r.GET("/posts", getAll)
	r.GET("/posts/:id", getById)
	r.POST("/posts", create)

	r.Run()
}
