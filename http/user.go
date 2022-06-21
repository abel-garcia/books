package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	domainErrors "github.com/wackGarcia/books/domain/errors"
	"github.com/wackGarcia/books/domain/user"
	"github.com/wackGarcia/books/utils/cryptography"
)

type UserCreateRequest struct {
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required,gte=8,lte=50" json:"password"`
}

func (h Handler) UserCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json UserCreateRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			appError := domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
			appError.Err = errors.Wrap(appError.Err, err.Error())
			ctx.Error(appError)
			return
		}

		password, err := cryptography.CreatePasswordWithSalt(json.Password)
		if err != nil {
			ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.UnknownError))
		}

		user := &user.User{
			UserName: json.UserName,
			Password: password,
		}

		if err := h.User.Create(ctx, user); err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
