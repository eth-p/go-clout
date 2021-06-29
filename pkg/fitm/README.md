# github.com/eth-p/clout/pkg/fitm
(**F**ormatter **i**n **t**he **M**iddle)

`fitm` is a package for parsing and manipulating [fmt](https://golang.org/pkg/fmt/) formatting strings.


## Installation

```go
import (
    "github.com/eth-p/clout/pkg/fitm"
)
```

## Example

This example will wrap all formatting arguments in parentheses.

```go
import "github.com/eth-p/clout/pkg/fitm"

func wrapped(v fitm.Verb, a interface{}) (fitm.Verb, interface{}) {
    return fitm.Preformatted("(" + v.Format(a) + ")")
}

func main() {
    clout.Printf(wrapped, "hello %s", "world")
    // -> "hello (world)"
}
```