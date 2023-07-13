package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestJsonToYamlCmd tests the JTY command
func TestJsonToYamlCmd(t *testing.T) {

	cmd := exec.Command("candy", "JTY", "-f", "testdata/JTY.json")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the output.yaml file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the output file with a new
	cmd = exec.Command("diff", "testdata/JTY_output.yaml", "output.yaml")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing output.yaml and testdata/test.yaml")
	}
	if string(output) != "" {
		t.Errorf("Expected output.yaml and testdata/test.yaml to be the same, but got error")
	}

}

// TestJsonToYamlCmdWithOutputFlag tests the JTY command with the output flag
func TestJsonToYamlCmdWithOutputFlag(t *testing.T) {

	cmd := exec.Command("candy", "JTY", "-f", "testdata/JTY.json", "-o", "JTY_output.yaml")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the JTY_output.yaml file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the output file with a new
	cmd = exec.Command("diff", "testdata/JTY_output.yaml", "JTY_output.yaml")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing JTY_output.yaml and testdata/test.yaml")
	}
	if string(output) != "" {
		t.Errorf("Expected JTY_output.yaml and testdata/test.yaml to be the same, but got error")
	}

}
