package studybase

import (
	"fmt"
	"sync"
	"time"
)

func RWMutexWaitQueue()  {
	rwMutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(6)

	read := func(name string) {
		fmt.Printf("read-%s begin\n", name)
		rwMutex.RLock()
		fmt.Printf("read-%s doing\n", name)
		time.Sleep(1 * time.Second)
		rwMutex.RUnlock()
		fmt.Printf("read-%s end\n", name)
		wg.Done()
	}

	write := func(name string) {
		fmt.Printf("write-%s begin\n", name)
		rwMutex.Lock()
		fmt.Printf("write-%s doing\n", name)
		time.Sleep(5 * time.Second)
		rwMutex.Unlock()
		fmt.Printf("write-%s end\n", name)
		wg.Done()
	}

	go read("01")
	go write("01")
	go read("02")
	go read("03")
	go read("04")
	go read("05")

	wg.Wait()
}
