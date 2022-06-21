package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	domainErrors "github.com/wackGarcia/books/domain/errors"
	"github.com/wackGarcia/books/utils/cryptography"
)

type AuthRequest struct {
	UserName string `binding:"required,userName" json:"username"`
	Password string `binding:"required,gte=8,lte=20" json:"password"`
}

func (h Handler) AuthHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json UserCreateRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		user, err := h.User.Get(ctx, json.UserName)
		if err != nil {
			ctx.Error(err)
			return
		}

		if err := cryptography.ComparePasswordWithSalt(json.Password, user.Password); err != nil {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotAuthorized)
			ctx.Error(appErr)
			return
		}

		token, err := cryptography.CreateJWTWithClaims(user.ID)
		if err != nil {
			log.Print(err)
			appErr := domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
			ctx.Error(appErr)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}
