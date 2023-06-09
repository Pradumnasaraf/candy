package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestKeyValueToJson(t *testing.T) {

	// it should convert a key-value file to json
	cmd := exec.Command("candy", "KVTJ", "-f", "testdata/env")

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
	cmd1 := exec.Command("diff", "testdata/env_output.json", "output.json")
	output, err = cmd1.CombinedOutput()
	if err != nil {
		t.Errorf("Error comparing output.json and testdata/test.json")
	}
	if string(output) != "" {
		t.Errorf("Expected output.json and testdata/test.json to be the same, but got error")
	}

}
