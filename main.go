package main

import (
	"fmt"

	"github.com/ks6088ts/handson-grpc-go/internal"
)

func main() {
	fmt.Printf("%v, %v\n", internal.Version, internal.Revision)
}
