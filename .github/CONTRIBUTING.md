# Contributing

All contributions are welcome! This guide will help you get started with the contribution process.

### Find/Create Issue

Before you start contributing, you need to find an issue you'd like to resolve. Take a look at [issues](https://github.com/alvii147/gloop/issues) that are currently open and unassigned, and see which one you'd like to tackle. Every PR must have a corresponding being resolved.

If you want to report a bug or request a feature, you may also [create an issue](https://github.com/alvii147/gloop/issues/new) yourself.

### Install Go

Install [Go](https://go.dev/) if you haven't already. You can check if you already have Go installed by running:

```bash
go version
```

### Install Git

Install [Git](https://git-scm.com/) if you haven't already. You can check if you already have Git installed by running:

```bash
git --version
```

### Install GNU Make

Install [GNU Make](https://www.gnu.org/software/make/) if you haven't already. You can check if you already have GNU Make installed by running:

```bash
make --version
```

### Fork this Repository

Create a [fork](https://github.com/alvii147/gloop/fork) of this repository on your own GitHub account.

## Make Changes

Clone your forked repository:

```bash
git clone https://github.com/<your-username>/<your-fork>.git
```

Once you've cloned it locally, make your desired changes.

## Update Documentation

If your changes include additions/updates to code, remember to:

* Add/update an example of usage in `example_test.go`
* Add/update the [features section](https://github.com/alvii147/gloop?#features) in `README.md`

## Running Tests

To run all tests with coverage, run:

```
make test
```

For this library, **100% coverage is mandatory.**

To run a specific test, set the `TESTCASE` variable:

```
make test TESTCASE=TestName
```

## Open Pull Request

Open a pull request between your fork and the `main` branch of this repository for your proposed changes. Remember to document the details of your proposal.
