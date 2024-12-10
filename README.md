# termunicode

This is the Go language version of the popular `is-unicode-supported` package. It provides a simple utility to check if the current terminal or console supports Unicode characters. It performs checks based on environment variables and terminal types, supporting both Windows and Unix-like systems.

## Installation

You can install this package using `go get`:

```bash
go get github.com/mlam-studio/termunicode
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/mlam-studio/termunicode"
)

func main() {
	if termunicode.IsUnicodeSupported() {
		fmt.Println("Unicode is supported")
	} else {
		fmt.Println("Unicode is not supported")
	}
}
```

## Features

- Works on both Windows and Unix-like systems (Linux, macOS).
- Detects terminal support for Unicode based on common terminal emulators.
- Includes checks for Windows Terminal, ConEmu, VS Code, and many more.