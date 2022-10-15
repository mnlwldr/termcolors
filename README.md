# whats-this

ANSI color formatting for output in terminal

# usage

```go
fmt.Printf("%sHello World %s\n", TXT_RED, RESET)
fmt.Printf("%sHello World %s\n", BACKGROUND_RED, RESET)
fmt.Printf("%sHello World %s\n", UNDERLINE_RED, RESET)
fmt.Printf("%sHello World %s\n", BOLD_RED, RESET)
```

# colors overview

```go
TXT_BLACK  = "\033[0;30m" // Black - Regular
TXT_RED    = "\033[0;31m" // Red
TXT_GREEN  = "\033[0;32m" // Green
TXT_YELLOW = "\033[0;33m" // Yellow
TXT_BLUE   = "\033[0;34m" // Blue
TXT_PURPLE = "\033[0;35m" // Purple
TXT_CYAN   = "\033[0;36m" // Cyan
TXT_WHITE  = "\033[0;37m" // White

BOLD_BLACK  = "\033[1;30m" // Black - Bold
BOLD_RED    = "\033[1;31m" // Red
BOLD_GREEN  = "\033[1;32m" // Green
BOLD_YELLOW = "\033[1;33m" // Yellow
BOLD_BLUE   = "\033[1;34m" // Blue
BOLD_PURPLE = "\033[1;35m" // Purple
BOLD_CYAN   = "\033[1;36m" // Cyan
BOLD_WHITE  = "\033[1;37m" // White

UNDERLINE_BLACK  = "\033[4;30m" // Black - Underline
UNDERLINE_RED    = "\033[4;31m" // Red
UNDERLINE_GREEN  = "\033[4;32m" // Green
UNDERLINE_YELLOW = "\033[4;33m" // Yellow
UNDERLINE_BLUE   = "\033[4;34m" // Blue
UNDERLINE_PURPLE = "\033[4;35m" // Purple
UNDERLINE_CYAN   = "\033[4;36m" // Cyan
UNDERLINE_WHITE  = "\033[4;37m" // White

BACKGROUND_BLACK  = "\033[40m" // Black - Background
BACKGROUND_RED    = "\033[41m" // Red
BACKGROUND_GREEN  = "\033[42m" // Green
BACKGROUND_YELLOW = "\033[43m" // Yellow
BACKGROUND_BLUE   = "\033[44m" // Blue
BACKGROUND_PURPLE = "\033[45m" // Purple
BACKGROUND_CYAN   = "\033[46m" // Cyan
BACKGROUND_WHITE  = "\033[47m" // White

RESET = "\033[0m" // Text Reset
```

[![Go Reference](https://pkg.go.dev/badge/github.com/mnlwldr/termcolors.svg)](https://pkg.go.dev/github.com/mnlwldr/termcolors)
