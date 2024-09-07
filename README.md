# Candy

<div align="center">

![Candy CLI logo](https://user-images.githubusercontent.com/51878265/224826395-f62efa65-f64c-4c2e-aa93-ad6f72e0d5d7.png)

[![Releases](https://github.com/Pradumnasaraf/candy/actions/workflows/releases.yml/badge.svg)](https://github.com/Pradumnasaraf/candy/actions/workflows/releases.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/Pradumnasaraf/candy.svg)](https://pkg.go.dev/github.com/Pradumnasaraf/candy)

</div>

**Candy** is a CLI tool that provides a basic set of commands to perform tedious tasks such as converting **YAML to JSON** or **JSON to YAML** directly from your terminal. It's built with [Go](https://github.com/golang/go), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), etc.

To install the Candy CLI, use the command `go install github.com/Pradumnasaraf/candy@latest`.
Go will automatically install it in your `$GOPATH/bin` directory, which should be in your `$PATH`.

Once installed, you can use the `candy` CLI command. To confirm installation, type `candy` at the command line.

> **Note** If you are getting an error like `command not found: candy`, then you need to add `$GOPATH/bin` to your `$PATH` environment variable. For that you can refer to [this](https://gist.github.com/Pradumnasaraf/ca6f9a0507089a4c44881446cdda4aa3). 

## 📚 Documentation

**Check out detailed documentation for getting started and using Candy CLI** [**here**](https://pradumnasaraf.github.io/candy/).

## ⭐️ Features

- Convert JSON file to YAML
- Convert YAML file to JSON
- Convert Key-Value to JSON
- Generate Dockerfile for different languages/framework
- Generate Kubernetes manifests for different objects
- Encode and Decode a string to base64
- More coming soon...

## 📝 Usage

```
Usage:
  candy [command] [flags]
  candy [command]

Available Commands:
  JTY         Converts a JSON into YAML.
  KVTJ        Converts Key-Value (text) to JSON.
  YTJ         Converts a YAML into JSON.
  docker      Docker related commands. Like generating a Dockerfile for a language.
  encode      It encodes and decodes a string to base64 and vice versa.
  help        Help about any command
  k8s         Kubernetes related commands. Like generating manifest files for kubernetes objects.
  update      Update candy to the latest version
  version     Know the installed version of candy
```

For detailed usage of each command, visit [here](https://pradumnasaraf.github.io/candy/)

Eg `candy JTY --file test.json` with convert JSON into YAML and create a `output.yaml` in your current directory.

Eg `candy docker dockerfile --lang go` to generate a `Dockerfile` template for go.

eg `candy k8s manifest --obj deployment` to generate a `deployment.yaml` file with deployment template.

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
