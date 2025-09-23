---
title: Development Process
description:
  Develop your Go-based CLI app with this guide. Set up the environment, follow
  Git Flow branching, adhere to coding standards, write unit/integration tests,
  and contribute via pull requests. Build and version your app using Semantic
  Versioning.
sidebar:
  order: 3
---

This document outlines the development process for
[Crisp](https://crisp.weburz.com), a command-line interface (CLI) tool
built with the [Go](https://go.dev) programming language. This section of the
documentation provides a detailed process for developing Crisp. Adhering to this
process ensures smooth collaboration and code quality.

Crisp is a linter for enforcing a uniform standard for writing
[Git](https://git-scm.com) commit messages. The goal of this project is to
provide an easy-to-use, efficient tool for linting Git commit messages in
accordance to the [Conventional Commits](https://www.conventionalcommits.org)
specifications. It is written in Go for its simplicity, performance, and ease of
use in building command-line tools.

## Setting Up the Development Environment

To get started with development, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone git@github.com:Weburz/crisp
   cd crisp
   ```

2. **Install the necessary tools**: Ensure you have the following tools
   installed:

   - [Pre-Commit](https://pre-commit.com) to run and manage the Git hooks.
   - [Go](https://go.dev) to develop and build Crisp itself.
   - An Integrated Development Environment (like
     [VSCode](https://code.visualstudio.com) with the
     [necessary plugin](https://code.visualstudio.com/docs/languages/go)
     installed) for writing code.

3. **Install dependencies**: Install all required dependencies including Go
   modules and other tools for linting/formatting:

   ```bash
   task setup
   ```

If you have followed the recommended steps until now, you should then be ready
to proceed with the rest of the instructions for development.

## Branching Strategy

We follow
[GitHub Flow](https://docs.github.com/en/get-started/using-github/github-flow)
for managing branches. Here are the key branches used in the development
process:

- **main**: The stable (and the default) release branch. All feature branches
  are merged here.
- **feature branches**: For each new feature or bug fix, create a feature branch
  from `main`. Name it something descriptive like `feature/add-new-command`.
- **hotfix branches**: For urgent fixes in the production releases, create a
  hotfix branch from `main`.

## Coding Standards

To maintain code consistency across the project, please adhere to the following
guidelines:

- **File structure**: Use a logical file structure to organize the project
  (e.g., `cmd` for CLI commands, `pkg` for core logic, `internal` for internal
  packages).
- **Go naming conventions**: Follow the official Go naming conventions. For
  example, use camelCase for variable names and PascalCase for function names
  and types.
- **Documentation**: Each package and function should have a description of its
  purpose. Use Go's docstring format for all exported functions and types.
- **Error handling**: Always handle errors and return descriptive error messages
  when necessary.

## Testing

We use Go's built-in testing framework to write unit tests and integration
tests.

1. **Unit tests**: Create unit tests for individual functions in the same
   package as the function being tested. Place tests in files ending with
   `_test.go`.

2. **Integration tests**: Test the full functionality of the CLI by simulating
   real-world usage. These tests can be placed in the `test` or `integration`
   directory.

To run the tests, use the following command:

```bash
go test ./...
```

To check for code coverage:

```bash
go test -cover ./...
```

## Building the CLI

To build the CLI application, run the following command:

```bash
go build -o ./bin/crisp
```

This will generate a binary file named `crisp` under the `bin` directory. You
can then run the binary directly to use your CLI tool like so:

```console
./bin/crisp
```

## Versioning

We follow Semantic Versioning for our releases. The versioning format is
`MAJOR.MINOR.PATCH`.

- **MAJOR** version: Incremented for incompatible API changes.
- **MINOR** version: Incremented for adding functionality in a
  backward-compatible manner.
- **PATCH** version: Incremented for backward-compatible fixes.

To tag a release, run:

```bash
git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push --tags
```
