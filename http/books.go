package http

import (
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

func (h Handler) BookCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json BookCreateRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		user := &book.Book{
			GoogleID:   json.GoogleID,
			Author:     json.Author,
			Title:      json.Title,
			Publisher:  json.Publisher,
			UserID:     "cf145774-28ee-4f48-b4fd-ae042955bc34",
			WishlistID: json.WishlistID,
		}

		if err := h.Book.Create(ctx, user); err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

func (h Handler) BookFindHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json BookCreateRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		user := &book.Book{
			GoogleID:   json.GoogleID,
			Author:     json.Author,
			Title:      json.Title,
			Publisher:  json.Publisher,
			UserID:     "cf145774-28ee-4f48-b4fd-ae042955bc34",
			WishlistID: json.WishlistID,
		}

		if err := h.Book.Create(ctx, user); err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
