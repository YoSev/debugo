# debugo

<img width="647" src="https://github.com/user-attachments/assets/2b5d9516-ae54-4868-ae5b-c9cef13015cc" alt="debugo" />

debugo is a lightweight Go implementation inspired by the popular [debug-js](https://github.com/debug-js/debug) library. This package aims to provide simple, human-friendly and foremost fine-graded debugging output for Go developers.

![CI](https://github.com/yosev/debugo/actions/workflows/go.test.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- Namespace-based (colored) debugging to categorize log output.
- Toggle debugging on or off via environment variables or programmatically.

## Installation

To install `debugo`, use `go get`:

```bash
go get github.com/yosev/debugo
```

## Usage

Using `debugo` is straightforward. Here's a basic example:

```go
package main

import (
	"github.com/yosev/debugo"
)

func main() {
	debugo.SetDebug("*") // overwrites DEBUGO env
	debug, _ := debugo.New("my:namespace")

	debug("This is a debug message.", 420.69, true, struct {
		Foo string
	}{Foo: "bar"})
}

// outputs: my:namespace This is a debug message. 420.69 true {Foo: bar} +0ms
```

### Environment Variables

You can control which debug namespaces are active using the `DEBUG` environment variable. For example:

```bash
export DEBUGO=my:namespace
```

This will enable debugging for the `my:namespace` namespace. To enable multiple namespaces, separate them with commas:

```bash
export DEBUGO=my:namespace,your:namespace
```

To enable all namespaces, use:

```bash
export DEBUGO=*
```

### Disabling Debugging

To turn off debugging, unset the `DEBUGO` environment variable:

```bash
unset DEBUGO
```

## Configuration

`debugo` allows you to customize its behavior through the `Options` struct. These options let you fine-tune how logs are output and processed.

### Available Options

```go
type Options struct {
    // Force log output independent of given namespace matching (default: false)
    ForceEnable bool
    // Use background colors over foreground colors (default: false)
    UseBackgroundColors bool
    // Use a static color (github.com/fatih/color) (default: random foreground color)
    Color *color.Color
    // Defines the pipe to output to, eg. stdOut (default: stdErr)
    Output *os.File
    // Write log files in their own go routine (maintains order)
    Threaded bool
    // Enable leading timestamps by adding a time format
    Timestamp *Timestamp
}
```

### Using Options

To create a new logger with specific options, use the `NewWithOptions` function:

#### Example

```go
package main

import (
    "github.com/yosev/debugo"
    "github.com/fatih/color"
    "os"
)

func main() {
    options := &debugo.Options{
        ForceEnable:         true,
        UseBackgroundColors: false,
        Color:               color.New(color.FgRed).Add(color.Underline),
        Output:              os.Stdout,
        Threaded:            true,
        Timestamp:           &debugo.Timestamp{Format: time.Kitchen},,
    }

    debug, _ := debugo.NewWithOptions("myapp", options)

    debug("This is a custom debug message with configured options.")
}
```

## Comparison to `debug-js`

While `debugo` is inspired by `debug-js`, it is a simplified version tailored for Go. It does not implement all features of `debug-js`, focusing on core debugging functionality for Go developers.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, feel free to open an issue or submit a pull request.

1. Fork the repository.
2. Create a feature branch: `git checkout -b my-new-feature`.
3. Commit your changes: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

## License

`debugo` is released under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [debug-js](https://github.com/debug-js/debug).
- Thanks to the open-source community for the inspiration and guidance.
