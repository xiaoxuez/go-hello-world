package context

import (
	"context"
	"fmt"
	"time"
)

func MainFunction() {
	ctx, cancel := context.WithCancel(context.Background())
	go A(ctx)
	go B(ctx)

	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(4 * time.Second)
}

func A(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("  A return ")
			return
		default:
			time.Sleep(3 * time.Second)
			fmt.Println(" A ")
		}
	}
}

func B(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("  B return ")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println(" B ")
		}
	}
}
