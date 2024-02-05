package communication

import (
	"fmt"
	"time"
)

func produce(count, intervall int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- i + 1
			time.Sleep(time.Second * time.Duration(intervall))
		}
	}()

	return ch
}

func Run() {
	fastProd := produce(4, 3)
	slowProd := produce(3, 4)

	listen1, listen2 := true, true

	for listen1 && listen2 {
		select {
		case num, ok := <-fastProd:
			if !ok {
				listen1 = false
				return
			}
			fmt.Printf("Item count fast production line: %d\n", num)
		case num, ok := <-slowProd:
			if !ok {
				listen2 = false
				return
			}
			fmt.Printf("Item count slow production line: %d\n", num)
		}
	}

}
