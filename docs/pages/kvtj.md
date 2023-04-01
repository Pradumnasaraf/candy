# KVTJ - Key-Value To JSON

KVTJ command will convert a key-value in text format to JSON format. We can specify the input file and output file.

## Usage

```bash
candy KVTJ [Flags]
```

> Flags for `KVTJ` subcommand

- `--file` or `-f`  - Input text file Key-Value pairs. It can be in a `.txt` format or a `.env` file. **This flag is required**.

> Note: The key-value pairs should be in the format `key=value`. Eg: `name=Pradumna`

- `--output` or `-o` - Output JSON file. By default, it will be `output.json` in the current directory.

```bash
candy KVTJ --input input.txt --output output.json
candy KVTJ --input .env // output.json will be created in the current directory
```

- `--print` or `-p` - Print the JSON output to the console without creating a file.

```bash
candy KVTJ --input input.txt --print
```

You can use help flag to get more information about the command.

```bash
candy KVTJ --help
```