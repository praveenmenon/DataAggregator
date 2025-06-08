package temperature

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Start(ctx context.Context, out chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			temp := 20 + rand.Intn(10)
			out <- fmt.Sprintf("Temperature: %d°C", temp)
			time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)
		}
	}
}
