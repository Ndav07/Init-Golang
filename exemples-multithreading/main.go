package main

import (
	"time"
	"fmt"
)

func main() {
	go task("A")
	go task("B")
	task("R")
	
	canal := make(chan string)

	go func()  {
		canal <- "Veio da t2"
	}()

	fmt.Println(<- canal)
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}