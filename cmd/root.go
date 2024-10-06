package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for the CLI
var rootCmd = &cobra.Command{
	Use:   "blogcli",
	Short: "Blog CLI to manage posts, tags, and categories",
	Long:  "A command-line tool for managing blog posts, tags, and categories globally",
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
