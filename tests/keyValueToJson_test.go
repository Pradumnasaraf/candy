package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestKeyValueToJson tests the KVTJ command
func TestKeyValueToJson(t *testing.T) {

	cmd := exec.Command("candy", "KVTJ", "-f", "testdata/env")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the output.json file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the output file with a new
	cmd = exec.Command("diff", "testdata/env_output.json", "output.json")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing output.json and testdata/test.json")
	}
	if string(output) != "" {
		t.Errorf("Expected output.json and testdata/test.json to be the same, but got error")
	}

}

// TestKeyValueToJson tests the KVTJ command with print flag.
func TestKeyValueToJsonWithPrint(t *testing.T) {

	expectedOutput := `{
  "MONGO": "mongodb://localhost:27017/test",
  "MONGO_DBNAME": "test",
  "NODE_ENV": "development",
  "PORT": "3000"
}`
	cmd := exec.Command("candy", "KVTJ", "-f", "testdata/env", "-p")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output

	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestKeyValueToJson tests the KVTJ command with output flag.
func TestKeyValueToJsonWithOutputFlag(t *testing.T) {

	cmd := exec.Command("candy", "KVTJ", "-f", "testdata/env", "-o", "KVTJ_output.json")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the KVTJ_output.json file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the output file with a new
	cmd = exec.Command("diff", "testdata/env_output.json", "KVTJ_output.json")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing KVTJ_output.json and testdata/test.json")
	}
	if string(output) != "" {
		t.Errorf("Expected KVTJ_output.json and testdata/test.json to be the same, but got error")
	}

}
