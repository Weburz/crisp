# Crisp: A Git Commit Message Linter

<!-- prettier-ignore-start -->
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/Weburz/crisp) ![Discord](https://img.shields.io/discord/1259044007342772256) ![GitHub License](https://img.shields.io/github/license/Weburz/crisp)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Weburz/crisp)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/weburz)
![GitHub Release](https://img.shields.io/github/v/release/Weburz/crisp)
<!-- prettier-ignore-end -->

Crisp is a linter for [git-commit](https://git-scm.com/docs/git-commit) messages
built on top of [Conventional Commits](https://www.conventionalcommits.org). It
was built in response to the following requirements we needed at
[Weburz](https://weburz.com).

1. Require a no nonsense and opionionated linter to enforce an uniform and
   standard commit message style guide.
2. An easy to setup and use binary executable so that we do not have to deal
   with broken Node.js dependencies.

## Usage Guide

There are two primary ways to get started with using Crisp:

1. Using [Pre-Commit](https://pre-commit.com) (**RECOMMENDED**).
2. Downloading and using the executable binary (check the [releases](./releases)
   page.).

For more detailed usage (and development) guidelines on the project, please
refer to the project's official documentations here -
[tech.weburz.com/crisp](https://tech.weburz.com/crisp).

### Using Crisp With Pre-Commit

It is **RECOMMENDED** to use Crisp through Pre-Commit. Follow these instructions
to use the hook accordingly:

1. Create a `.pre-commit-config.yaml` file and add the following contents to it:

   ```yaml
   repos:
     - repo: https://github.com/Weburz/crisp
       rev: <VERSION>
       hooks:
         - id: crisp
           name: lint git-commit messages
   ```

2. Install the pre-commit hook in your local Git repository by invoking:

   ```console
   pre-commit install --install-hook
   ```

3. Thereafter you can either commit some changes to the local Git repository or
   you can directly invoke the hook by running:

   ```console
   pre-commit run --all-files --verbose
   ```

If your Pre-Commit configuration are correct then you will find it invoked and
linted your commit appropriately.

### Using the Binary Executable

To download, setup and use the executable binary directly on your system, follow
these instructions:

1. Download the binary from the [releases](./releases) page.
2. Copy the binary to a location on your filesystem (e.g. `~/.local/bin`) by
   invoking:

   ```console
   mv ~/Downloads/crisp ~/.local/bin
   ```

3. Ensure the path where `crisp` is available is added to `$PATH` by running:

   ```console
   echo $PATH
   ```

   If the path you installed `crisp` (e.g. `~/.local/bin`) to is not listed in
   the output above, then you will probably have to edit your `.bashrc` or
   `.zshrc` file to update the `$PATH`.

## Why Crisp Exists?

We built Crisp at Weburz due to limitations of `commitlint` which were hampering
our development workflows. With it, we expect to resolve some of those concerns
and we hope others outside Weburz will be able to make some use of Crisp as
well!

1. `commitlint` WAS NOT opionionated enough for our requirements. Neither do we
   want to deal with shareable configurations published on
   [npm registry](https://www.npmjs.com). We expected `commitlint` to strictly
   adhere to the "Conventional Commits" specifications but was deeply
   disappointed.

2. Dependency on Node.js was not something we appreciated due to the reasons
   mentioned above. Among other reasons include, requiring to keep a track of
   dependency updates just for a `git-commit` message linter was not a good use
   of our time.

3. We wanted an easier way to use and ship the tool as an executable binary so
   that our developers can quickly up and running using it instead of bickering
   about messaging style guide.

## What Crisp isn't?

Crisp isn't;

1. An enhancement of `commitlint` (even though it is based on its intended
   specifications).

2. It will not have a functionality to comply with the idea of "shareable
   configurations" although a limited set of configuration options are
   acceptable (PRs are welcome of course! <3).

3. A clone of `commitlint` with a similar CLI UI/UX as documented
   [the `commitlint` docs](https://commitlint.js.org/reference/cli.html).

## Usage and Distribution Rights

Crisp is licensed under a free and open-source license which makes it free to
use, copy and distribute its source code as long as you adhere to the terms and
conditions of the licensing terms. For more information on the licensing
details, check out the [LICENSE](./LICENSE) document.
