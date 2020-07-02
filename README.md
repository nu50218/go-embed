# go-embed

## Install

`$ go get -u github.com/nu50218/go-embed`

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nu50218/go-embed"
)

type T1 struct {
	A int
	B time.Duration
	C string
}

type T2 struct {
	A int
	B int64
	D bool
	C string
}

func main() {
	t1 := T1{
		A: 100,
		B: time.Millisecond,
		C: "hoge",
	}
	t2 := T2{}

	printT2(t2)
	// {
	//   A: 0
	//   B: 0
	//   C:
	//   D: false
	// }

	if err := embed.Embed(&t2, &t1); err != nil {
		log.Fatal(err)
	}

	printT2(t2)
	// {
	//   A: 100
	//   B: 1000000
	//   C: hoge
	//   D: false
	// }
}
```
