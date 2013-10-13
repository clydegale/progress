# progress [![Build Status](https://secure.travis-ci.org/jzelinskie/progress.png)](http://travis-ci.org/jzelinskie/progress)

progress is a package to easily generate a progress bar for your CLI application.
Being written in Go, it makes sense that it is made to be safely usable by multiple goroutines.


## docs

Check out the docs at [godoc.org](http://godoc.org/github.com/jzelinskie/progress)

## example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jzelinskie/progress"
)

func main() {
	pb := progress.New(os.Stdin, 10)
	pb.Draw()
	pb.DrawEvery(1 * time.Second)
	for i := 0; i < 10; i++ {
		pb.Increment()
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("\n")
}
```
