// Copyright © 2025 Ping Identity Corporation

// Package browser provides internal utilities for opening URLs in the system's default browser.
// This package is not part of the public SDK API and should only be used internally by the
// authentication flow implementations.
package browser

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// CanOpen checks if we can open a browser in the current environment.
// Returns true if a display/GUI environment is detected and the browser command exists.
func CanOpen() bool {
	// Check for display environment (Linux/Unix)
	// or assume GUI is available on Windows/macOS
	switch runtime.GOOS {
	case "linux", "freebsd", "openbsd", "netbsd":
		// On Linux/Unix, check if we're in a graphical environment
		// by looking for common display environment variables
		if _, err := exec.LookPath("xdg-open"); err != nil {
			return false
		}
		// Could also check for DISPLAY or WAYLAND_DISPLAY env vars
		// but xdg-open handles that internally
		return true
	case "darwin":
		// macOS: check if 'open' command exists (it always should)
		_, err := exec.LookPath("open")
		return err == nil
	case "windows":
		// Windows: check if 'explorer' command exists
		_, err := exec.LookPath("explorer")
		return err == nil
	default:
		return false
	}
}

// Open attempts to open a URL in the default browser.
// If the browser cannot be opened, an error message is printed to stderr.
func Open(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("explorer", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening browser: %v\nPlease go to the URL manually: %s\n", err, url)
	}
}
