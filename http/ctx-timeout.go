package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	go add_number(ctx, 0)
	subtract_number(ctx, 0)

}

func add_number(ctx context.Context, i int) {
	for {
		i += 1
		select {
		case <-ctx.Done():
			fmt.Printf("add number is done number is : %d \n", i)
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Printf("add number is running, number is : %d \n", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func subtract_number(ctx context.Context, i int) {
	for {
		i -= 1
		select {
		case <-ctx.Done():
			fmt.Printf("subtract number is done number is : %d \n", i)
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Printf("subtract number is running, number is : %d \n", i)
			time.Sleep(1 * time.Second)
		}
	}
}
