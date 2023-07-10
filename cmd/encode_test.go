package cmd

import (
	"encoding/base64"
	"os/exec"
	"strings"
	"testing"
)

func TestEncodeCmd(t *testing.T) {
	stringToEncode := "password"

	// Execute the encode command
	cmd := exec.Command("candy", "encode", stringToEncode)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the output
	expectedOutput := base64.StdEncoding.EncodeToString([]byte(stringToEncode))
	got := strings.TrimSpace(string(output)) // Convert output to string and remove leading/trailing spaces
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}
}
