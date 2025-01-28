# brr

![logo](assets/logo.png)

`brr` is a tool to fetch and process repositories, producing a ready-to-go `repositories.yaml` file for use with the [aww](https://github.com/bmichalkiewicz/aww) CLI.
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

## Commands

```bash
# example: brr gitlab -g "terraform" -t "ghyr$jrnjyuehd2"
brr gitlab --help

# example: brr git git@github.com:bmichalkiewicz/brr.git
brr add --help
```
