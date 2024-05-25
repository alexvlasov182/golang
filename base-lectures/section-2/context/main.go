package main

import (
	"context"
	"fmt"
	"time"
)

// 1. context.Background() - on the high level
// 2. context.TODO() - when you are not sure what context to use
// 3. context.Value - you should use it rearly and put only unnesseary parameters
// 4. ctx - put in the function only as first element

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	ctx = context.WithValue(ctx, "id", 1)

	parse(ctx)

}

func parse(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Println(id)
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parseing completed")
			return
		case <-ctx.Done():
			fmt.Println("dedline exceded")
			return
		}
	}
}
