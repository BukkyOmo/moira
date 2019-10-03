# Contributing to Moira

First off, thanks for taking the time to contribute!

The following is a set of guidelines for contributing to Moira and its packages, which are hosted in the 
[Moira](https://github.com/moira-alert) Organization on GitHub. These are mostly guidelines, not rules. Use your best 
judgment, and feel free to propose changes to this document in a pull request.

**Table of Content**

* [How Can I Contribute?](#how-can-i-contribute)
    * [Submitting an Issue](#submitting-an-issue)
    * [Contributing a Code](#contributing-a-code)
* [Style Guides](#style-guides)
    * [Git Style Guide](#git-style-guide)
    * [Go Code Style Guide](#go-code-style-guide)
    * [Testing Strategy](#testing-strategy)
* [Development](#development)

## How Can I Contribute?

### Submitting an Issue

If you find bugs or you want propose an idea that you think will be suitable for moira you can create an issue. The main 
issue tracker is https://github.com/moira-alert/moira/issues. Issues that belongs for moira backend and web are placed here. 
Issues that are related to python client and ansible role should be placed in issue tracker for related repo.

We have a templates for two types of issues: Feature Request and a Bug. If you want to submit a bug please be patient 
and fill all required fields. This will help us to define the problem. 

Try to follow this advices to fill a great issue:

* Check that issue tracker do not contain issues similar to your.
* Use a clear and descriptive title for the issue to identify the suggestion.
* Provide a step-by-step description of the suggested enhancement in as many details as possible.
* Provide specific examples to demonstrate the steps. Include copy/pasteable snippets which you use in those examples, as Markdown code blocks.
* Describe the current behavior and explain which behavior you expected to see instead and why.
* Include screenshots and animated GIFs which help you demonstrate the steps to reproduce issue.
* Explain why this enhancement would be useful to most Moira users.
* Specify which version of Moira you're using.

## Contributing a Code

If you want to contribute a code you should choose an issue and comment that you want to work on it. If you cannot find 
an issue for your contribution please fill an issue first. Please **DO NOT** work on code before you file an issue and comment
about your wish to work on it.

Here you can find issues that are good for first time contributors.

* [Good First Issue](https://github.com/moira-alert/moira/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22)
* [Help Wanted](https://github.com/moira-alert/moira/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

During implementation and submission of code follow style guides listed below.

## Style Guides

### Git Style Guide

We are using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) philosophy to submit the code.
If you are not familiar with Conventional Commits it's highly recommended to read 
[documentation](https://www.conventionalcommits.org/en/v1.0.0/) about it.

In general the rules for commiting are following:

* Commit message should have three blocks: title, body and footer divided with empty lines.
* At least title should be specified.
* Start first line with type of commit. You can find commit types below.
* Place summary in the first line.
* If applicable specify the scope of commit.
* In body describe the changes you did. Write mostly **why** you did changes instead **what** you changed.
* You can split a body to multiple paragraphs using blank lines.
* In footer place an additional information about your commit. Specify issue number in footer with `Relates` or `Closes` label.

Structure of commit:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

We are using following commit types in our repo:

* **build:** Changes that affect the build system or external dependencies
* **ci:** Changes to our CI configuration files and scripts
* **docs:** Documentation only changes
* **feat:** A new feature
* **fix:** A bug fix
* **perf:** A code change that improves performance
* **refactor:** A code change that neither fixes a bug nor adds a feature
* **style:** Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **test:** Adding missing tests or correcting existing tests

### Go Code Style Guide

We are using [GolangCI-Lint](https://github.com/golangci/golangci-lint) to check our code for following the code style.
Check your code with `make lint` before you will submit it. Code with broken style will not pass the CI and will not be
reviewed. In future core contributors team can change the rules for linter in project.

### Testing Strategy

We are using unit and integration tests to check that Moira behaviour satisfies our expectations and we expect that
contributors will cover with tests changed code as much as possible. 
Use [tests skipping](https://golang.org/pkg/testing/#hdr-Skipping) in short mode for tests that uses database.
If you fix a bug first cover a bug with test that reproduce this bug then fix it.

## Development

To build and run tests you need local redis listening on port 6379.
Easiest way to get redis is via docker:

```bash
docker run -p 6379:6379 -d redis:alpine
```

To run test use ``go test ./...`` or run [GoConvey](http://goconvey.co/):

```bash
go get github.com/smartystreets/goconvey
goconvey
```

For full local deployment of all services, including web, graphite and metrics relay (may be slow on first launch) use:

```bash
docker-compose up
```

Before push your changes don't forget about linter:

```bash
make lint
```
