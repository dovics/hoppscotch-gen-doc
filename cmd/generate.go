package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/dovics/hoppscotch-gen-doc/internal/generator"
)

var (
	inputFile       string
	outputFile      string
	executeGET      bool
	timeout         int
	serverURL       string
	targetServerURL string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Markdown documentation from Hoppscotch JSON",
	Long:  `Generate Markdown documentation from a Hoppscotch JSON collection file.`,
	Args:  cobra.NoArgs,
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input Hoppscotch JSON file (required)")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output Markdown file (optional, prints to stdout if not specified)")
	generateCmd.Flags().BoolVarP(&executeGET, "execute", "x", false, "Execute GET requests and include responses in documentation")
	generateCmd.Flags().IntVarP(&timeout, "timeout", "t", 10, "Request timeout in seconds")
	generateCmd.Flags().StringVar(&serverURL, "server", "", "Replace endpoint host in documentation (e.g., https://api.example.com)")
	generateCmd.Flags().StringVar(&targetServerURL, "target-server", "", "Replace endpoint host only when executing requests (original URL shown in documentation)")
	generateCmd.MarkFlagRequired("input")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	// Create options
	opts := &generator.Options{
		ExecuteGET:      executeGET,
		Timeout:         timeout,
		ServerURL:       serverURL,
		TargetServerURL: targetServerURL,
	}

	// Generate markdown
	markdown, err := generator.Generate(data, opts)
	if err != nil {
		return fmt.Errorf("error generating markdown: %w", err)
	}

	// Output
	if outputFile != "" {
		if err := os.WriteFile(outputFile, []byte(markdown), 0644); err != nil {
			return fmt.Errorf("error writing output file: %w", err)
		}
		fmt.Printf("Markdown documentation generated: %s\n", outputFile)
	} else {
		fmt.Print(markdown)
	}

	return nil
}
