package main

import (
	"fmt"
	"github.com/praveenmenon/DataAggregator/actor1"
	"github.com/praveenmenon/DataAggregator/actor2"
	"github.com/praveenmenon/DataAggregator/actor3"
	"github.com/praveenmenon/DataAggregator/models"
	"log"
	"math/rand"
	"os"
	"time"
)

//It takes teh query string and any number of searches
func FirstGoroutine(query string, replicas ...models.Search) models.Result {
	//create a channels of results as bug as our replica
	c := make(chan models.Result, len(replicas))

	// A small closure which take i and invokes query on search i -> replicas[i]
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	// Fires goroutines one for each replica
	for i := range replicas {
		go searchReplica(i)
	}

	//return the first result
	return <- c
}

func SearchEngine(query string)(results []models.Result){
	// Created a channel which receives 3 results
	c := make(chan models.Result, 3)
	go func() {c <- FirstGoroutine(query, actor1.Web1, actor1.Web2, actor1.Web3)}()
	go func() {c <- FirstGoroutine(query, actor2.Image1, actor2.Image2, actor2.Image3)}()
	go func() {c <- FirstGoroutine(query, actor3.Video1, actor3.Video2, actor3.Video3)}()

	// I kept a timer of 80 milliseconds. I am returning a timeout channel after 80 milliseconds
	// the timeout condition will be executed and since only one case at a time is considered,
	// the result will not be appended and stop sill be printed on the terminal and return immediately.
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i< cap(c); i++ {
		select {
		case result := <-c:
				results = append(results, result)
		case <- timeout:
				fmt.Println("stop")
				return
		}
	}
	return
}

func main() {
	//run the aggregator function 50 times
	const n = 50
	f, err := os.OpenFile("output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for i :=0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		start := time.Now()
		// Trying to replicate a google SEO data aggregator with "Tesla" as the search key
		results := SearchEngine("Tesla")
		elapsed := time.Since(start)
		for _, result := range results{
			//prints the data from each actor to the output.log file
			if _, err := f.WriteString(string(result)); err != nil {
				log.Println(err)
			}
		}
		//prints the total time elapsed to get data from the actors
		if _, err := f.WriteString( "===================================== Total time elapsed: " + elapsed.String() +" =====================================\n\n"); err != nil {
			log.Println(err)
		}
	}
}
