package deployments

import "testing"

func TestDeploymentService(t *testing.T) {
	tests := make(map[string]func(t *testing.T, service *DeploymentService))

	tests["create deployment and list deployments"] = testServiceCreateList

	setupServiceIntegrationTest()

	// for situation, fn := range tests {
	// 	t.Run(situation, func(t *testing.T) {
	// 		fn(t, NewDeploymentService())
	// 	})
	// }
}

func setupServiceIntegrationTest() {

}

func testServiceCreateList(t *testing.T, service *DeploymentService) {

}
