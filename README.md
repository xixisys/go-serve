# go-serve

a simple serve pkg for testing.

### usage

```go
import "github.com/xixisys/go-serve/pkg"

func TestSomeFunc() {
    s := pkg.NewStaticServe(":8000", "test/testdata")
    defer s.End()
    s.Start()

    // your test code
    // enjoy it.
}
```