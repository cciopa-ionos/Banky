package account

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAccountListCommand(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	bankyPath := filepath.Join(tmpDir, "banky.json")

	// Example data
	data := `[{"Id":"id1","Name":"Ana","Deposit":1000},{"Id":"id2","Name":"George","Deposit":1500}]`

	err := os.WriteFile(bankyPath, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write banky.json: %v", err)
	}

	// Set temp file
	os.Setenv("BANKY_PATH", bankyPath)

	// Capture the output
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run account list
	AccountCmd.SetArgs([]string{"list"})
	err = AccountCmd.Execute()

	// Close pipe
	w.Close()
	os.Stdout = oldOut

	// Read the output
	outBytes, _ := io.ReadAll(r)
	output := string(outBytes)

	// Check for errors
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Final Check
	if !strings.Contains(output, "Names:") {
		t.Errorf("Output missing 'Names:': %s", output)
	}
	if !strings.Contains(output, "Ana") || !strings.Contains(output, "George") {
		t.Errorf("Output missing expected names 'Ana' or 'George': %s", output)
	}
}
