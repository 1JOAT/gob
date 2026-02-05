package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/1joat/gob/internal/ui"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the current project",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner(Version)
		ui.PrintInfo("Building project...")

		// Check for target
		mainPath := "./cmd/api"
		if _, err := os.Stat(mainPath); os.IsNotExist(err) {
			fmt.Printf("Error: %s directory not found.\n", mainPath)
			return
		}

		// Create bin dir if not exists
		if err := os.MkdirAll("bin", 0755); err != nil {
			fmt.Printf("Failed to create bin directory: %v\n", err)
			return
		}

		// Run go build
		buildProcess := exec.Command("go", "build", "-o", "bin/app.exe", mainPath)
		buildProcess.Stdout = os.Stdout
		buildProcess.Stderr = os.Stderr

		if err := buildProcess.Run(); err != nil {
			fmt.Printf("Build failed: %v\n", err)
			return
		}

		ui.PrintSuccess("Build complete! Binary available at ./bin/app.exe")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
