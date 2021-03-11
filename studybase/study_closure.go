package studybase

import (
	"fmt"
	"sync"
	"time"
)

func ClosureSample() {
	fmt.Println("main enter")
	wg := sync.WaitGroup{}
	wg.Add(5)
	// i escape to heap, 每个协程都访问共享变量i
	for i := 0; i < 5; i ++ {
		go func(val int) {
			fmt.Printf("%d -> %d \n", i, val)
			wg.Done()
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("main exit")
}

