package kubernetes

import (
	"github.com/spf13/cobra"
)

// KubernetesCmd is the command for Kubernetes related commands.
var KubernetesCmd = &cobra.Command{
	Use:   "k8s [command] [flags]",
	Short: "Kubernetes related commands. Like generating manifest files for kubernetes objects.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		checkNilErr(err)
	},
}

func init() {
	KubernetesCmd.AddCommand(kubernetesManifestCmd)
}
