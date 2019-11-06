package main

import (
	"fmt"

	"github.com/thsanto/hello/pkg/idrps"
)

func main() {
	generator, err := idrps.NewIDRPS(7)

	if err != nil {
		panic(err)
	}

	id, err := generator.Generate()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", id.Base64)

}
