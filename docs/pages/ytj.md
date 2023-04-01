# YTJ - YAML to JSON

YTJ command will convert a YAML file to JSON file. We cam specify the input file and output file.

## Usage

```bash
canndy YTJ [Flags]
```

> Flags for `YTJ` subcommand

- `--file` or `-f` - Input YAML file. **This flag is required**.
- `--output` or `-o` - Output JSON file. By default, it will be `output.json` in the current directory.

Eg:

```bash
candy YTJ --input input.yaml --output output.json
candy YTJ --input user.yaml // output.json will be created in the current directory
```

You can use help flag to get more information about the command.

```bash
candy YTJ --help
```
