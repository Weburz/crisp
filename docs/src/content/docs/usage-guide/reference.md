---
title: Command-Line Interface (CLI) Reference
description: The CLI reference of Crisp
---

Crisp as a CLI tool is rather quite simple and straight-forward to use. It has
been kept that way intentionally and its simplicity is one of the core design
principle of the tool. Hence this section of the documentation provides a
detailed reference of the very list of commands available for `crisp`.

While this section of the documentation exists for an easier reference browsing
over your preferred web browser, it is highly recommended you use `help` command
for quick reference instead.

## Reference

| Command      | Description                                                 |
| ------------ | ----------------------------------------------------------- |
| `completion` | Generate the autocompletion script for the specified shell. |
| `help`       | Help about any command for `crisp`.                         |
| `message`    | Lint a Git commit message using `crisp`.                    |
| `version`    | Print the version and build information of `crisp`.         |

### `completion`

The `crisp completion` subcommand provides the following arguments and the
purposes they fulfill.

| Subcommand   | Description                                        |
| ------------ | -------------------------------------------------- |
| `bash`       | Generate the autocompletion script for Bash.       |
| `fish`       | Generate the autocompletion script Fish.           |
| `powershell` | Generate the autocompletion script for PowerShell. |
| `zsh`        | Generate the autocompletion script Zsh.            |

TODO: Add some examples of its usage.

### `help`

Print a helpful usage message of a command/subcommand.

**Examples**:

```console
crisp help version
```

```console
crisp help message
```

### `message`

Lint a Git commit message using this command in accordance to the
[Conventional Commits v1.0.0 specification](https://conventionalcommits.org).

**Examples**:

```console
crisp message "chore: fix an annoying bug"
```

```console
echo "feat: add an amazing feature" | crisp message --stdin
```

### `version`

Print valuable build and version information of Crisp to `STDOUT` useful for
debugging purposes mostly.

**Examples**:

```console
crisp version
```
