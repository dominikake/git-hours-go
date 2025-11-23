![logo](figures/git-hours.svg)

Git-hours is a command that calculates working time using the git log data.

## Installation

### Install using Golang
```bash
go install github.com/dominikake/git-hours-go@latest
```

### Build from source
```bash
git clone https://github.com/dominikake/git-hours-go.git
cd git-hours-go
go build
```

## Set Environment
Put the built binary into the bin folder, which is set in the $PATH environment variable.
Then it will be recognized as a subcommand of git.
Because if a command starts with 'git-', git automatically recognizes it as a subcommand of git.

## How to use

1. Open the terminal
1. Move to your git repository
1. Type as shown below

```bash
$ git hours
From "2019-03-01 00:00:00 +0900" to "2019-03-31 23:59:59 +0900" : 13h20m9s
```
- The value of timezone offset is automatically set depending on your region.
- By default, the start date is set as the first day of last month and the end date is set as the last day of last month. 

## Detail Options

### Help
```
$ git hours -help


  -author string
    	author name
  -auto-dates
    	use first and last commit dates as range
  -before string
    	before date (default "2020-03-31 23:59:59 +0900")
  -debug
    	debug mode
  -duration string
    	git log duration (default "1h")
  -help
    	print help
  -since string
    	since(after) date (default "2020-03-01 00:00:00 +0900")
```

### Auto-dates
Use the `-auto-dates` flag to automatically set the date range from the first to the last commit in the repository. This is useful when you want to analyze the entire history of a project without manually specifying date ranges.

```bash
$ git hours -auto-dates
```

### Since, Before
You can set the start date and the end date with this options.
If you don't enter any value, the start date and the end date will be set as the first and last date of last month by default.

```bash
$ git hours -since 2019-02-01 -before today
```

If you want to set timezone value, put timezone value at the end of the command as shown below.

```bash
$ git hours -since "2019-03-29 13:55:33 +0800"
```

If you want to customize every value, enter the date, time, and timezone offset value as shown below.

```bash
$ git hours -since "2019-03-01 00:00:00 +0900" -before "2019-03-31 23:59:59 +0900"
```

### Author
If you want to know data of particular user,  use the `-author` option as shown below.

```bash
$ git hours -author name
```

Also, You can set more than one user as shown below.

```bash
$ git hours -author name1,name2
```

### Duration
Git-hours calculates working time based on duration. If interval between git commits is less than duration value, Git-hours considers working time was continued. With `-duration` option, you can set duration as you want.
By default, duration is set to 1 hour.

If you want to set duration to 30min, type as shown below.

```bash
$ git hours -duration 0.5h
```

### Debug
With `-debug` option, You can see the details.
You can see information as following:

- interval between git commits
- time
- author
- commit message

```bash
$ git hours -debug
```

Output example:
```
	 2019-03-31 23:28:40 +0900 kim hanwoong edit go fmt
2m26s >
	 2019-03-31 23:31:06 +0900 kim hanwoong added description.
6m34s >
	 2019-03-31 23:37:40 +0900 kim hanwoong edit comment
1m46s >
	 2019-03-31 23:39:26 +0900 hanwoong kim Update README.md
38s >
	 2019-03-31 23:40:04 +0900 hanwoong kim Update README.md
1m12s >
	 2019-03-31 23:41:16 +0900 hanwoong kim Update README.md
From 2019-02-01 to 2019-03-31 : 13h1m48s
```


## Why did I fork it?
I found git-hours to be a useful utility for a work project where I needed to track development time across repositories. The original tool by lazypic provided a solid foundation for calculating working time from git commit history.

I started adding a feature to automatically calculate date ranges based on the actual commit history in a repository, which evolved into the `-auto-dates` flag. This enhancement makes it easier to analyze entire project histories without manually specifying date ranges.

All credit for the original implementation and concept goes to lazypic. This fork simply builds upon their excellent work with additional functionality that I found useful for my use case.
