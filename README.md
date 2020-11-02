# go-junit-report

Allows go applications to generate an xml report suitable for applications expecting JUnit xml reports (e.g. [Jenkins](http://jenkins-ci.org)).

## Installation

Install or update using the `go get` command:

```bash
go get -u github.com/mjidris/go-junit-report
```

## Usage

Applications can import the `junit` package and create a `Report` to track their test results. A report is simply a collection of `Package` structs, which themselves are a collection of `Test` structs. Each individual test case should be represented by a single `Test` struct. Users can track the result and duration of the test case by updating the `Result` and `Duration` properties respectively.

When all test cases are recorded, a report can be generated using the `Write` function which accepts a `Report` struct along with a file name. If the file name does not end with a `.xml` extension, one will be appended.

## Note

This repository was forked from [github.com/jstemmer/go-junit-report](https://github.com/jstemmer/go-junit-report) and has been changed significantly. Any contributions to the original repository can be done there.
