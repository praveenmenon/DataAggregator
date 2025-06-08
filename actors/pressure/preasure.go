package pressure

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Start generates fake pressure data and sends it to the data channel until the context is cancelled.
func Start(ctx context.Context, out chan<- string) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Pressure] Stopping pressure sensor.")
			return
		case <-ticker.C:
			pressure := 950 + rand.Float64()*100 // Simulated pressure between 950–1050 hPa
			out <- fmt.Sprintf("pressure:%.2f", pressure)
		}
	}
}
