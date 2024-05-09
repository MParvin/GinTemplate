package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-template",
	Short: "A Golang Gin template project",
	Long:  `A Golang Gin template project with a simple REST API, user authentication, and Swagger documentation.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-template.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
