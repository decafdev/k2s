package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serviceRootCmd = &cobra.Command{
	Use:   "service",
	Short: "root command for service operations",
	Long:  `root command for service operations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service root command")
	},
}

func init() {
	rootCmd.AddCommand(serviceRootCmd)
}
