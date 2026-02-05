package ui

import (
	"fmt"
)

// Color
const (
	Reset  = "\033[0m"
	Cyan   = "\033[36m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Bold   = "\033[1m"
)

// PrintBanner displays a colorful splash screen for gob.
func PrintBanner(version string) {
	banner := `
   ____   ___  ____  
 / ___| / _ \| __ ) 
| |  _ | | | |  _ \ 
| |_| || |_| | |_) |
 \____| \___/|____/ %s

   %s%s Modern Go Framework & CLI %s
`
	fmt.Printf(banner, Reset, Bold+Purple, version, Reset)
}

// PrintInfo prints a formatted info message.
func PrintInfo(msg string) {
	fmt.Printf("%s[info] %s%s\n", Cyan, msg, Reset)
}

// PrintSuccess prints a formatted success message.
func PrintSuccess(msg string) {
	fmt.Printf("%s[success] %s%s\n", Blue, msg, Reset)
}
