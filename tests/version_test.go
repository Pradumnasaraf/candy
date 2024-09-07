package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestVersionCmd tests the version command
func TestVersionCmd(t *testing.T) {

	expectedOutput := "candy version"
	cmd := exec.Command("./candy", "version")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := string(output)[:13]

	if !strings.Contains(got, expectedOutput) {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
