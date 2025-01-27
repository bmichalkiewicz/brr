# brr

![logo](assets/logo.png)

`brr` is a tool to fetch and process repositories, producing a ready-to-go `repositories.yaml` file for use with the aww CLI.
## Features

- Fetch repositories from GitLab groups and projects.
- Save repository information in a structured YAML format for easy integration with other tools.

## YAML Structure

The `brr` tool writes repository data in the following structure:

```yaml
- name: <group_name>
  skip: <true|false>
  projects:
    - url: <project_name_1>
    - url: <project_name_2>
    - ...
```

# SYNOPSIS

brr

```
[--output|-o]=[value]
[--update|-p]
```

**Usage**:

```
brr [GLOBAL OPTIONS] [command [COMMAND OPTIONS]] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--output, -o**="": Path where `brr` outputs the repositories.yaml

**--update, -p**: Updates groups already present in the repository file.


# COMMANDS

## gitlab

Download GitLab repositories

**--groups, -g**="": A comma-separated list of GitLab group names to fetch repositories. (default: [])

**--token, -t**="": GitLab API token for authentication.

**--url, -u**="": URL of the GitLab instance.  (default: https://gitlab.com)

## add

Add a repository to the repositories file.

**--url**="": The SSH URL of the repository to add.
