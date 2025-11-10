package main

import (
	"log"

	deepspec "github.com/commercetools/deepspec/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deepspec",
	Short: "DeepSpec - Better spec driven development",
	Run: func(cmd *cobra.Command, args []string) {
		// Default: start the TUI
		if err := deepspec.StartTUI(); err != nil {
			log.Fatalf("TUI error: %v", err)
		}
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the MCP server",
	Run: func(cmd *cobra.Command, args []string) {
		server := deepspec.NewServer()
		if err := server.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
