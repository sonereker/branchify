# Brancify

Small tool to generate branch name from a Jira issue key.

## Config

Create a config file named `.brancify` in `$HOME` dir with keys;

```
base_url: <JIRA_BASE_URL>
username: <JIRA_USERNAME>
password: <JIRA_PASSWORD>
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