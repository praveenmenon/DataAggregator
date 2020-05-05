# Golang Data Aggregator

A simple data aggregator that subscribes to data published by all actors/asynchronous processes.
* There are 3 actors which are created as packages.
    * Web actor
    * Image actor
    * Video actor
* The models have required dataTypes for the data retrieved from each actor.

## Setup 
* Let's assume you created a program in `$GOPATH/src/example` directory:
* Copy the project to your directory under `src/github.com/preveenmenon`
* Navigate to your package's directory and use the go tool to build and run your code.
* The quickest way to run it is using `go run:`

    `$ go run aggregator.go`
* You can also build the project and run the executable file.
    ```
       $ go build
       $ ls
       $ ./DataAggregator
    ```
* The file output.log will contain the aggregated data form out actors.

