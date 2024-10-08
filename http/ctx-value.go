package main

import (
	"context"
	"fmt"
)


func main(){
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "jafar")
	
	f1(ctx)
	f2(ctx)
}

func f1(ctx context.Context) {
	fmt.Printf("f1 .... %s \n", ctx.Value("name"))
}
func f2(ctx context.Context) {
	fmt.Printf("f2 .... %s \n", ctx.Value("name"))
}