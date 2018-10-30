package main

import (
	"fmt"
)

type Takler interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g *Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Takler
	talker = &Greeter{"Tom"}
	talker.Talk()
}
