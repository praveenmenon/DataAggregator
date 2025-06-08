package supervisor

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func Watch(cancelFunc context.CancelFunc) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Type 'stop' to terminate: ")
		text, _ := reader.ReadString('\n')
		if text == "stop\n" {
			fmt.Println("Shutting down...")
			cancelFunc()
			return
		}
	}
}
