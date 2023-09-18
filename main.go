package main

import (
	"fmt"
	"net/http"
	"time"
)

func sendKeepAliveRequest() {
	serverURL := "https://casefile-api.onrender.com/"

	// Send a GET request to the root route
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Printf("Error sending keep-alive request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Keep-alive request sent successfully.")
	} else {
		fmt.Printf("Error sending keep-alive request. Status code: %d\n", resp.StatusCode)
	}
}

func main() {

	keepAliveInterval := 4 * time.Hour
	sendKeepAliveRequest()

	// Schedule the next keep-alive request every 4 hours
	ticker := time.NewTicker(keepAliveInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sendKeepAliveRequest()
		}
	}
}
