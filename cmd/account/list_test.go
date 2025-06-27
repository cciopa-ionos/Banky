package account

import (
	"bankycli/internal/core"
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

	err := os.WriteFile(bankyPath, []byte(data), 0666)
	if err != nil {
		t.Fatalf("Failed to write banky.json: %v", err)
	}

	// Set temp file
	os.Setenv("BANKY_PATH", bankyPath)

	// Capture the output
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	AccountCmd.SetArgs([]string{"list"})
	err = AccountCmd.Execute()

	w.Close()
	os.Stdout = oldOut

	outBytes, _ := io.ReadAll(r)
	output := string(outBytes)

	core.Check(err)

	if !strings.Contains(output, "Names:") {
		t.Errorf("Missing 'Names:': %s", output)
	}
	if !strings.Contains(output, "Ana") || !strings.Contains(output, "George") {
		t.Errorf("No 'Ana' or 'George': %s", output)
	}
}
