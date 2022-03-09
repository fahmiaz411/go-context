package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T){
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue (t *testing.T){
	A := context.Background()

	B := context.WithValue(A, "b", "B")
	C := context.WithValue(A, "c", "C")

	D := context.WithValue(B, "d", "D")
	E := context.WithValue(B, "e", "E")

	F := context.WithValue(C, "f", "F")

	fmt.Println(D,)
	fmt.Println(E,)
	fmt.Println(F,)

	fmt.Println(F.Value("f"))
	fmt.Println(F.Value("c"))
	fmt.Println(F.Value("a"))
}

func CreateCounter (ctx context.Context) chan int {
	destination := make(chan int)

	go func (){
		defer close(destination)
		counter:=1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++	
				time.Sleep(time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T){
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination{
		if n > 10 {
			cancel()
			break
		} else {
			fmt.Println("Counter", n)
		}
	}

	cancel()
		
	// time.Sleep(time.Second)
	
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T){
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, time.Second * 5)
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination{
		if n > 10 {
			cancel()
			break
		} else {
			fmt.Println("Counter", n)
		}
	}
		
	// time.Sleep(time.Second)
	
	fmt.Println(runtime.NumGoroutine())	
}

func TestContextWithDeadline(t *testing.T){
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination{
		if n > 10 {
			cancel()
			break
		} else {
			fmt.Println("Counter", n)
		}
	}
		
	// time.Sleep(time.Second)
	
	fmt.Println(runtime.NumGoroutine())	
}