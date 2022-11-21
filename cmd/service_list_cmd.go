package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/techdecaf/k2s/v2/pkg/sdk"
)

var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists services deployed using the k2s operator",
	Long:  `lists services deployed using the k2s operator`,
	Run: func(cmd *cobra.Command, args []string) {
		client := sdk.NewClient("http://localhost:3000")
		res, err := client.ListDeployments()
		if err != nil {
			log.Fatal(err)
		}
		sdk.PrettyPrintJSON(res)
	},
}

func init() {
	serviceRootCmd.AddCommand(serviceListCmd)
}
