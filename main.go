package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "db-book-sql"
)

const PORT = ":8080"

type Book struct {
	BookID      uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null;type:varchar(191)" json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var (
	db  *gorm.DB
	err error
)

func init() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	route := gin.Default()

	bookRoutes := route.Group("/books")
	{
		bookRoutes.POST("/", createBook)
		bookRoutes.GET("/", getBooks)
		bookRoutes.GET("/:bookId", getBookById)
		bookRoutes.PUT("/:bookId", updateBook)
		bookRoutes.DELETE("/:bookId", deleteBook)
	}

	route.Run(PORT)
}

func getBooks(c *gin.Context) {
	books := []Book{}

	err = db.Find(&books).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	books := c.Param("bookId")

	parseId, err := strconv.Atoi(books)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid book id",
		})
		return
	}

	book := Book{
		BookID: uint(parseId),
	}

	err = db.First(&book).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "book not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) {
	bookId := c.Param("bookId")

	parseId, err := strconv.Atoi(bookId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid book id",
		})
		return
	}
	var bookRequest BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(422, gin.H{
			"message": "invalid request body",
		})
		return
	}

	book := Book{
		BookID: uint(parseId),
	}

	err = db.Model(&book).Updates(Book{Title: bookRequest.Title, Author: bookRequest.Author, Description: bookRequest.Desc}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
	bookId := c.Param("bookId")

	parseId, err := strconv.Atoi(bookId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid book id",
		})
		return
	}

	err = db.Delete(&Book{}, parseId).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book has been deleted successfully",
	})
}

func createBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(422, gin.H{
			"message": "invalid request body",
		})
		return
	}

	book := Book{
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Desc,
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(201, book)
}
