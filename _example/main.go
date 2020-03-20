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

func printT2(t2 T2) {
	fmt.Println("{")
	fmt.Printf("  A: %v\n", t2.A)
	fmt.Printf("  B: %v\n", t2.B)
	fmt.Printf("  C: %v\n", t2.C)
	fmt.Printf("  D: %v\n", t2.D)
	fmt.Println("}")
}

func main() {
	t1 := T1{
		A: 100,
		B: time.Millisecond,
		C: "hoge",
	}
	t2 := T2{}

	printT2(t2)

	if err := embed.Embed(&t2, &t1); err != nil {
		log.Fatal(err)
	}

	printT2(t2)
}
