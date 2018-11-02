package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	wait := time.After(5 * time.Second)
	<-wait
	fmt.Println("End")
}
