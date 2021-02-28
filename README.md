# random-string
Golang lib to generate a random string efficiently, with specified length / chars

Installation
---------

```go
import "github.com/thewolfnl/random-string"
```

Functions
---------

### func New

```go
func New() string
```

New returns a new random string of the standard length, consisting of
standard characters.

### func NewLen

```go
func NewLen(length int) string
```

NewLen returns a new random string of the provided length, consisting of
standard characters.

### func NewLenChars

```go
func NewLenChars(length int, chars string) string
```

NewLenChars returns a new random string of the provided length, consisting
of the provided string of allowed characters (maximum 63).

# sources
based off of:
- https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
- https://github.com/dchest/uniuri