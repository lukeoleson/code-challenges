package main

import "fmt"

func main() {

	//sequentialFib(10)

	concurrentFib(10)

}

func sequentialFib(n int) {
	f1, f2 := 0, 1
	fmt.Printf("%v\n%v\n", f1, f2)
	for i := 0; i < 10; i++ {
		f1, f2 = f2, f1 + f2
		fmt.Println(f2)
	}
}

// concurrentFib is concurrent in the sense that it uses two go routines to
// calculate the fibonacci numbers - it is not calculating them concurrently
// though - just printing in one routine and calculating in the other.
//
// Example of the generator pattern.
func concurrentFib(n int) {

	for i := range fibonacci(n) {
		fmt.Println(i)
	}
}

func fibonacci(n int) chan int {
	ch := make(chan int)
	go func() {
		for i, j, k := 0, 1, 0; k < n ; k++ {
			ch <-i
			i, j = i + j, i

		}
		close(ch)
	}()
	return ch
}