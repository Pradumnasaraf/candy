package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestJsonToYamlCmd(t *testing.T) {

	// it should convert a json file to yaml

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
