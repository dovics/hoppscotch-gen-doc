package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hoppscotch-gen-doc",
	Short: "A CLI tool to generate Markdown docs from Hoppscotch JSON collections",
	Long: `Hoppscotch to Markdown Generator

A CLI tool that converts Hoppscotch JSON collections into well-formatted
Markdown API documentation.

Features:
- Support for multi-level folder structures
- Parse Hoppscotch JSON collection files
- Generate Markdown documentation with:
  - Table of Contents
  - HTTP methods with visual badges
  - Request headers and parameters
  - Request bodies (formatted JSON)
  - Authentication information
  - Full description support`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
