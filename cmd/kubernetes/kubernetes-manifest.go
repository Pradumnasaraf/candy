package kubernetes

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	k8Obj string
)

// kubernetesManifestCmd is the command for generating a manifest file for a kubernetes object
var kubernetesManifestCmd = &cobra.Command{
	Use:   "manifest [flags]",
	Short: "Generates manifest file for different objects.",
	Run: func(cmd *cobra.Command, args []string) {

		k8Obj = strings.ToLower(k8Obj)
		switch k8Obj {
		case "deployment":
			createManifestFile("deployment.yaml", deployment)
		case "pod":
			createManifestFile("pod.yaml", pod)
		case "service":
			createManifestFile("service.yaml", service)
		case "ingress":
			createManifestFile("ingress.yaml", ingress)
		case "secret":
			createManifestFile("secret.yaml", secret)
		case "configmap":
			createManifestFile("configmap.yaml", configmap)
		case "persistentvolume", "pv":
			createManifestFile("persistentvolume.yaml", pv)
		case "persistentvolumeclaim", "pvc":
			createManifestFile("persistentvolumeclaim.yaml", pvc)
		default:
			log.Print("Currently we don't support manifest generation for " + k8Obj + ".")
		}
	}}

func createManifestFile(filename string, obj string) {
	file, err := os.Create(filename)
	checkNilErr(err)

	defer func() { _ = file.Close() }()

	_, err = file.WriteString(obj)
	checkNilErr(err)

	log.Print(filename + " created successfully.")
}

func init() {
	kubernetesManifestCmd.Flags().StringVarP(&k8Obj, "obj", "o", "", "Kubernetes object to generate manifest for.")
	err := kubernetesManifestCmd.MarkFlagRequired("obj")
	checkNilErr(err)

}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
