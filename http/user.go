package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	domainErrors "github.com/wackGarcia/books/domain/errors"
	"github.com/wackGarcia/books/domain/user"
	"github.com/wackGarcia/books/utils/cryptography"
)

type UserRequest struct {
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required,gte=8,lte=20" json:"password"`
}

// Create User godoc
// @Summary Create a new user.
// @Description This endpoint registries users to use the API..
// @Tags users
// @Accept application/json
// @Produce json
// @Param user body UserRequest{} true "create user struct"
// @Success 201
// @Failure 400,500
// @Router /user [post]
func (h Handler) UserCreateHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json UserRequest

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
