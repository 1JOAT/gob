package cli

import (
	"fmt"
	"os"

	"github.com/1joat/gob/internal/scaffold"
	"github.com/spf13/cobra"
)

var dbType string = "mongodb"

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new gob project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := scaffold.GenerateProject(projectName, dbType); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&dbType, "db", "d", "mongodb", "Database type (mongodb, postgres, etc.)")
	rootCmd.AddCommand(newCmd)
}
