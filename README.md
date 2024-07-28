# Crisp: A Git Commit Message Linter

<!-- prettier-ignore-start -->
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/Weburz/crisp) ![Discord](https://img.shields.io/discord/1259044007342772256) ![GitHub License](https://img.shields.io/github/license/Weburz/crisp)
<!-- prettier-ignore-end -->

Crisp is an opinionated linter for
[`-git-commit`](https://git-scm.com/docs/git-commit) messages built on top of
the [Conventional Commits](https://www.conventionalcommits.org) specifications.
It is built in response to cater to certain requirements that
[`commitlint`](https://commitlint.js.org) could not fulfill our needs at
[Weburz](https://weburz.com).

**NOTE**: Crisp is still a **WORK-IN-PROGRESS** project and **IS NOT**
recommended for usage in production environments yet!

## Why Crisp Exists?

We built Crisp at Weburz due to limitations of `commitlint` which were hampering
our development workflows. Hence, with Crisp, we expect to resolve some of those
concerns and we hope others outside Weburz will be able to make some use of
Crisp as well!

1. **Unopinionated nature of `commitlint`**: When using `commitlint` for some of
   our projects we realised `commitlint` does not strictly adhere to its
   specifications! It is possible to share custom configurations as `npm`
   packages for everyone else to use. We believe a linter should be strict and
   opionionated, hence Crisp will strictly adhere to the Conventional Commits
   specs sheet!
2. **Dependency on Node.js**: A minor concern for us when using `commitlint` was
   its dependency on [Node.js](https://nodejs.org) for its runtime. We want our
   CLI tools to be lightweight and be shipped as executable binaries!
3. **Self-sufficient, lightweight and ease-to-use**: With respect to the
   aforementioned pointers, `commitlint` can be time-consuming and (kind of)
   difficult to use with all its configuration options. We want to make it
   easier to use Crisp so that developers can focus on writing code instead of
   configuring (and debating about) a linter.

## What Crisp isn't?

Crisp isn't;

1. An enhancement of `commitlint` (even though it is based on its intended specs
   sheet).
2. It will not have a functionality to comply with the idea of "shareable
   configurations".
3. A clone of `commitlint` with a similar CLI UI/UX as documented
   [here in the `commitlint` docs](https://commitlint.js.org/reference/cli.html).

## Usage and Distribution Rights

Crisp is licensed under a free and open-source license which makes it free to
use, copy and distribute its source code as long as you adhere to the terms and
conditions of the licensing terms. For more information on the licensing
details, check out the [LICENSE](./LICENSE) document.
