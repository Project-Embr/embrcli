package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "embrcli",
	Short: "Cli application to interface with Embr",
	Long: `Embrcli is a CLI application for interacting with Project Embr's API.
This is a tool to speed up and automate developer workflows`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.embrcli.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
