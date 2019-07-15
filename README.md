# Environmentalist

`[tree_hugging_gopher.jpg]`

## Usage

```
Usage:
  environmentalist [flags]
  environmentalist [command]

Available Commands:
  help        Help about any command
  run         run the environmentalist daemon
  stop        stop the envrionmentalist daemon

Flags:
  -h, --help                help for environmentalist
  -n, --node                the node flag tells environmentalist that we want to watch files associated with NodeJS
  -p, --php                 the php flag tells environmentalist that we want to watch files associated with PHP
  -v, --vault-type string   the vault flag tells environmentalist what vault we want to extract secrects from (default "hashicorp-vault")
      --version             version for environmentalist

Use "environmentalist [command] --help" for more information about a command.
```