package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// VERSION - This is converted to the git tag at compile time using tasks run build command
var VERSION = "0.0.0"
var SERVICE_NAME = "k2s-operator"

var cfgFile string

// rootCmd represents the base command when called without any sub commands
var rootCmd = &cobra.Command{
	Use:   "k2s",
	Short: "staggeringly simple and opinionated kubernetes deployments",
	Long:  `staggeringly simple and opinionated kubernetes deployments`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if ver, _ := cmd.Flags().GetBool("version"); ver {
			fmt.Println(VERSION)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.io.techdecaf.template.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Prints application version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var configFile = ""

	if cfgFile != "" {
		// Use config file from the flag.
		configFile = cfgFile
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configFile = filepath.Join(home, "k2s.yaml")
	}

	viper.SetConfigFile(configFile)

	// Search config in home directory for SetConfigName (without extension).
	// viper.AddConfigPath(home)
	// viper.SetConfigName("kli")

	// create config file if one does not already exist
	file, _ := os.OpenFile(configFile, os.O_RDONLY|os.O_CREATE, 0644)
	file.Close()

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("failed to read in config file")
	}
}
