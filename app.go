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

func update(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Text  string
	}

	c.Bind(&body)

	var post Post

	DB.First(&post, id)

	post.Title = body.Title
	post.Text = body.Text

	DB.Save(&post)

	c.JSON(200, post)
}

func delete(c *gin.Context) {
	id := c.Param("id")

	var post Post

	DB.Delete(&post, id)

	c.Status(200)
}

func main() {
	DB, err = gorm.Open(postgres.Open("postgres://asoggexd:9tqUok-Xb7IyYHHzsry-c1HdLPYYgsMd@kandula.db.elephantsql.com/asoggexd"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//DB.AutoMigrate(&Post{})

	r := gin.Default()

	r.GET("/api", hello)
	r.GET("/api/posts", getAll)
	r.GET("/api/posts/:id", getById)
	r.POST("/api/posts", create)
	r.PUT("/api/posts/:id", update)
	r.DELETE("/api/posts/:id", delete)
	r.StaticFile("/", "./static/index.html")

	r.Run()
}
