package aggregator

import (
	"context"
	"fmt"
	"os"
	"time"
)

func Start(ctx context.Context, dataCh <-chan string) {
	file, err := os.Create("output.log")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(file, "Aggregator stopped.")
			return
		case msg := <-dataCh:
			timestamp := time.Now().Format(time.RFC3339)
			fmt.Fprintf(file, "[%s] %s\n", timestamp, msg)
		}
	}
}
