package account

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAccountListCommand(t *testing.T) {
	tmpDir := t.TempDir()
	bankyPath := filepath.Join(tmpDir, "banky.json")

	data := `[{"Id":"id1","Name":"Ana","Deposit":1000},{"Id":"id2","Name":"George","Deposit":1500}]`

	if err := os.WriteFile(bankyPath, []byte(data), 0666); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", bankyPath, err)
		os.Exit(1)
	}

	// Set temp file
	if err := os.Setenv("BANKY_PATH", bankyPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting BANKY_PATH: %v\n", err)
		os.Exit(1)
	}

	// Capture the output
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating pipe: %v\n", err)
		os.Exit(1)
	}
	os.Stdout = w

	AccountCmd.SetArgs([]string{"list"})
	if err := AccountCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}

	w.Close()
	os.Stdout = oldOut

	outBytes, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from pipe: %v\n", err)
		os.Exit(1)
	}
	output := string(outBytes)

	if !strings.Contains(output, "Names:") {
		fmt.Fprintf(os.Stderr, "Output missing 'Names:': %s\n", output)
		os.Exit(1)
	}
	if !strings.Contains(output, "Ana") || !strings.Contains(output, "George") {
		fmt.Fprintf(os.Stderr, "Output missing 'Ana' or 'George': %s\n", output)
		os.Exit(1)
	}
}
