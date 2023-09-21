# Go Etrier
A minimalistic go error handling.

## Usage Example
```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	etrier "github.com/mandriota/go-etrier"
)

func main() {
	l := log.New(os.Stderr, "", log.LstdFlags) 
	etrier.SetHandler(func (err error) {
		l.Fatalln(err)
	})
	
	fmt.Println(foo())
}

func foo() int {
	return etrier.Try1(strconv.Atoi("foo"))
}
```
