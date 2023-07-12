package test

import (
	"os/exec"
	"strings"
	"testing"
)

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
