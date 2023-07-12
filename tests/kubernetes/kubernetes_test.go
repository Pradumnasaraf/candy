package docker

import (
	"os/exec"
	"strings"
	"testing"
)

// TestDockerCmd tests the docker command
func TestKubernetesCmd(t *testing.T) {

	expectedOutput := "Kubernetes related commands."

	cmd := exec.Command("candy", "k8s")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:28])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
