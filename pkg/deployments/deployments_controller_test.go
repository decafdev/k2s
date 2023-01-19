package deployments

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestDeploymentsController(t *testing.T) {
	tests := make(map[string]func(t *testing.T, client *resty.Client))

	tests["create deployment and list deployments"] = testControllerCreateList

	setupControllerIntegrationTest()

	// for situation, fn := range tests {
	// 	t.Run(situation, func(t *testing.T) {
			
	// 	})
	// }

	// type expected struct {
	// 	code int
	// 	body string
	// }

	// type given struct {
	// 	method   string
	// 	endpoint string
	// 	// body     string
	// 	expected expected
	// }

	// scenario := make(map[string]given)

	// scenario["when I want to list all k2s application deployments"] = given{
	// 	method:   http.MethodGet,
	// 	endpoint: "/deployments",
	// 	expected: expected{
	// 		code: http.StatusOK,
	// 		body: `[{}]`,
	// 	},
	// }

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

func setupControllerIntegrationTest() {

}

func testControllerCreateList(t *testing.T, client *resty.Client) {

}
