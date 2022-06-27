package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mockRepository "github.com/wackGarcia/books/data/mocks/user"
	"github.com/wackGarcia/books/domain/user"
)

func TestCreateHandler(t *testing.T) {

	testCases := map[string]struct {
		statusCode int
		user       UserRequest
	}{
		"success": {
			statusCode: http.StatusOK,
			user: UserRequest{
				UserName: "abelgarcia",
				Password: "12345678",
			},
		},
		// Fail password empty.
		"badRequest": {
			statusCode: http.StatusOK,
			user: UserRequest{
				UserName: "abelgarcia",
			},
		},
	}

	for testName, testCase := range testCases {
		testName, testCase := testName, testCase
		t.Run(testName, func(t *testing.T) {
			userRepository := mockRepository.NewUserMockRepository()
			userService := user.NewUserService(userRepository)

			// Prepare body request.
			requestBody, err := json.Marshal(map[string]interface{}{
				"username": testCase.user.UserName,
				"password": testCase.user.Password,
			})
			if err != nil {
				t.Fatalf("Error marshaling body request %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatalf("error creating http request %v", err)
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = req
			h := &Handler{User: userService}
			h.UserCreateHandler()(ctx)

			if status := w.Code; status != testCase.statusCode {
				t.Errorf("unexpected HTTP response code, expected: %d, response: %d", testCase.statusCode, w.Code)
			}
		})
	}
}
