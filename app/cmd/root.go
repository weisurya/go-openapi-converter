package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "Convert OpenAPI spec into other formats",
		Long:  "go-openapi-converter is a golang-based tool to convert OpenAPI spec into other formats. Currently, it only supports .docx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, buddy! Thank you for using go-openapi-converter.")
		},
		Example: "./go-openapi-converter",
	}
)

// Execute to run this command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
