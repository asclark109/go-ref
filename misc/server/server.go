package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"proj1/feed"
)

type Config struct {
	Encoder *json.Encoder // Represents the buffer to encode Responses
	Decoder *json.Decoder // Represents the buffer to decode Requests
	Mode    string        // Represents whether the server should execute
	// sequentially or in parallel
	// If Mode == "s"  then run the sequential version
	// If Mode == "p"  then run the parallel version
	// These are the only values for Version
	ConsumersCount int // Represents the number of consumers to spawn
}

// Run starts up the twitter server based on the configuration
// information provided and only returns when the server is fully
// shutdown.
func Run(config Config) {

	// this main goroutine Run() acts as the producer for the task queue
	decoder := config.Decoder
	encoder := config.Encoder
	feed := feed.NewFeed()

	if config.Mode == PARALLEL_FLAG {
		// (1) spawn a certain number of goroutines (config.ConsumersCount)
		// each will be executing func consumer(...).
		//
		fmt.Println("not yet implemented")
		return
	}

	// SEQUENTIAL MODE
	if config.Mode == SEQUENTIAL_FLAG || config.Mode == "" {
		for {
			var request Request
			if err := decoder.Decode(&request); err == io.EOF {
				// fmt.Println("end of file")
				break
			} else if err != nil {
				log.Fatal(err)
				return
			}

			cmd := identifyCmd(&request)
			switch {
			case cmd == ADD_CMD:
				handleAddRequest(encoder, &request, &feed)
			case cmd == REMOVE_CMD:
				handleRemoveRequest(encoder, &request, &feed)
			case cmd == CONTAINS_CMD:
				handleContainsRequest(encoder, &request, &feed)
			case cmd == FEED_CMD:
				handleFeedRequest(encoder, &request, &feed)
			case cmd == DONE_CMD:
				return
			}

		}
	}

	// for {
	// check request queue (Decoder).
	// if non-empty, pull off request, handle it.
	// generate response, enqueue it onto response queue (encoder)
	// repeat until done request received

	// request_queue.
	// }

}

// string mappings (enums)
const (
	PARALLEL_FLAG   string = "p"
	SEQUENTIAL_FLAG string = "s"
	ADD_CMD         string = "ADD"
	REMOVE_CMD      string = "REMOVE"
	CONTAINS_CMD    string = "CONTAINS"
	FEED_CMD        string = "FEED"
	DONE_CMD        string = "DONE"
	UNKNOWN_CMD     string = ""
)

// // Request struct represents a schema with which to interpret incoming requests
// type Request struct {
// 	Command   string
// 	Id        int
// 	Body      string
// 	Timestamp int
// }

func identifyCmd(request *Request) string {
	cmd := request.Command
	switch {
	case cmd == ADD_CMD:
		return ADD_CMD
	case cmd == REMOVE_CMD:
		return REMOVE_CMD
	case cmd == CONTAINS_CMD:
		return CONTAINS_CMD
	case cmd == FEED_CMD:
		return FEED_CMD
	case cmd == DONE_CMD:
		return DONE_CMD
	}
	return UNKNOWN_CMD
}

// func handleCmd(request *Request) {
// 	cmd := identifyCmd(request)
// 	if cmd == UNKNOWN_CMD {
// 		fmt.Println("ERR: did not recognize command")
// 		return
// 	}

// 	switch {
// 	case cmd == ADD_CMD:
// 		handleAddRequest(request)
// 		// case cmd == REMOVE_CMD:
// 		// 	return REMOVE_CMD
// 		// case cmd == CONTAINS_CMD:
// 		// 	return CONTAINS_CMD
// 		// case cmd == FEED_CMD:
// 		// 	return FEED_CMD
// 		// case cmd == DONE_CMD:
// 		// 	return DONE_CMD
// 	}
// }

func handleAddRequest(encoder *json.Encoder, request *Request, userFeed *feed.Feed) {
	if identifyCmd(request) != ADD_CMD {
		fmt.Println("ERR: incorrectly called this function!")
	}

	(*userFeed).Add(request.Body, float64(request.Timestamp))

	response := map[string]interface{}{}
	response["success"] = true
	response["id"] = request.Id

	if err := encoder.Encode(&response); err != nil {
		fmt.Println("ERR: had trouble generating response:")
		fmt.Println(err)
		return
	}
}

func handleRemoveRequest(encoder *json.Encoder, request *Request, userFeed *feed.Feed) {
	if identifyCmd(request) != REMOVE_CMD {
		fmt.Println("ERR: incorrectly called this function!")
	}

	deleted := (*userFeed).Remove(request.Timestamp)

	response := map[string]interface{}{}
	response["success"] = deleted
	response["id"] = request.Id

	if err := encoder.Encode(&response); err != nil {
		fmt.Println("ERR: had trouble generating response:")
		fmt.Println(err)
		return
	}
}

func handleContainsRequest(encoder *json.Encoder, request *Request, userFeed *feed.Feed) {
	if identifyCmd(request) != CONTAINS_CMD {
		fmt.Println("ERR: incorrectly called this function!")
	}

	containsTimestamp := (*userFeed).Contains(request.Timestamp)

	response := map[string]interface{}{}
	response["success"] = containsTimestamp
	response["id"] = request.Id

	if err := encoder.Encode(&response); err != nil {
		fmt.Println("ERR: had trouble generating response:")
		fmt.Println(err)
		return
	}
}

func handleFeedRequest(encoder *json.Encoder, request *Request, userFeed *feed.Feed) {
	if identifyCmd(request) != FEED_CMD {
		fmt.Println("ERR: incorrectly called this function!")
	}

	var contents *map[float64]string = (*userFeed).ContentsAsMap()
	var contentsSlice []interface{} = make([]interface{}, len(*contents))

	response := map[string]interface{}{}
	// response["success"] = containsTimestamp
	response["id"] = request.Id
	response["feed"] = contentsSlice

	idx := 0
	for timestamp, body := range *contents {
		contentsSlice[idx] = map[string]interface{}{"body": body, "timestamp": timestamp}
		idx++
	}

	if err := encoder.Encode(&response); err != nil {
		fmt.Println("ERR: had trouble generating response:")
		fmt.Println(err)
		return
	}
}

// func consumer(haredContext *SharedContext) {
// 	// each consumer goroutine will try to grab one task from queue
// 	// consumer will process rqequest and send appropriate response back.
// 	// when consumer finishes executing task, it checks the queue to grab another task.
// 	// if there are no tasks in queue then it will need to wait for more tasks to process or exits function if there
// 	// are no more remaining tasks to complete.
// }

// func producer(sharedContext *SharedContext) {
// 	// this is the main go routine
// 	// reads in from json.Decorder.
// 	// place tasks inside a queue data structure and do:
// 	// if there is a consumer waiting for work then place a task in side queue and wake up
// 	// else just place task on queue.
// 	// repeat

// 	// if receiving a done command, just issue. let a consumer thread handle the end

// }

type SharedContext struct {
	Encoder *json.Encoder // Represents the buffer to encode Responses
}

type Request struct {
	Command   string
	Id        int
	Body      string
	Timestamp float64
}

type AddResponse struct {
	success bool
	id      int
}
