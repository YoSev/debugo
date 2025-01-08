# debugo

debugo is a lightweight Go implementation inspired by the popular [debug-js](https://github.com/debug-js/debug) library. This package aims to provide simple, human-friendly debugging output for Go developers.

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
    debugo.SetNamespace("*") // overwrites DEBUGO env
    debug, _ := debugo.New("my:namespace")

    debug("This is a debug message.")
}
```

### Environment Variables

You can control which debug namespaces are active using the `DEBUG` environment variable. For example:

```bash
export DEBUGO=my:namespace
```

This will enable debugging for the `my:namespace` namespace. To enable multiple namespaces, separate them with commas:

```bash
export DEBUG=my:namespace,your:namespace
```

To enable all namespaces, use:

```bash
export DEBUG=*
```

### Disabling Debugging

To turn off debugging, unset the `DEBUGO` environment variable:

```bash
unset DEBUG
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
