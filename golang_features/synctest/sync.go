package sync

import (
	"fmt"
	"sync"
	"time"
)

func DoSomeWork() {
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Second *):
				fmt.Println(time.Since(now))
				if time.Since(now) >= time.Second*3 {
					return
				}
			case <-time.After(time.Second * 2):
				panic("FAIL")
			}
		}
	}()

	wg.Wait()
}
