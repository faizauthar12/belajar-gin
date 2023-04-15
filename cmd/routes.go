package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	route := gin.Default()

	bookRoutes := route.Group("/books")
	{
		bookRoutes.POST("/", app.books.Create)
		bookRoutes.GET("/", app.books.GetAll)
		bookRoutes.GET("/:bookId", app.books.GetById)
		bookRoutes.PUT("/:bookId", app.books.Update)
		bookRoutes.DELETE("/:bookId", app.books.Delete)
	}

	return route
}
