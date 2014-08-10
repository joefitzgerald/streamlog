# Streamlog

[![wercker status](https://app.wercker.com/status/3c161febd57a5afe98400a45214b02d2/m "wercker status")](https://app.wercker.com/project/bykey/3c161febd57a5afe98400a45214b02d2)

The go standard library includes a default
[log.Logger](http://golang.org/pkg/log/#Logger) that targets `stderr`. This
library provides a StreamLogger that targets both `stdout` and `stderr`.

## Usage

`go get -u github.com/joefitzgerald/streamlog`

```go
import "github.com/joefitzgerald/streamlog"

func main() {
  l := streamlog.New()
  l.Println(streamlog.Out, "Hello")
  l.PrintlnOut("World")
  l.Println(streamlog.Err, "Earth")
  l.PrintlnErr("#%!!")
}
```
