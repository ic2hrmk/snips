package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/ic2hrmk/snips/golang/access-trouble-shooter/pinger"
	"github.com/ic2hrmk/snips/golang/access-trouble-shooter/server"
)

var (
	port        string
	pingLogFile string
	pingHost    string
	pingDelay   time.Duration
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port to host server on")
	flag.StringVar(&pingLogFile, "log", "ping.log", "File to write ping errors")
	flag.StringVar(&pingHost, "host", "http://google.com", "Remote site to 'ping'")
	flag.DurationVar(&pingDelay, "delay", time.Minute, "Delay for PING on remote resource")
	flag.Parse()
}

func main() {
	log.Println("> HOST PORT:", port)
	log.Println("> LOG FILE: ", pingLogFile)
	log.Println("> PING HOST:", pingHost)
	log.Println("> DELAY:    ", pingDelay)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		troubleShootServer := server.NewTroubleShootServer()
		if err := troubleShootServer.Run(port); err != nil {
			log.Println("ERROR", err.Error())
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		remotePinger := pinger.NewExternalResourcePinger()
		remotePinger.LoopHttpPing(pingHost, pingLogFile, pingDelay)
		wg.Done()
	}()

	wg.Wait()
}
