package sync_test

import (
	"context"
	"fmt"
	sync "leetcode/golang_features/synctest"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

func TestSharedValue(t *testing.T) {
	synctest.Run(func() {
		var shared atomic.Int64
		go func() {
			shared.Store(1)
			time.Sleep(2 * time.Microsecond)
			shared.Store(2)
		}()

		// Check the shared value after 5 microseconds
		time.Sleep(1 * time.Microsecond)
		synctest.Wait()
		if shared.Load() != 2 {
			t.Errorf("shared = %d, want 2", shared.Load())
		}
	})
}

func TestTimingWithoutSynctest(t *testing.T) {
	synctest.Run(func() {
		start := time.Now().UTC()
		time.Sleep(5 * time.Second)
		t.Log(time.Since(start))
	})
}

func TestAfterFunc(t *testing.T) {
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(context.Background())

		afterFuncCalled := false
		context.AfterFunc(ctx, func() {
			afterFuncCalled = true
		})

		// Cancel the context and wait for the AfterFunc to complete
		cancel()
		synctest.Wait()

		// Now we can safely check that the callback has been called
		fmt.Printf("after context is canceled: afterFuncCalled=%v\n", afterFuncCalled)
	})
}

func TestTimingWithSynctest(t *testing.T) {
	synctest.Run(func() {
		t.Log(time.Now().UnixNano())

		var now int64
		for range 10000000 {
			now = time.Now().UnixNano()
		}

		t.Log(now)
	})
}

// Output:
// 946684800000000000
// 946684800000000000

func TestTimeWait(t *testing.T) {
	synctest.Run(func() {
		doneCh := make(chan bool)
		go func() {
			for range 100 {
				t.Log(time.Now())
				time.Sleep(time.Millisecond)
			}
		}()

		go func() {
			t.Log("time before sleep:", time.Now())
			time.Sleep(time.Second * 10)
			t.Log("time after sleep:", time.Now())
			close(doneCh)
		}()

		<-doneCh
	})
}

func TestDoSomeWork(t *testing.T) {
	synctest.Run(func() {
		sync.DoSomeWork()
	})
}
