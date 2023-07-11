package test

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestRootCmd(t *testing.T) {

	expectedOutput := "Do all your tedious tasks with a single command"

	cmd := exec.Command("candy")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	got := string(output)[:47]
	if got != expectedOutput {
		fmt.Println(got)
		fmt.Println(expectedOutput)
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
