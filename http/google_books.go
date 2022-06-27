package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Find Book in Google godoc
// @Summary Find any book Google.
// @Description This endpoint allows find books using the `GOOGLE API`.
// @Tags Google books
// @Accept json
// @Produce json
// @Param   apikey     path    string     true        "AIzaSyBhW1SGJWkF4KSo3QOhLpFYA0c_IQSVQ14"
// @Param   query     query    string     true        "clean architecture"
// @Success 200 google structure
// @Failure 400,403,500
// @Router /google/book/:apikey [get]
func (h Handler) GoogleBookFindHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var response interface{}
		apiKey := ctx.Param("apikey")
		search := ctx.Query("q")
		u, err := url.Parse("https://www.googleapis.com/books/v1/volumes")
		if err != nil {
			ctx.Error(err)
			return
		}

		query := u.Query()
		query.Set("q", search)
		query.Set("key", apiKey)
		u.RawQuery = query.Encode()

		resp, err := http.Get(u.String())
		if err != nil {
			ctx.Error(err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ctx.Error(err)
			return
		}

		if err := json.Unmarshal(body, &response); err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(resp.StatusCode, response)
	}
}
