package models

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

type BooksModel struct {
	DB *gorm.DB
}

func (m *BooksModel) Create(c *gin.Context) {
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

	err := m.DB.Create(&book).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(201, book)
}

func (m *BooksModel) Delete(c *gin.Context) {
	bookId := c.Param("bookId")

	parseId, err := strconv.Atoi(bookId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid book id",
		})
		return
	}

	err = m.DB.Delete(&Book{}, parseId).Error

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

func (m *BooksModel) Update(c *gin.Context) {
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

	err = m.DB.Model(&book).Updates(Book{Title: bookRequest.Title, Author: bookRequest.Author, Description: bookRequest.Desc}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (m *BooksModel) GetAll(c *gin.Context) {
	books := []Book{}

	err := m.DB.Find(&books).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (m *BooksModel) GetById(c *gin.Context) {
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

	err = m.DB.First(&book).Error

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
