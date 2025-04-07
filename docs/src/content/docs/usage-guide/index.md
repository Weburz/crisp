---
title: Getting Started
description: Get started with using Crisp to lint your commit messages.
---

Getting started with using Crisp is very straight-forward and not very involved
in any way! To get started with using Crisp for your requirements, you have the
following options:

1. Use as a [Pre-Commit](https://pre-commit.com) hook (**RECOMMENDED**).
2. Use as a standalone tool, either built from source or downloaded using a
   package manager ([Homebrew](https://brew.sh), APT and so on) (**docs are
   WIP**, hence contributions are welcome).
3. Use in a CI environment (**docs are WIP**, hence contributions are welcome).

The following sections of the documentations will walk you through the various
ways to get started with using Crisp.

## Pre-Commit Hook

The easiest and most straight-forward approach for using Crisp is to utilise it
as a Pre-Commit hook. So, the following steps will help you with the process:

1. Ensure Pre-Commit is locally installed and accessible (following the
   [official installation guidelines](https://pre-commit.com#installation)).

2. Create a `.pre-commit-config.yaml` file with the following contents:

   ```yaml
   repos:
     - repo: https://github.com/Weburz/crisp
       rev: "v0.1.5"
       hooks:
         - id: crisp
   ```

3. Setup the hooks for your project by executing:

   ```console
   pre-commit install --install-hooks
   ```

4. To test out whether Crisp is working as part of your Pre-Commit hooks, try
   adding a dummy commit like so:

   ```console
   touch .gitkeep
   git add .gitkeep
   git commit -m 'chore: added a dummy .gitkeep file'
   ```

If Crisp was setup properly, you will see either a passing/failing message on
STDOUT. If the output is a failing message then your commit message obviously
does not follow the
[Conventional Commits specifications](https://conventionalcommits.org) and you
will have to improve it accordingly.

## Install Using a Package Manager

Crisp can also be used as a standalone tool, installed using your OS's package
manager of choice. Hence, this section of the docs will walk you through the
process of getting started with its usage accordingly.

Crisp is available on all major platforms including some of the major Linux
distributions, MacOS and Windows. To get Crisp through your preferred choice of
package manager on the platform you are using, follow the relevant instructions
provided below:

### Homebrew (TBA)

### Winget (TBA)

### Debian/Ubuntu (TBA)

### AUR (TBA)

### Fedora, RHEL, Centos (TBA)

Post installation, run the following command to test whether Crisp was installed
appropriately and is usable;

```console
crisp version
```

The command above should print useful build information relevant to the platform
you are using Crisp on which can also prove valuable during debugging sessions
if you are facing issues with it.

## Build From Source

Crisp can also be built from source and used accordingly. The added benefit of
this approach is you will get to the test out the absolute latest
functionalities of the tool before they are stabilised and shipped. Follow these
steps if you are interested in building Crisp from source code:

1. Clone the source code's remote repository from our GitHub organisation by
   invoking:

   ```console
   git clone --depth 1 git@github.com:Weburz/crisp
   ```

2. Ensure you have [Go](https://go.dev) and [Task](https://taskfile.dev)
   installed. Check it by invoking the following commands:

   ```console
   go version
   ```

   ```console
   task --version
   ```

3. When you are ready with the build tools, invoke the following build command;

   ```console
   task build
   ```

   The command will build the executable binary in the `bin` directory.

4. Copy the built executable binary to the `~/.local/bin` directory (on
   Linux/MacOS) or `C:\Program Files\Crisp` (on Windows).

   **Linux/MacOS**:

   ```console
   mv ./bin/crisp ~/.local/bin/crisp
   ```

   **Windows**:

   ```powershell
   Copy-Item ".\bin\crisp" -Destination "C:\Program Files\Crisp"
   ```

5. With the executable binary copied to an appropriate location, you should now
   update the `$PATH` variable of your system so that the `crisp` command can be
   invoked from anywhere on your filesystem.

   **Linux/MacOS**:

   To update the `$PATH`, you will have to add the following lines to the
   `.bashrc`/`.zshrc` file (only on Linux and MacOS systems):

   ```sh
   export PATH="$PATH:$HOME/.local/bin"
   ```

   **Windows**:

   Run the following command in a PowerShell session;

   ```powershell
   $ENV:PATH += ";C:\Program Files\Crisp"
   ```

## Usage in CI Environments

TODO: Add docs on using Crisp to lint PR titles and more
