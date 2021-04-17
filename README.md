# Branchify

[![Go Report Card](https://goreportcard.com/badge/github.com/sonereker/branchify)](https://goreportcard.com/report/github.com/sonereker/branchify)

Small tool to generate branch name from a Jira issue key in the format `[ISSUE_TYPE]/[ISSUE_KEY]-[ISSUE_SUMMARY]`.

## Config

Create a config file named `.branchify` in `$HOME` dir with keys;

```
base_url: <JIRA_BASE_URL>
username: <JIRA_USERNAME>
password: <JIRA_PASSWORD>
```

## Install

```
go get github.com/sonereker/branchify/cmd/branchify
```

## Usage

```
Usage: main [--prefix PREFIX] [--key required KEY REQUIRED]

Options:
  --prefix PREFIX, -p PREFIX
                         Branch type prefix default is /bugfix [default: bugfix/]
  --key required KEY REQUIRED, -k KEY REQUIRED
                         Jira Issue Key
  --help, -h             display this help and exit
```

### Example

```
./branchify -k MYAL-9748
bugfix/MYAL-9748-requested-at-sent-in-two-formats-hence-breaking-the-logging
```
