package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C

	fmt.Println(time)
}

// time.After
func TestTimerAfter(t *testing.T) {
	channel := time.After(1 * time.Second)
	fmt.Println(time.Now())

	time := <-channel

	fmt.Println(time)
}

// time.AfterFunc
func TestTimerAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Execute after 1 second")
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()
}
