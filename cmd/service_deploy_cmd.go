package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serviceDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a service k2s operator",
	Long:  `deploy a service k2s operator`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service deploy command")
	},
}

func init() {
	serviceRootCmd.AddCommand(serviceDeployCmd)
}
