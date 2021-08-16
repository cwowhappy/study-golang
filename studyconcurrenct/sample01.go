package studyconcurrenct

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DoneSample() {
	wait := sync.WaitGroup{}
	wait.Add(10)
	done := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i ++ {
			fmt.Printf("Producer process[%d] begin\n", i)
			done<- struct{}{}
			fmt.Printf("Producer process[%d] end\n", i)
		}
	}()

	go func() {
		for i := 1; i <= 10; i ++ {
			fmt.Printf("Consumer process[%d] begin\n", i)
			<- done
			fmt.Printf("Consumer process[%d] end\n", i)
			wait.Done()
		}
	}()

	wait.Wait()
}

func DoneSample01() {
	wait := sync.WaitGroup{}
	wait.Add(1)
	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("Producer process begin\n")
		close(done)
		fmt.Printf("Producer process end\n")
		time.Sleep(3 * time.Second)
		wait.Done()
	}()

	for i := 1; i <= 10; i ++ {
		go func(index int) {
			fmt.Printf("Consumer process[%d] begin\n", index)

			select {
			case <-done:
				fmt.Print()
			}
			fmt.Printf("Consumer process[%d] end\n", index)
		}(i)
	}

	wait.Wait()
}


func CancelContextSample() {
	ctx, cancel := context.WithCancel(context.Background())
	wait := sync.WaitGroup{}
	wait.Add(1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Producer process begin\n")
		cancel()
		fmt.Printf("Producer process end\n")
		time.Sleep(2 * time.Second)
		wait.Done()
	}()

	for i := 1; i <= 10; i ++ {
		go func(index int, cancelCtx context.Context) {
			fmt.Printf("Consumer process[%d] begin\n", index)
			select {
			case <-cancelCtx.Done():
				fmt.Print()
			}
			fmt.Printf("Consumer process[%d] end\n", index)
		}(i, ctx)
	}

	wait.Wait()
}