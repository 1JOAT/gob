package cli

import (
	"fmt"
	"os"

	"github.com/1joat/gob/internal/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gob",
	Short: "gob is a CLI for Go web applications",
	Long: `A fast and flexible CLI for scaffolding and managing 
Go web applications with modern aesthetics.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner(Version)
		ui.PrintInfo("Welcome to gob! Use 'gob --help' for available commands.")
	},
}

const Version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gob",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintInfo(fmt.Sprintf("gob version %s", Version))
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
