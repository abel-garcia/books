package http

import (
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

func (h Handler) WishListCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json WishListCreateRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		id, err := h.WishList.Create(ctx, &wishlist.WishList{
			Name:   json.Name,
			UserID: "cf145774-28ee-4f48-b4fd-ae042955bc34",
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

func (h Handler) WishListFindHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		wishlistID := ctx.Param("wishlistID")

		if len(wishlistID) > 0 && len(wishlistID) != 36 {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			ctx.Error(appError)
			return
		}

		list, err := h.WishList.Get(ctx, &wishlist.WishList{
			ID:     wishlistID,
			UserID: "cf145774-28ee-4f48-b4fd-ae042955bc34",
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
