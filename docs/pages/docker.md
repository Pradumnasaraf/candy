# docker

docker command provides a set of commands to perform docker related tasks. For now, it only provides a command to generate a **Dockerfile** for a language/framework. It creates a Dockerfile in the current directory with the content depending on the language/framework you choose.

## Usage

```
candy docker [command]
```

### Subcommands

- `dockerfile` - Generate a Dockerfile for a language/framework

> Flags for `dockerfile` subcommand

- `--lang` - Language/Framework for which you want to generate a Dockerfile. **This flag is required**.

Eg:
```bash
candy docker dockerfile --lang goland
canndy docker dockerfile --lang node
```

## Supported Languages/Frameworks

- Go/Golang
- Node
- Python
- Ruby
- Rust