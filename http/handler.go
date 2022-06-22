package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (handler Handler) Http() http.Handler {
	server := gin.Default()
	server.Use(HandlerError)
	group := server.Group(v1)
	{
		// AUTH ENPOINT
		group.POST("/auth", handler.AuthHandler())

		// USERS EDNPOINT
		group.POST("/user", handler.UserCreateHandler())

		// BOOKS ENPOINTS
		group.POST("/book", handler.BookCreateHandler())

		group.GET("/book/:bookID", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{"status": "ok"})
		})

		// WISHLIST ENDPOINTS
		group.POST("/wishlist", handler.WishListCreateHandler())
		group.GET("/wishlist/:wishlistID", handler.WishListFindHandler())
		group.GET("/wishlist", handler.WishListFindHandler())
	}

	return server
}
