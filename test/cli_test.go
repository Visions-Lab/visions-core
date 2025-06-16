package cronmgr_test

import (
	"os/exec"
	"testing"
)

func TestVisionsCoreHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "--help")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run CLI help: %v\nOutput: %s", err, string(output))
	}
	if len(output) == 0 || !contains(string(output), "Usage:") {
		t.Fatalf("Help output missing or incorrect: %s", string(output))
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr)))
}
