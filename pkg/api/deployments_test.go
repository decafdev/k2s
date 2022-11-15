package api

import (
	"net/http"
	"testing"
)

func TestDeploymentsController(t *testing.T) {

	// expected struct
	type expected struct {
		code int
		body string
	}
	// given struct
	type given struct {
		method   string
		endpoint string
		body     string
		expected expected
	}

	scenario := make(map[string]given)

	scenario["when I want to list all k2s application deployments"] = given{
		method:   http.MethodGet,
		endpoint: "/deployments",
		expected: expected{
			code: http.StatusOK,
			body: `[{}]`,
		},
	}

	// for when, given := range scenario {
	// 	t.Run(when, func(t *testing.T) {
	// 		res := httptest.NewRecorder()
	// 		context, app := gin.CreateTestContext(res)
	// 		context.Request, _ = http.NewRequest(given.method, given.endpoint, bytes.NewBuffer([]byte(given.body)))
	// 		deployments.NewDeploymentController(app, config)

	// 		// act
	// 		app.ServeHTTP(res, context.Request)

	// 		// assert
	// 		assert.JSONEq(t, given.expected.body, string(res.Body.Bytes()))
	// 		assert.Equal(t, given.expected.code, res.Code)
	// 	})
	// }
}
