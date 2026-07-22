package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// ==================================================
// Go Cheat Sheet #3 (Concurrency & standard library)
// ==================================================

func title(s string) {
	fmt.Printf("\n===== %s =====\n", s)
}

func goroutineDemo() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Worker 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Worker 2")
	}()

	wg.Wait()
}

func channelDemo() {
	ch := make(chan string)

	go func() {
		ch <- "Hello Channel"
	}()

	fmt.Println(<-ch)
}

func bufferedChannelDemo() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println(<-ch, <-ch)
}

func selectDemo() {
	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "Done"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

func mutexDemo() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func atomicDemo() {
	var value int64

	atomic.AddInt64(&value, 1)
	atomic.AddInt64(&value, 5)

	fmt.Println(value)
}

func contextDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func httpServerDemo() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Go")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go server.ListenAndServe()

	time.Sleep(100 * time.Millisecond)

	server.Shutdown(context.Background())
}

func envDemo() {
	os.Setenv("APP_NAME", "GoDemo")
	fmt.Println(os.Getenv("APP_NAME"))
}

func tickerDemo() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("Tick")
	}
}

func timerDemo() {
	timer := time.NewTimer(100 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer expired")
}

func loggingDemo() {
	log.Println("Application started")
}

func main() {

	title("Goroutines")
	goroutineDemo()

	title("Channels")
	channelDemo()

	title("Buffered Channels")
	bufferedChannelDemo()

	title("Select")
	selectDemo()

	title("Mutex")
	mutexDemo()

	title("Atomic")
	atomicDemo()

	title("Context")
	contextDemo()

	title("Timer")
	timerDemo()

	title("Ticker")
	tickerDemo()

	title("Environment Variables")
	envDemo()

	title("Logging")
	loggingDemo()

	title("HTTP Server")
	httpServerDemo()

	title("Done")
}
