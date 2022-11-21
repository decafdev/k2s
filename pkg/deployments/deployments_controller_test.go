package deployments_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
	mock_depl_srv "github.com/techdecaf/k2s/v2/pkg/deployments/mock"
	"github.com/techdecaf/k2s/v2/pkg/logger"
	"github.com/techdecaf/k2s/v2/pkg/state"
)

var VERSION = "0.0.0"
var SERVICE_NAME = "k2s-operator"

func TestDeploymentsController(t *testing.T) {
	type given struct {
		method        string
		url           string
		body          []byte
		expectations  func(srv *mock_depl_srv.MockDeploymentSrv)
		checkResponse func(t *testing.T, res *httptest.ResponseRecorder)
	}

	testPost := make(map[string]given)

	depl := &state.DeploymentDTO{Name: "whoami", Image: "traefik/whoami", Version: "1.0.0"}
	values := map[string]string{"name": "whoami", "image": "traefik/whoami", "version": "1.0.0"}
	json_data, err := json.Marshal(values)
	require.NoError(t, err)

	testPost["when I make a proper POST request to create a deployment"] = given{
		method: string(http.MethodPost),
		url:    "/deployments",
		body:   json_data,
		expectations: func(srv *mock_depl_srv.MockDeploymentSrv) {
			srv.EXPECT().CreateDeployment(gomock.Eq(depl)).Times(1).Return(nil)
		},
		checkResponse: func(t *testing.T, res *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, res.Code)
			requireBodyMatch(t, nil, res.Body)
		},
	}

	// TODO: add more given cases

	for when, given := range testPost {
		t.Run(when, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// setup mock depl srv and expectations
			srv := mock_depl_srv.NewMockDeploymentSrv(ctrl)
			given.expectations(srv)

			// TODO: abstract/mock over next few lines
			os.Setenv("SERVICE_NAME", SERVICE_NAME)
			os.Setenv("VERSION", VERSION)

			configService, err := config.NewConfigService(os.Environ()...).Validate()
			require.NoError(t, err)
			log := logger.NewLogger(configService)
			l := log.WithFields(logrus.Fields{"module": "deployments"})
			//

			// setup test server and controller
			res := httptest.NewRecorder()
			context, app := gin.CreateTestContext(res)
			deployments.NewDeploymentController(app, srv, l)

			// setup request
			context.Request, err = http.NewRequest(given.method, given.url, bytes.NewBuffer(given.body))
			context.Request.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)

			// send request
			app.ServeHTTP(res, context.Request)
			given.checkResponse(t, res)
		})
	}
}

func requireBodyMatch(t *testing.T, depl *state.DeploymentDTO, body *bytes.Buffer) {
	bytes, err := io.ReadAll(body)
	require.NoError(t, err)

	var actual *state.DeploymentDTO
	err = json.Unmarshal(bytes, &actual)
	require.NoError(t, err)
	require.Equal(t, depl, actual)
}

// func TestDeploymentsController(t *testing.T) {

// 	// expected struct
// 	type expected struct {
// 		code int
// 		body string
// 	}
// 	// given struct
// 	type given struct {
// 		method   string
// 		endpoint string
// 		body     string
// 		expected expected
// 	}

// 	scenario := make(map[string]given)

// 	scenario["when I want to list all k2s application deployments"] = given{
// 		method:   http.MethodGet,
// 		endpoint: "/deployments",
// 		expected: expected{
// 			code: http.StatusOK,
// 			body: `[{}]`,
// 		},
// 	}

// 	// for when, given := range scenario {
// 	// 	t.Run(when, func(t *testing.T) {
// 	// 		res := httptest.NewRecorder()
// 	// 		context, app := gin.CreateTestContext(res)
// 	// 		context.Request, _ = http.NewRequest(given.method, given.endpoint, bytes.NewBuffer([]byte(given.body)))
// 	// 		deployments.NewDeploymentController(app, config)

// 	// 		// act
// 	// 		app.ServeHTTP(res, context.Request)

// 	// 		// assert
// 	// 		assert.JSONEq(t, given.expected.body, string(res.Body.Bytes()))
// 	// 		assert.Equal(t, given.expected.code, res.Code)
// 	// 	})
// 	// }
// }
