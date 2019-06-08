/*
*
*		Example of channel write with timeout without using additional
*		goroutines. Take a look at method 'channelWriter' :)
*
 */

package main

import (
	"context"
	"log"
	"time"
)

type event struct{}

func main() {
	var (
		timeoutToWrite = 2 * time.Second // Time for writing
		readDelay      = 3 * time.Second // Delay before event will be read by reader
	)

	if timeoutToWrite > readDelay {
		log.Println("expected to successfully read event :)")
	} else if timeoutToWrite < readDelay {
		log.Println("expected to lose event :(")
	} else {
		log.Println("result is unknown :|")
	}

	//
	// Push event and try to read it
	//
	var (
		eventStream = make(chan event)
		newEvent    event
		ctx         = context.Background()
	)

	ctx, _ = context.WithTimeout(ctx, timeoutToWrite)

	go channelReader(readDelay, eventStream)
	channelWriter(ctx, newEvent, eventStream)
}

func channelWriter(ctx context.Context, event event, stream chan event) {
	select {

	case stream <- event:
		log.Println("channelWriter: event is consumed by reader")

	case <-ctx.Done():
		log.Println("channelWriter: event write is canceled")

	}
}

func channelReader(readDelay time.Duration, eventStream chan event) {
	for {
		time.Sleep(readDelay)
		<-eventStream
		log.Println("channelReader: stream event is read")
	}
}
