package pinger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ExternalResourcePinger struct{}

func NewExternalResourcePinger() *ExternalResourcePinger {
	return &ExternalResourcePinger{}
}

func (rcv *ExternalResourcePinger) LoopHttpPing(
	resourceAddress string,
	logFilePath string,
	delay time.Duration,
) {
	client := http.Client{}

	for {
		_, err := client.Get(resourceAddress)
		if err != nil {
			log.Printf("<- FAILED TO ACCESS RESOURCE [%s], %s\n",
				resourceAddress, err.Error())

			logMessage := fmt.Sprintf("[%s] FAILED TO ACCESS RESOURCE [%s], %s\n",
				time.Now().Format("2006-01-2 15:04:05"), resourceAddress, err.Error())

			if err := appendDataToFile(logFilePath, []byte(logMessage)); err != nil {
				log.Printf("X- UNABLE TO WRITE LOG TO [%s], %s", logFilePath, err.Error())
			}
		} else {
			log.Printf("V- REMOTE RESOURCE [%s] IS AVAILABLE", resourceAddress)
		}

		time.Sleep(delay)
	}
}

func appendDataToFile(path string, data []byte) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
