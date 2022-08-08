package main

import (
	"fmt"

	"github.com/tamilmaaran/gofirst/greetings"
	"github.com/tamilmaaran/gofirst/hello"

	"github.com/tamilmaaran/gofirst/_01_math"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(greetings.Greet())

	hello.Say()
	hello.Say()

	fmt.Println(_01_math.SayMath())
}
