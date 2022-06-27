package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wackGarcia/books/domain/book"
	"github.com/wackGarcia/books/domain/user"
	"github.com/wackGarcia/books/domain/wishlist"
)

const (
	v1 = "v1"
)

// Set up services from domain ->
type Handler struct {
	User     *user.Service
	Book     *book.Service
	WishList *wishlist.Service
}

// @title Challenge Books API
// @version 1.0
// @description This API expose the methods to manage the book lits. A MELI challenge.

// @contact.name API Support
// @contact.email abelgarcia38348@gmail.com

// @license.name MIT License
// @license.url https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt

// @BasePath /v1
// @schemes http https
func (handler Handler) Http() http.Handler {
	server := gin.Default()
	server.Use(HandlerError)

	// swagger middleware documentation.
	server.Static("/docs", "./docs")
	url := ginSwagger.URL("http://localhost:8989/docs/swagger.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	group := server.Group(v1)

	group.POST("/auth", handler.AuthHandler())
	group.POST("/user", handler.UserCreateHandler())
	group.GET("/health", handler.HealthCheck())
	group.Use(Authorization)
	{
		// BOOKS
		group.POST("/book", handler.BookCreateHandler())
		group.GET("/book", handler.BookFindHandler())
		group.DELETE("/book/:bookID", handler.BookDeleteHandaler())

		// GOOGLE BOOKS
		group.GET("/google/book/:apikey", handler.GoogleBookFindHandler())

		// WISHLIST
		group.POST("/wishlist", handler.WishListCreateHandler())
		group.GET("/wishlist/:ID", handler.WishListFindHandler())
		group.GET("/wishlist", handler.WishListFindHandler())
		group.DELETE("/wishlist/:wishlistID", handler.WishListDeleteHandler())

		// ADD AND DELETE BOOKS FROM WISHLIST
		group.DELETE("/wishlist/:wishlistID/book/:bookID", handler.WishListDeleteBookHandler())
		group.PUT("/wishlist/:wishlistID/book/:bookID", handler.WishListAddBookHandler())
	}

	return server
}
