package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wackGarcia/books/domain/book"
	domainErrors "github.com/wackGarcia/books/domain/errors"
)

type BookCreateRequest struct {
	GoogleID   string `binding:"required" json:"google_id"`
	Author     string `binding:"required,gte=3,lte=50" json:"author"`
	Title      string `binding:"required,gte=3,lte=50" json:"title"`
	Publisher  string `binding:"required,gte=3,lte=50" json:"publisher"`
	WishlistID string `json:"whislist_id"`
}

// Create Book godoc
// @Summary Create book.
// @Description This endpoint allows create books.
// @Tags books
// @Accept json
// @Produce json
// @Param user body BookCreateRequest{} true "create book struct"
// @Success 201
// @Failure 400,403,500
// @Router /book [post]
func (h Handler) BookCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json BookCreateRequest
		userID, _ := ctx.Get("X-USER-ID")

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		user := &book.Book{
			GoogleID:  json.GoogleID,
			Author:    json.Author,
			Title:     json.Title,
			Publisher: json.Publisher,
			UserID:    fmt.Sprintf("%v", userID),
		}

		bookID, err := h.Book.Create(ctx, user)

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"id": bookID})
	}
}

func (h Handler) BookFindHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		userID, _ := ctx.Get("X-USER-ID")

		googleID := ctx.Query("google_id")
		author := ctx.Query("author")
		title := ctx.Query("title")
		publisher := ctx.Query("publisher")
		wishListID := ctx.Query("wish_list_id")

		books, err := h.Book.GetAll(ctx, &book.Book{
			GoogleID:   googleID,
			Author:     author,
			Title:      title,
			Publisher:  publisher,
			UserID:     fmt.Sprintf("%v", userID),
			WishlistID: wishListID,
		})
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, books)
	}
}

func (h Handler) BookDeleteHandaler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bookID := ctx.Param("bookID")
		userID, _ := ctx.Get("X-USER-ID")

		if len(bookID) > 0 && len(bookID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		err := h.Book.Delete(ctx, bookID, fmt.Sprintf("%v", userID))
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
