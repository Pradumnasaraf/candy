package test

import (
	"encoding/base64"
	"os/exec"
	"strings"
	"testing"
)

// Test encode command
func TestEncodeCmd(t *testing.T) {

	stringToEncode := "password"
	cmd := exec.Command("candy", "encode", stringToEncode)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the output
	expectedOutput := base64.StdEncoding.EncodeToString([]byte(stringToEncode))
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}
}

// Test encode command with decode flag
func TestEncodeCmdWithDecodeFlag(t *testing.T) {

	stringToDecode := "cGFzc3dvcmQ="
	cmd := exec.Command("candy", "encode", "-d", stringToDecode)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the output
	expectedOutput, err := base64.StdEncoding.DecodeString(stringToDecode)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	got := strings.TrimSpace(string(output))
	if got != string(expectedOutput) {
		t.Errorf("expected %v, but got: %v", string(expectedOutput), got)
	}

}
