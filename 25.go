package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("sleeping...")
	sleep(3 * time.Second)
	fmt.Print(" woke up!")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
