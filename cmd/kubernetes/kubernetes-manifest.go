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
	Use:   "manifest",
	Short: "Generates manifest file for different objects.",
	Run: func(cmd *cobra.Command, args []string) {

		k8Obj = strings.ToLower(k8Obj)
		switch {
		case k8Obj == "deployment":
			createManifestFile("deployment.yaml", deployment)
		case k8Obj == "pod":
			createManifestFile("pod.yaml", pod)
		case k8Obj == "service":
			createManifestFile("service.yaml", service)
		case k8Obj == "ingress":
			createManifestFile("ingress.yaml", ingress)
		case k8Obj == "secret":
			createManifestFile("secret.yaml", secret)
		case k8Obj == "configmap":
			createManifestFile("configmap.yaml", configmap)
		case k8Obj == "persistentvolume" || k8Obj == "pv":
			createManifestFile("persistentvolume.yaml", pv)
		case k8Obj == "persistentvolumeclaim" || k8Obj == "pvc":
			createManifestFile("persistentvolumeclaim.yaml", pvc)
		default:
			log.Print("Currently we don't support manifest generation for " + k8Obj + ".")
		}
	}}

func createManifestFile(filename string, obj string) {
	file, err := os.Create(filename)
	checkNilErr(err)
	
	defer file.Close()

	_, err = file.WriteString(obj)
	checkNilErr(err)

	log.Print(filename + " created successfully.")
}

func init() {
	KubernetesCmd.AddCommand(kubernetesManifestCmd)

	kubernetesManifestCmd.Flags().StringVarP(&k8Obj, "obj", "o", "", "Kubernetes object to generate manifest for.")
	kubernetesManifestCmd.MarkFlagRequired("obj")
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}