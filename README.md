# go.eth-p.dev/clout
(**C**ommand **L**ine **Out**put)

`clout` is a package that helps you print user-friendly output messages from your Go command line applications.


## Installation

```go
import (
    "go.eth-p.dev/clout"
)
```

## Why clout?

Clout helps you print readable and consistent messages using a familiar `klog`-like API. It's unobtrusive, and designed to enable you to write command line tools without having to worry about the specifics of ANSI colors or stdout/stderr best practices.

### Messages Types

Instead of asking you to figure out which output stream a message should be destined for, `clout` provides with you different types of messages that you can print:

|Constant|Usage|
|:--|:--|
|`Status`|An update to the program's current status.|
|`Info`|An informational message.|
|`Warning`|A warning about a potential issue.|
|`Deprecation`|A warning about a feature which will be removed or unsupported in the future.|
|`Error`|A severe error.|

By default, `clout` will direct these messages into an appropriate output stream. Warning and error messages will go to the standard error, and all other messages will go to the standard output.  

### Configurable Verbosity

Just like `klog`, `clout` supports different verbosity levels. If you want to provide extra debug information without littering the code with `if`-statements, you can do that:

```go
clout.V(2).Infof("Processing %s", file)            // Visible by default.
clout.V(4).Statusf("%s: Unmarshalling yaml", file) // Only visible with verbosity 4 or higher.
```

#### Best Practices

The following table shows the best practices for using verbosity levels. It's based on the [Kubernetes logging best practices](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md):

|Level|For|
|:-:|:--|
|V(0)|Programmer errors, logging extra info about a panic, cli argument handling.|
|V(1)|Information about config, errors.|
|**V(2)**|System state, log messages.|
|V(3)|Extended info about system state changes.|
|V(4)|Logging in "thorny parts of code".|
|V(5)|Trace level verbosity.|

### Color Support

When your terminal supports colors, giant walls of plain text can be unwieldy. `clout` helps you with that by providing color support (Linux/MacOS only) with no extra burden on you:

- Warning messages will be prefixed with "warning: " and automatically colored yellow.
- Deprecation messages will be prefixed with "deprecated: " and automatically colored yellow.
- Error messages will be prefixed with "error: " and automatically colored red.

And if you have any text that you feel should be highlighted to stand out (e.g. paths)? You can simply wrap parameters in a `Highlight`, and `clout` will handle it. 

#### Conditional Colors

Best of all, colors are enabled conditionally. If someone pipes your command's output, colors will be disabled automatically. `clout` even supports the `NO_COLOR` standard ;)



## Example

```go
import (
    "go.eth-p.dev/clout"
    "go.eth-p.dev/clout/pkg/highlight"
)

func main() {
    clout.V(2).Infof("Initializing application...")
    // -> Initializing application...
    
    clout.V(3).Info("Args: %#v", os.Args)
    // This won't print anything, because the default verbosity is level 2.
	
    if len(os.Args) != 2 {
        clout.V(1).Error("not enough arguments")
        // -> error: not enough arguments
        return
    }
    
    clout.V(2).Infof("Hello, %s.", highlight.Cyan(os.Args[1]))
    // -> Hello, SOME_EXECUTABLE_NAME
    // On supported systems, SOME_EXECUTABLE_NAME will be cyan.
}
```

For more detailed examples, feel free to check out the [examples directory](examples).


## License

[MIT License](LICENSE.md)
