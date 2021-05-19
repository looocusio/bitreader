# bitreader

Bitreader is Golang library for reading bit with offset and length.

Craete new reader with bytes array, then you can use `SliceToInt` to get result of slice by your offset and length.

## Installation

```bash
go get github.com/looocusio/bitreader
```

## Usage

```go
import (
    "fmt"

    "github.com/looocusio/bitreader"
)

func main() {
    r := bitreader.NewReader([]byte{3, 255})
    result, err := r.SliceToInt(0, 8)
    if err != nil {
        fmt.Printf("failed slice to int: %s", err)
    }
    fmt.Println(result)
    // Output:
    // 3
}

```

## Contributing

Contributions are welcome.
