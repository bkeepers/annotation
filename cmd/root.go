package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bkeepers/annotation/parser"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "annotation",
	Short: "Convert linter output into annotations for the GitHub Checks API",
	Run: func(cmd *cobra.Command, args []string) {
		annotations, err := parser.Parse(bufio.NewReader(os.Stdin))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		output, err := json.Marshal(annotations)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s", output)
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
