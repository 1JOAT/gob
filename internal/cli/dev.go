package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/1joat/gob/internal/ui"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start the development server with hot-reload",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner(Version)
		ui.PrintInfo("Starting development server...")

		mainPath := "./cmd/api"
		mainFile := "./cmd/api/main.go"
		if _, err := os.Stat(mainFile); os.IsNotExist(err) {
			fmt.Printf("Error: %s not found. Please run this from the project root.\n", mainFile)
			return
		}

		// Create .gob directory for the dev binary
		devDir := ".gob"
		if err := os.MkdirAll(devDir, 0755); err != nil {
			fmt.Printf("Error: failed to create %s directory: %v\n", devDir, err)
			return
		}

		devBinary := filepath.Join(devDir, "dev-app")
		if runtime.GOOS == "windows" {
			devBinary += ".exe"
		}

		ui.PrintInfo("Watching for changes in .go files...")
		fmt.Println("Press Ctrl+C to stop.")

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		restart := make(chan bool, 1)

		// Initial start
		go func() {
			restart <- true
		}()

		// File watcher  polling
		go func() {
			lastMod := make(map[string]time.Time)
			for {
				changed := false
				err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
					if err != nil || info.IsDir() || filepath.Ext(path) != ".go" || strings.HasPrefix(path, ".gob") {
						return nil
					}

					if info.ModTime().After(lastMod[path]) {
						if !lastMod[path].IsZero() {
							changed = true
						}
						lastMod[path] = info.ModTime()
					}
					return nil
				})

				if err == nil && changed {
					restart <- true
				}
				time.Sleep(500 * time.Millisecond)
			}
		}()

		var currentProcess *exec.Cmd

		for {
			select {
			case <-stop:
				killProcess(currentProcess)
				fmt.Println("\nStopping development server...")
				return
			case <-restart:
				if currentProcess != nil {
					ui.PrintInfo("Change detected! Rebuilding...")
					killProcess(currentProcess)
					// Small delay to allow Windows to release the port
					time.Sleep(800 * time.Millisecond)
				}

				// Build the binary to a consistent path to avoid the firewall prompt
				buildCmd := exec.Command("go", "build", "-o", devBinary, mainPath)
				buildCmd.Stdout = os.Stdout
				buildCmd.Stderr = os.Stderr
				if err := buildCmd.Run(); err != nil {
					fmt.Printf("Build failed: %v\n", err)
					continue
				}

				currentProcess = exec.Command(devBinary)
				currentProcess.Stdout = os.Stdout
				currentProcess.Stderr = os.Stderr

				if err := currentProcess.Start(); err != nil {
					fmt.Printf("Failed to start process: %v\n", err)
				}
			}
		}
	},
}

func killProcess(cmd *exec.Cmd) {
	if cmd == nil || cmd.Process == nil {
		return
	}

	if runtime.GOOS == "windows" {
		kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
		kill.Run()
	} else {
		cmd.Process.Kill()
	}
}

func init() {
	rootCmd.AddCommand(devCmd)
}
