package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/1joat/gob/internal/ui"
	"github.com/spf13/cobra"
)

const repoOwner = "1joat"
const repoName = "gob"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gob to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintInfo("Checking for updates...")

		latestVersion, downloadURL, err := getLatestRelease()
		if err != nil {
			ui.PrintInfo(fmt.Sprintf("Failed to check for updates: %v", err))
			return
		}

		if latestVersion == Version {
			ui.PrintSuccess(fmt.Sprintf("gob is already up to date (v%s)", Version))
			return
		}

		ui.PrintInfo(fmt.Sprintf("New version available: v%s (current: v%s)", latestVersion, Version))
		ui.PrintInfo("Updating...")

		if err := doUpdate(downloadURL); err != nil {
			ui.PrintInfo(fmt.Sprintf("Update failed: %v", err))
			return
		}

		ui.PrintSuccess(fmt.Sprintf("Successfully updated to v%s!", latestVersion))
	},
}

func getLatestRelease() (string, string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("github api returned status: %s", resp.Status)
	}

	var release struct {
		TagName string `json:"tag_name"`
		Assets  []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", "", err
	}

	version := strings.TrimPrefix(release.TagName, "v")

	// Find the correct asset for the current OS/Arch
	targetAsset := fmt.Sprintf("gob_%s_%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		targetAsset += ".exe"
	}

	var downloadURL string
	for _, asset := range release.Assets {
		if asset.Name == targetAsset {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		// If no specific asset found, general message
		return version, "", fmt.Errorf("no binary found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	return version, downloadURL, nil
}

func doUpdate(url string) error {
	// Simple self-update via go install

	ui.PrintInfo("Running 'go install github.com/1joat/gob/cmd/gob@latest'...")
	cmd := exec.Command("go", "install", "github.com/1joat/gob/cmd/gob@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
