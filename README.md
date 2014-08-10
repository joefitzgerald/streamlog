# Streamlog

The go standard library includes a default
[log.Logger](http://golang.org/pkg/log/#Logger) that targets `stderr`. This
library provides a StreamLogger that targets both `stdout` and `stderr`.

## Usage

```go
import github.com/joefitzgerald/streamlog

func main() {
  l := streamlog.New()
  l.Println(streamlog.Out, "Hello")
  l.PrintlnOut("World")
  l.PrintlnErr("#%!!")
}
```
