# Contributing to gob

First off, thank you for considering contributing to gob! It's people like you that make it a great framework.

## How Can I Contribute?

### Reporting Bugs
- Use the Bug Report template to report issues.
- Provide a clear description and steps to reproduce.

### Suggesting Features
- Use the Feature Request template for new ideas.
- Explain the "why" and how it benefits the framework.

### Pull Requests
1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. Ensure the test suite passes.
4. Issue that pull request!

## Development Setup

1. Clone the repo.
2. Run `go mod download`.
3. Build the CLI: `go build -o gob.exe ./cmd/gob`.
4. Test scaffolding: `./gob.exe new test-project`.

## Style Guide
- Follow standard Go formatting (`go fmt`).
- Keep the CLI output professional and text-based (no emojis).
- Document new exported functions.
