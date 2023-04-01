# JTY - JSON to YAML

JYT command will convert a JSON file to YAML file. We cam specify the input file and output file.

## Usage

```
candy JTY [Flags]
```

> Flags for `JTY` subcommand

- `--input` - Input JSON file. **This flag is required**. 
- `--output` - Output YAML file. By default, it will be `output.yaml` in the current directory. 

Eg:

```bash
candy JTY --input input.json --output output.yaml
candy JTY --input user.json // output.yaml will be created in the current directory
```

You can use help flag to get more information about the command.

```bash
candy JTY --help
```
