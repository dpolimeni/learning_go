package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//c := make(chan int)
	c := make(chan int, 50)
	for i := 0; i < 60; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		value := rand.Int()
		c <- value
		fmt.Println("sent", value)
		fmt.Println(len(c))
		if len(c) < 50 {
			time.Sleep(time.Millisecond * 50)
		} else {
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
