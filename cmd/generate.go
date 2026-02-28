package cmd

import (
	"fmt"
	"os"
	"strings"

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
	varStrings      []string
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
	generateCmd.Flags().StringArrayVarP(&varStrings, "var", "v", []string{}, "Variable substitutions in format 'key=value' (can be used multiple times)")
	generateCmd.MarkFlagRequired("input")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	// Parse variable substitutions
	vars := make(map[string]string)
	for _, varStr := range varStrings {
		parts := strings.SplitN(varStr, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid variable format: %s (expected key=value)", varStr)
		}
		vars[parts[0]] = parts[1]
	}

	// Create options
	opts := &generator.Options{
		ExecuteGET:      executeGET,
		Timeout:         timeout,
		ServerURL:       serverURL,
		TargetServerURL: targetServerURL,
		Vars:            vars,
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
