package main

import (
	"fmt"

	"github.com/tamilmaaran/gofirsty/greetings"
	"github.com/tamilmaaran/gofirsty/hello"

	"github.com/tamilmaaran/gofirsty/_01_math"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(greetings.Greet())

	hello.Say()
	hello.Say()

	fmt.Println(_01_math.SayMath())
}
