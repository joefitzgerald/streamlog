# Streamlog

[![wercker status](https://app.wercker.com/status/3c161febd57a5afe98400a45214b02d2/m "wercker status")](https://app.wercker.com/project/bykey/3c161febd57a5afe98400a45214b02d2)

## Overview

The go standard library includes a default
[log.Logger](http://golang.org/pkg/log/#Logger) that targets `stderr`. Sometimes
when writing a command line interface (CLI), the output to `stderr and `stdout`
is significant and you want to write tests to verify the output is correct.

## Differences With The Standard Library `log.Logger`

`streamlog.Logger` drops the `Fatal`, `Fatalf`, `Fatalln`, `Panic`, `Panicf`,
and `Panicln` methods from `log.Logger`, and modifies the `Print`, `Printf`, and
`Println` methods to include a `streamlog.OutputStream` argument (`streamlog.Out`
and `streamlog.Err` are possible values). It also adds `PrintOut`, `PrintfOut`,
`PrintlnOut`, `PrintErr`, `PrintfErr`, and `PrintlnErr` methods.

A `streamlog.Logger` is included that internally has an `out` and an
`err` `log.Logger`.

* `streamlog.New()`: Provides you a `Logger` targeting `os.Stdout` and `os.Stderr`
  with a prefix of `""` and flags of `log.LstdFlags`
* `streamlog.NewWithWriters(...)`: Provides you a `Logger` targeting the supplied
  out and err `io.Writer`, prefix, and flags
* `streamlog.NewWithWriters(...)`: Provides you a `Logger` targeting the supplied
  out and err `standardlog.Logger`, prefix, and flags

Note: `standardlog.Logger` is satisfied by the standard library `log.Logger`.

`streamlog.NewWithWriters` provides an easy way to write tests to verify your
log output. `streamlog.NewWithLoggers` provides an easy way to completely change
the behavior of the logging.

## Usage

`go get -u github.com/joefitzgerald/streamlog`

```go
package main

import "github.com/joefitzgerald/streamlog"

func main() {
  l := streamlog.New()
  l.SetPrefix("") // You can set the prefix of the underlying log.Loggers
  l.SetFlags(0)   // You can set the flags of the underlying log.Loggers
  l.Println(streamlog.Out, "Hello")
  l.PrintlnOut("World")
  l.Println(streamlog.Err, "Earth")
  l.PrintlnErr("#%!!")
  logStuffOut(l, "$")
}

func logStuffOut(l streamlog.Logger, s string) {
  l.Println(streamlog.Out, s)
}
```
