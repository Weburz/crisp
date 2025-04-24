---
title: Architectural Overview
---

The internal logic that powers Crisp are divided into three distinct individual
parts:

1. The **reader** which is responsible for reading the Git commit message either
   from STDIN (piped in) or from the `$GIT_DIR/COMMIT_EDITMSG` file (see the
   [official Git docs](https://git-scm.com/docs/git-commit/2.2.3#Documentation/git-commit.txt-codeGITDIRCOMMITEDITMSGcode)).

2. The `parser` which is responsible for accepting the inputted commit message
   and then parsing it into a Go struct for further data processing.

3. The `validator` which is responsible for running some validation logic on the
   parsed data.

Under the hood, all three of the aforementioned components work in tandem to
lint your Git commit messages. The following diagram will provide a better
understanding of the underlying logic of the software.

```mermaid
sequenceDiagram
    autonumber

    actor User

    box rgb(117, 115, 109) Crisp Internals
        participant Reader
        participant Parser
        participant Validator
    end

    Note over User,Reader: The commit message linting should<br/>ideally be done using the<br/>Pre-Commit framework!
    User ->> Reader : Make a Git commit

    activate User

    activate Reader
        alt read from STDIN
            Reader -->> Parser : Read from piped input
        else read from COMMIT_EDITMSG
            Reader -->> Parser : Read from file
        end
        opt pass message as flag
            Reader -->> Parser : Read from CLI flag
        end
    deactivate Reader

    activate Parser
        Parser -->> Parser : Parse the message into a struct
        Parser -->> Validator : Crisp validates the message
    deactivate Parser

    activate Validator

            Validator ->> User : Crisp responds with valid/invalid message
    deactivate Validator

    deactivate User
```
