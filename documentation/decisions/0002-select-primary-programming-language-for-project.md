# 2. Select primary programming language for project

Date: 2021-04-16

## Status

Accepted

## Context

### Ideal

Ideally we are using a language that will allow us to conveniently interact with GitHub and extract characteristics of the source represented in the repositories.

### Reality

In reality we don't know what languages have libraries to aid GitHub interaction and are suited to source code analysis.

### Consequences

* complexity dealing with GitHub API hindering progress
* potential performance issues extracting source code characteristics

## Options

### Golang

`+`

* lines of code library: https://github.com/boyter/scc (2.4K Star)
* GitHub library: https://github.com/google/go-github (7.3K Star)
* Git library: https://github.com/go-git/go-git (2.2K Star)
* developer interest
* familiarity with lines of code library `scc`

`-`

* resume driven development

### Python

`+`

* lines of code application: https://github.com/roskakori/pygount (62 Star)
* GitHub library: https://github.com/PyGithub/PyGithub (4.2K Star)
* Git library: https://github.com/gitpython-developers/GitPython (3.1K Star)
* opportunity to build on knowledge of the language

`-`

* `pygount` seems to be an application not a library
* low star count on lines of code repository

### Perl

`+`

* lines of code application: https://github.com/AlDanial/cloc (11.4K Star)
* GitHub library: https://github.com/fayland/perl-net-github (104 Star)
* Git library: https://github.com/book/Git-Repository (29 Star)
* familiarity with lines of code library `cloc`

`-`

* `cloc` seems to be an application not a library
* lack of developer interest

### Rust

`+`

* lines of code library: https://github.com/XAMPPRocky/tokei (4.8K Star)
* GitHub library: https://github.com/github-rs/github-rs (299 Star)
* Git library: https://github.com/rust-lang/git2-rs (834 Star)
* developer interest

`-`

* low star counts on Git/GitHub repositories
* resume driven development

## Decision

Will use Golang for implementing the project behaviours.

## Consequences

* delays in project development due to choice not being familiar to developer
* strong community around language may allow for more contributors
