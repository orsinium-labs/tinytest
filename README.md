# tinytest

The only assertion library that works with [TinyGo](https://tinygo.org/).

* **Why does nothing else work?** All assertion libraries use [reflect](https://pkg.go.dev/reflect) to compare values and to generate error messages. However, TinyGo [has very limited support for reflection](https://tinygo.org/docs/reference/lang-support/#reflection) which is not enough for any of the assertion libraries, including [testify](https://github.com/stretchr/testify) and [is](https://github.com/matryer/is).
* **How does this work?** Instead of reflection, tinytest uses generics for comparing values and very fmt for generating error messages.
* **Why would I need an assertion library?** Testing with just the standard library is quite verbose. Tests should be easy to write (people already don't like writing tests, don't make it even more difficult) and easy to read and understand.
* **Is it stable?** What is already here is pretty stable and reliable. The API might change slightly if I come up with a better one. New features and assertions might be added over time if there is demand. Error messages can change in any release.

Features:

* 🪶 Zero dependencies
* 🐹 Pure Go
* ✅ Works both with Go and TinyGo
* 🧠 A simple API inspired by [is](https://github.com/matryer/is) and [testify](https://github.com/stretchr/testify)

## 📦 Installation

```bash
go get github.com/orsinium-labs/tinytest
```

## 🔧 Usage

```go
import (
    "testing"
    "github.com/orsinium-labs/tinytest/is"
)

func TestHello(t *testing.T) {
    c := is.NewRelaxed(t)
    is.Equal(c, hello(), "Hello world!")
}
```

Running tests:

```bash
tinygo ./...
```

If you want to negate the check, wrap `c` into `is.Not`:

```go
// asserts that the string is not empty
is.Equal(is.Not(c), hello(), "")
```

If you want to provide an additional error message, use [t.Log](https://pkg.go.dev/testing#T.Log) or [t.Logf](https://pkg.go.dev/testing#T.Logf):

```go
t.Log("greet Joe")
is.Equal(c, Greet("Joe"), "Hello, Joe")
```

Or [t.Run](https://pkg.go.dev/testing#T.Run) to run a subtest with the given name:

```go
t.Run("greet_joe", func(t *testing.T) {
    is.Equal(c, Greet("Joe"), "Hello, Joe")
})
```

📚 Check out [documentation](https://pkg.go.dev/github.com/orsinium-labs/tinytest/is) for the list of available assertions.

## 🙅 Known limitations

If you try to call `testing.T.FailNow` from TinyGo tests, you'll get the following error message:

```text
FailNow is incomplete, requires runtime.Goexit()
```

Which means that currently calling `FailNow` does not interrupt the tests. For this reason, tinytest provide only `NewRelaxed` function that configures tinygo to use `Fail` instead of `FailNow` and do not interrupt tests on failures. In future versions, we might to find a workaround for this limitation and then tinytest will also provide a `New` function that interrupts the tests immediately.

As a workaround, you can explicitly check at critical points if the test has failed and interrupt execution:

```go
if t.Failed() {
    return
}
```
