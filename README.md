# Candy

<div align="center">

![Candy CLI logo](https://user-images.githubusercontent.com/51878265/224826395-f62efa65-f64c-4c2e-aa93-ad6f72e0d5d7.png)

[![Releases](https://github.com/Pradumnasaraf/candy/actions/workflows/releases.yml/badge.svg)](https://github.com/Pradumnasaraf/candy/actions/workflows/releases.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/Pradumnasaraf/candy.svg)](https://pkg.go.dev/github.com/Pradumnasaraf/candy)

</div>

**Candy** is a CLI tool that provides a basic set of commands to perform tedious tasks such as converting **YAML to JSON** or **JSON to YAML** directly from your terminal. It's build with [Go](https://github.com/golang/go), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), etc

To install the Candy CLI, use the command `go install github.com/Pradumnasaraf/candy@latest`.
Go will automatically install it in your `$GOPATH/bin` directory, which should be in your `$PATH`.

Once installed, you can use the `candy` CLI command. To confirm installation, type `candy` at the command line.

> **Note** If you are getting an error like `command not found: candy`, then you need to add `$GOPATH/bin` to your `$PATH` environment variable. For that you can refer to [this](https://gist.github.com/Pradumnasaraf/ca6f9a0507089a4c44881446cdda4aa3) 

## ⭐️ Features

- Convert JSON file to YAML
- Convert YAML file to JSON
- Generate Dockerfile for different languages/framework
- Generate Kubernetes manifests for different objects.
- More comming soon...

## 📝 Usage

```
Usage:
  candy [command]

Available Commands:
  JTY         Converts a JSON into YAML.
  YTJ         Converts a YAML into JSON.
  docker      Docker related commands. Like generating a Dockerfile for a language.
  k8s         Kubernetes related commands. Like generating manifest files for kubernetes objects.
```

Eg `candy JTY --file test.json` with convert JSON into YAML and create a `output.yaml` in you currect directory.

Eg `candy docker dockerfile --lang go` to genrate a `Dockerfile` template for go.

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
