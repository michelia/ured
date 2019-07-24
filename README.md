# ured(util of redis)

使用 [github.com/mediocregopher/radix]()操作redis.

简单的封装. Do + Cmd + FlatCmd + Pipeline, and type alias.

# 样例
```go
package main

import (
	"github.com/michelia/ured"
	"github.com/michelia/ulog"
)

func main() {
	c := ured.Config{
		Addr:      "127.0.0.1:6379",
		Passwd:    "xxxxx",
		DbNum:     3,
	}
	slog := ulog.NewConsole()
	red := ured.New(slog, c)
	_ = red.Do(nil, "SET", "foo", "someval")
	foo := ""
	_ = red.Do(&foo, "GET", "foo", "someval")
	slog.Print("foo: "foo)
	// print
	// foo: someval
}

```

### 参考
- [https://godoc.org/github.com/mediocregopher/radix]()
