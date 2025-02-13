package main

import (
	"encoding/json"
	"fmt"
	"os"
	"proj1/server"
	"proj1/utils"
	"strconv"
)

func main() {

	// (1) interpret args. terminate early if bad args are supplied.
	// (2) otherwise generate config and start server

	// CHECK: 1 arg provided
	arguments := os.Args[1:]
	if !(len(arguments) == 1) {
		fmt.Println(USAGE)
		return
	}

	// CHECK: arg is int
	numConsumersStr := arguments[0]
	if !(utils.IsPostiveInt(numConsumersStr)) {
		fmt.Println(USAGE)
		return
	}
	numConsumers, _ := strconv.Atoi(numConsumersStr)

	// run program (generate config and launch server)
	encoder := json.NewEncoder(os.Stdout)
	decoder := json.NewDecoder(os.Stdin)
	mode := "s"
	config := server.Config{encoder, decoder, mode, numConsumers}
	server.Run(config)
}

const (
	USAGE string = `Usage: twitter <number of consumers>
    <number of consumers> = the number of goroutines (i.e., consumers) to be part of the parallel version.`
)
