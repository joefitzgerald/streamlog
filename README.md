# Streamlog

[![wercker status](https://app.wercker.com/status/3c161febd57a5afe98400a45214b02d2/m "wercker status")](https://app.wercker.com/project/bykey/3c161febd57a5afe98400a45214b02d2)

The go standard library includes a default
[log.Logger](http://golang.org/pkg/log/#Logger) that targets `stderr`. This
library provides a StreamLogger that targets both `stdout` and `stderr`.

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
