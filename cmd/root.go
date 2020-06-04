package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	Text = "text"
	Json = "json"
)

var OutputType string

var rootCmd = &cobra.Command{
	Use:   "cider",
	Short: "A CLI tool for CIDR blocks",
	Long: ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&OutputType, "output", "o", Text, "Output type: text, json")
}

