package scaffold

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/1joat/gob/internal/ui"
)

//go:embed templates/*
var templates embed.FS

func GenerateProject(name, dbType string) error {
	ui.PrintInfo(fmt.Sprintf("Scaffolding new project: %s with %s database...", name, dbType))

	// Create project directory
	if err := os.MkdirAll(name, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Walk through embedded templates
	err := fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "templates" {
			return nil
		}

		// Calculate relative path from "templates" root
		relPath := strings.TrimPrefix(path, "templates/")

		// Ensure path uses platform-specific separators for target
		relPath = filepath.FromSlash(relPath)

		// If the file is a .tmpl, remove the suffix for the target path
		targetRelPath := strings.TrimSuffix(relPath, ".tmpl")
		targetPath := filepath.Join(name, targetRelPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// Read file content
		content, err := templates.ReadFile(path)
		if err != nil {
			return err
		}

		// Basic template replacement for specific files
		if strings.HasSuffix(targetRelPath, "go.mod") || strings.HasSuffix(targetRelPath, "main.go") {
			strContent := string(content)
			// Replace the template package name with the new project name
			strContent = strings.ReplaceAll(strContent, "github.com/1joat/gob/internal/scaffold/templates", name)
			content = []byte(strContent)
		}

		fmt.Printf("   Creating %s...\n", targetRelPath)
		// Write to target
		return os.WriteFile(targetPath, content, 0644)
	})

	if err != nil {
		return fmt.Errorf("failed to scaffold project: %w", err)
	}

	ui.PrintSuccess("Project created successfully!")
	fmt.Printf("   cd %s\n", name)
	fmt.Println("   go mod tidy")
	return nil
}
