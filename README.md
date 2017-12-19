# Pandoc based Profile

[![Build Status](https://travis-ci.org/holderbaum/resume.svg?branch=master)](https://travis-ci.org/holderbaum/resume)

This is my personal profile for recruiters or potential clients.
It is continuously built and released using Travis CI.

To build it locally run `./go build`. Just make sure to install a proper LaTeX distribution and pandoc.

## CI/CD Setup

Every push triggers a build on Travis CI. If the commit is tagged, the resulting PDF will also be relased on Github.
