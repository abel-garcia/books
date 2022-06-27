package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	domainErrors "github.com/wackGarcia/books/domain/errors"
	"github.com/wackGarcia/books/domain/wishlist"
)

type WishListCreateRequest struct {
	Name string `binding:"required,gte=3,lte=50" json:"name"`
}

type WishListResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Create Book WishList godoc
// @Summary Create a new book wishlist.
// @Description This endpoint allows create a new book wishlist.
// @Tags wishlist
// @Accept application/json
// @Produce json
// @Param wishlist body WishListCreateRequest{} true "create whislist struct"
// @Success 200 {object} WishListResponse{}
// @Failure 400,403,409,500
// @Router /wishlist [post]
func (h Handler) WishListCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json WishListCreateRequest
		userID, _ := ctx.Get("X-USER-ID")

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		id, err := h.WishList.Create(ctx, &wishlist.WishList{
			Name:   json.Name,
			UserID: fmt.Sprintf("%v", userID),
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	}
}

// Find Book WishList godoc
// @Summary Find book wishlist.
// @Description This endpoint allows find a specific `Book WishList` by ID or get all `Book WishList` by user.
// @Tags wishlist
// @Accept json
// @Produce json
// @Param   ID     path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Success 200 {object} []WishListResponse{}
// @Failure 400,403,500
// @Router /wishlist [get]
func (h Handler) WishListFindHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		wishlistID := ctx.Param("ID")
		userID, _ := ctx.Get("X-USER-ID")

		if len(wishlistID) > 0 && len(wishlistID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		list, err := h.WishList.Get(ctx, &wishlist.WishList{
			ID:     wishlistID,
			UserID: fmt.Sprintf("%v", userID),
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		response := []WishListResponse{}
		for _, v := range list {
			response = append(response, WishListResponse{
				ID:        v.ID,
				Name:      v.Name,
				UserID:    v.UserID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.CreatedAt,
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}

// Delete Book WishList godoc
// @Summary Delete book wishlist.
// @Description This endpoint allows delet a specific `Book WishList` this remove all books in relation to.
// @Tags wishlist
// @Accept json
// @Produce json
// @Param   wishlistID    path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Success 200
// @Failure 400,403,500
// @Router /wishlist [delete]
func (h Handler) WishListDeleteHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		wishlistID := ctx.Param("wishlistID")
		userID, _ := ctx.Get("X-USER-ID")

		if len(wishlistID) > 0 && len(wishlistID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		id, err := h.WishList.Delete(ctx, wishlistID, fmt.Sprintf("%v", userID))
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}

// Add Book To WishList godoc
// @Summary Add a new book to wishlist.
// @Description This endpoint allows to add books to wishlist.
// @Tags wishlist book
// @Produce json
// @Param   wishlistID     path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Param   bookID     path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Failure 400,403,409,500
// @Router  /wishlist/:wishlistID/book/:bookID [post]
func (h Handler) WishListAddBookHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		wishlistID := ctx.Param("wishlistID")
		bookID := ctx.Param("bookID")
		userID, _ := ctx.Get("X-USER-ID")

		if len(wishlistID) > 0 && len(wishlistID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		if len(bookID) > 0 && len(bookID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		if err := h.WishList.AddBook(ctx, wishlistID, bookID, fmt.Sprintf("%v", userID)); err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

// Delete Book From WishList godoc
// @Summary Delete an specific book from wishlist.
// @Description This endpoint allows delete a specific book from an wishlist.
// @Tags wishlist book
// @Accept json
// @Produce json
// @Param   wishlistID     path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Param   bookID     path    string     false        "49c88b9c-e324-42cf-802f-897f5b16cd37"
// @Success 200
// @Failure 400,403,500
// @Router /wishlist/:wishlistID/book/:bookID [delete]
func (h Handler) WishListDeleteBookHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		wishlistID := ctx.Param("wishlistID")
		bookID := ctx.Param("bookID")
		userID, _ := ctx.Get("X-USER-ID")

		if len(wishlistID) > 0 && len(wishlistID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		if len(bookID) > 0 && len(bookID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		if err := h.WishList.DeleteBook(ctx, wishlistID, bookID, fmt.Sprintf("%v", userID)); err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
