package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestVersionCmd tests the version command
func TestVersionCmd(t *testing.T) {

	expectedOutputString := "1."
	cmd := exec.Command("candy", "version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	got := string(output)[:5]

	if !strings.Contains(got, expectedOutputString) {
		t.Errorf("expected the version to start with %v, but got: %v", expectedOutputString, got)
	}

}
