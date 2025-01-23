package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	backgroundContext := context.Background()
	fmt.Println(backgroundContext)

	todoContext := context.TODO()
	fmt.Println(todoContext)
}

func TestWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) // Output: F
	fmt.Println(contextF.Value("c")) // Output: C
	fmt.Println(contextF.Value("b")) // Output: nil, different parent A -> C -> F
	fmt.Println(contextA.Value("b")) // Output: nil, parent cant invoke a child value
}

func CreateCounter(ctx context.Context, wg *sync.WaitGroup) chan int {
	destination := make(chan int)
	wg.Add(1)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Counter goroutine stopped")
				wg.Done()
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate slow process
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutine:", runtime.NumGoroutine())

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	destination := CreateCounter(ctx, wg)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 5 {
			break
		}
	}
	// cancel the context to stop the goroutine
	cancel()

	wg.Wait()

	fmt.Println("Total goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutine:", runtime.NumGoroutine())

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // best practice to invoke cancel even when using context.WithTimeout

	destination := CreateCounter(ctx, wg)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	wg.Wait()

	fmt.Println("Total goroutine:", runtime.NumGoroutine())
}
