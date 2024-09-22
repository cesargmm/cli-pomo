# cli-pomo

`cli-pomo` is a simple terminal pomodoro written in Go.

## Requirements

In order to build this project you will need to install the following applications:

* Go (Golang) v1.23.1

## Parameters

You can manually adjust the following parameters:

`--focus` specify how many minutes will the focus session take

`--break` specify how many minutes will the break session take

> [!note]
> If no parameters are specified, the program will use the default values:
> * focus: 25 minutes
> * break: 5 minutes

## Examples

```bash
cli-pomo -focus 45 -break 10
```

