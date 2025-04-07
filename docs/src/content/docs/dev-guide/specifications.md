---
title: Software Requirements Specifications (SRS)
sidebar:
  order: 1
---

This section of the documentations contains the "Software Requirements
Specifications (SRS)" for Crisp. All feature enhancements and related
development on the project will be based on the criterias listed here on this
document. Any other feature request or behaviour of the tool which can be
considered out-of-scope of this document will not be worked upon. In case, a
functionality or behaviour has been heavily requested by community members, the
document will first have to be updated accordingly before development on the
functionality can start taking shape.

## Purpose

Crisp is intended to be used as a tool to lint and enforce a standard practice
to write Git commit messages. It enforces the rules according to the
specifications laid out in the
[Conventional Commits v1.0.0](https://conventionalcommits.org) document. To put
it simply, Crisp is a linter for Git commit messages!

## Scope

The core functionality of Crisp should be to enforce the rules and guidelines
for authoring Git commit messages in accordance to the Conventional Commits
v1.0.0 specifications. Additional functionalities can be added to the project
but they should not deviate from the core functionality.

On top of the core functionality mentioned above, the following functionalities
(and their similarities) should be outright considered out of scope for the
project:

1. Be a complete clone of an existing tool like
   [commitlint](https://commitlint.js.org) and so on.

2. Provide shareable configurations (like `commitlint` does) since it beats the
   purpose of enforcing strict rules.

Exceptions do exist and hence you are welcome to open a discussion on a specific
functionality to be implemented on Crisp.

## References

Crisp stands on the shoulders of giants and development on the project would not
be possible without the existence of the mature tools like the following:

1. [Commitlint](https://commitlint.js.org)
2. [Pre-Commit](https://pre-commit.com)
3. [Git Hooks](https://git-scm.com/docs/githooks)

## Summary

Crisp is a powerful tool designed to help teams maintain consistent and
high-quality Git commit messages by strictly following the Conventional Commits
v1.0.0 standard. Its primary focus is to enforce rules that ensure every commit
message is clear, concise, and properly formatted. While Crisp's main goal is to
improve commit message practices, it won't venture into areas like providing
customizable configurations or replicating features from tools like
`commitlint`. If you're looking for a tool that ensures consistency and helps
streamline your development process, Crisp is the perfect choice.
