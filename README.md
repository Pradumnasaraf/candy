# Candy

<div align="center">

![Candy CLI logo](https://user-images.githubusercontent.com/51878265/224826395-f62efa65-f64c-4c2e-aa93-ad6f72e0d5d7.png)

</div>

**Candy** is a CLI tool that provides a basic set of commands to perform tedious tasks such as converting **YAML to JSON** or **JSON to YAML** directly from your terminal. It's build with [Go](https://github.com/golang/go), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), etc

To install the Candy CLI, use the command `go install github.com/Pradumnasaraf/candy@latest`.
Go will automatically install it in your `$GOPATH/bin` directory, which should be in your `$PATH`.

Once installed, you can use the `candy` CLI command. To confirm installation, type `candy` at the command line.

## ⭐️ Features

- Convert JSON file to YAML
- Convert YAML file to JSON
- More comming soon...

## 📝 Usage

```
Usage:
  candy [flags]
  candy [command]

Available Commands:
  JTY         Converts a JSON file to YAML
  YTJ         Converts a YAML file to JSON
  help        Help about any command
```

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
