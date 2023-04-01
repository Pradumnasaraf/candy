# K8s - Kubernetes

k8s provides a set of commands to perform kubernetes related tasks. For now, we can genrate manifests file for different objects like deployment, service, ingress, etc. It creates a manifest file in the current directory with the content depending on the object you choose.

## Usage

```
candy k8s [command]
```

### Subcommands

- `manifest` - Generate a manifest file for a kubernetes object like deployment, service, ingress, etc.

> Flags for `manifest` subcommand

- `--object` or `-o` - Kubernetes object for which you want to generate a manifest file. **This flag is required**.

## Supported Objects

- Deployment
- Service
- Ingress
- ConfigMap
- PersistentVolume
- PersistentVolumeClaim
- Secret

you can use help flag to get more information about the command.

```bash
candy k8s --help
```