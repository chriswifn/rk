# 🌳 Go Bonzai File Comparison Branch

File Comparison using a slightly modified version of the Rabin Karp Algorithm
to determine propabilities of plagarism accross student submissions.

## Installation

Standalone
```
go install github.com/chriswifn/rk/cmd/rk@latest
```

Composed

```go
package z

import (
    Z "github.com/rwxrob/bonzai/z"
    "github.com/chriswifn/rk"
)

var Cmd = &Z.Cmd{
    Name: `z`,
    Commands: []*Z.Cmd{help.Cmd, rk.Cmd},
}
```

## Tab Completion
To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C rk rk
```

