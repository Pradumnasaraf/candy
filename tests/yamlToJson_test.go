package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestYamlToJsonCmd tests the yamlToJson command
func TestYamlToJsonCmd(t *testing.T) {

	// Execute the yamlToJson command
	cmd := exec.Command("candy", "YTJ", "-f", "testdata/YTJ.yaml")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expecredOutput := "Operation completed successfully. Check the output.json file."
	got := strings.TrimSpace(string(output))
	if got != expecredOutput {
		t.Errorf("expected %v, but got: %v", expecredOutput, got)
	}

	// Validate the output file with a new
	cmd1 := exec.Command("diff", "testdata/YTJ_output.json", "output.json")
	output, err = cmd1.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing output.json and testdata/test.json")
	}
	if string(output) != "" {
		t.Errorf("Expected output.json and testdata/test.json to be the same, but got error")
	}

}

// TestYamlToJsonCmd tests the yamlToJson command with output file flag.
func TestYamlToJsonCmdWithOutfile(t *testing.T) {

	// Execute the yamlToJson command
	cmd := exec.Command("candy", "YTJ", "-f", "testdata/YTJ.yaml", "-o", "YTJ_output.json")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expecredOutput := "Operation completed successfully. Check the YTJ_output.json file."
	got := strings.TrimSpace(string(output))
	if got != expecredOutput {
		t.Errorf("expected %v, but got: %v", expecredOutput, got)
	}

	// Validate the output file with a new
	cmd1 := exec.Command("diff", "testdata/YTJ_output.json", "YTJ_output.json")
	output, err = cmd1.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing YTJ_output.json and testdata/test.json")
	}
	if string(output) != "" {
		t.Errorf("Expected YTJ_output.json and testdata/test.json to be the same, but got error")
	}

}
