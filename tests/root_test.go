package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestRootCmd tests the root command
func TestRootCmd(t *testing.T) {

	expectedOutput := "Do all your tedious tasks with a single command"

	cmd := exec.Command("candy")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	got := strings.TrimSpace(string(output)[:47])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestRootCmdHelp tests the help flag for the root command
func TestRootCmdHelp(t *testing.T) {

	expectedOutput := "Do all your tedious tasks with a single command"

	cmd := exec.Command("candy", "--help")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	got := strings.TrimSpace(string(output)[:47])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
