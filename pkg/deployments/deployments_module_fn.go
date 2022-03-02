package deployments

import "github.com/techdecaf/k2s/v2/pkg/global"

func Module(app *global.Dependencies) (err error) {
	deploymentsTable, err := app.Streams.Table("deployments")
	if err != nil {
		return err
	}

	NewDeploymentController(app.Gin, NewDeploymentService(deploymentsTable))
	return
}
