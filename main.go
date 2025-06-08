package main

import (
	"context"
	"data-aggregator/actors/pressure"
	"data-aggregator/actors/temperature"
	"data-aggregator/aggregator"
	"data-aggregator/supervisor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dataCh := make(chan string)

	go temperature.Start(ctx, dataCh)
	go pressure.Start(ctx, dataCh)
	go aggregator.Start(ctx, dataCh)
	go supervisor.Watch(cancel)

	<-ctx.Done()
}
