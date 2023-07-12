package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestJsonToYamlCmd tests the jsonToYaml command
func TestJsonToYamlCmd(t *testing.T) {

	// Execute the jsonToYaml command
	cmd := exec.Command("candy", "JTY", "-f", "testdata/JTY.json")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expecredOutput := "Operation completed successfully. Check the output.yaml file."
	got := strings.TrimSpace(string(output))
	if got != expecredOutput {
		t.Errorf("expected %v, but got: %v", expecredOutput, got)
	}

	// Validate the output file with a new
	cmd1 := exec.Command("diff", "testdata/JTY_output.yaml", "output.yaml")
	output, err = cmd1.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing output.yaml and testdata/test.yaml")
	}
	if string(output) != "" {
		t.Errorf("Expected output.yaml and testdata/test.yaml to be the same, but got error")
	}

}

// TestJsonToYamlCmd tests the jsonToYaml command with output file flag.
func TestJsonToYamlCmdWithOutfile(t *testing.T) {

	// Execute the jsonToYaml command
	cmd := exec.Command("candy", "JTY", "-f", "testdata/JTY.json", "-o", "JTY_output.yaml")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expecredOutput := "Operation completed successfully. Check the JTY_output.yaml file."
	got := strings.TrimSpace(string(output))
	if got != expecredOutput {
		t.Errorf("expected %v, but got: %v", expecredOutput, got)
	}

	// Validate the output file with a new
	cmd1 := exec.Command("diff", "testdata/JTY_output.yaml", "JTY_output.yaml")
	output, err = cmd1.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing JTY_output.yaml and testdata/test.yaml")
	}
	if string(output) != "" {
		t.Errorf("Expected JTY_output.yaml and testdata/test.yaml to be the same, but got error")
	}

}
