package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for {
			if time.Now().Second()%9 == 0 {
				cancel()
			}
		}
	}()

	go add(ctx, 1)
	sub(ctx, 1)
}

func add(ctx context.Context, i int) {
	for {
		i += 1
		select {
		case <-ctx.Done():
			fmt.Printf("add is cancelled ! num is : %d \n", i)
			time.Sleep(2 * time.Second)
			return
		default:
			{
				fmt.Printf("add is processing ! num is : %d \n", i)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func sub(ctx context.Context, i int) {
	for {
		i -= 1
		select {
		case <-ctx.Done():
			fmt.Printf("sub is cancelled ! num is : %d \n", i)
			time.Sleep(2 * time.Second)
			return
		default:
			{
				fmt.Printf("sub is processing ! num is : %d \n", i)
				time.Sleep(2 * time.Second)
			}
		}
	}
}
