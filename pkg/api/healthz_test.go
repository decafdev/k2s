package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/util"
)

func TestHealthzController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	util.NewConfig("VERSION=99.99.99").Validate()
	type expected struct {
		response string
		code     int
	}
	// given struct
	type given struct {
		endpoint string
		body     string
		expected expected
	}

	tests := make(map[string]given)

	tests["GET /healthz"] = given{
		endpoint: "/healthz",
		body:     "{}",
		expected: expected{
			response: `{
        "hostname": "",
        "name": "k2s-operator",
        "version": "99.99.99"
      }`,
			code: 200,
		},
	}

	// for when, given := range tests {
	// 	t.Run(when, func(t *testing.T) {
	// 		res := httptest.NewRecorder()
	// 		context, app := gin.CreateTestContext(res)
	// 		context.Request, _ = http.NewRequest(http.MethodGet, given.endpoint, bytes.NewBuffer([]byte(given.body)))
	// 		healthz.NewHealthzController(app, config)
	// 		// act
	// 		app.ServeHTTP(res, context.Request)

	// 		// assert
	// 		snaps.MatchSnapshot(t, string(res.Body.Bytes()))

	// 	})
	// }
}
