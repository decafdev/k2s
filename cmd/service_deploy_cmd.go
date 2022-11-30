package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
	"github.com/techdecaf/k2s/v2/pkg/sdk"
)

var serviceDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a service k2s operator",
	Long:  `deploy a service k2s operator`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetString("version")
		port, _ := cmd.Flags().GetString("port")
		image, _ := cmd.Flags().GetString("image")

		deployment := deployments.DeploymentDTO{
			Name:    name,
			Version: version,
			Port:    port,
			Image:   image,
		}

		res, err := sdk.NewClient("http://localhost:3000").CreateDeployment(deployment)
		if err != nil {
			log.Fatal(err)
		}

		sdk.PrettyPrintJSON(res)
	},
}

func init() {
	serviceRootCmd.AddCommand(serviceDeployCmd)
	// flags
	serviceDeployCmd.Flags().String("name", "", "the name of the service you wish to deploy")
	serviceDeployCmd.MarkFlagRequired("name")

	serviceDeployCmd.Flags().String("version", "", "a semantic version example: 1.0.1")
	serviceDeployCmd.MarkFlagRequired("version")

	serviceDeployCmd.Flags().String("port", "80", "the port number your service will listen on example: 80")

	serviceDeployCmd.Flags().String("image", "", "a docker image to use for this deployment example: traefik/whoami:v1.8.7")
	serviceDeployCmd.MarkFlagRequired("image")
}
