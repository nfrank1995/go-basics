package concurrency

import (
	"fmt"
	"time"
)

func helloChannel(c chan string) {
	c <- "Hello"
	time.Sleep(time.Second * 2)
	c <- "World!"
	close(c)
}

func countdown(s int, ch chan int) {
	for i := 0; i < s; i++ {
		ch <- s - i
		time.Sleep(time.Second)
	}
	ch <- 0
	close(ch)

}

func Run() {
	c := make(chan string)
	ch := make(chan int)

	go helloChannel(c)
	go countdown(10, ch)

	fmt.Println("Response from goroutine:")
	for res := range c {
		fmt.Println(res)
	}

	fmt.Println("Countdown:")
	for res := range ch {
		fmt.Printf("%d\n", res)
	}
}
