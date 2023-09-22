package main

import (
	"fmt"
	"net/http"
	"time"
)

func sendKeepAliveRequest() {
	
	serverURL := "https://casefile-api.onrender.com/" 

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

	// Schedule the next keep-alive request every 2 minutes
	ticker := time.NewTicker(keepAliveInterval)
	defer ticker.Stop()

	// Start an HTTP server on port 3000
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is your HTTP server response.")
		})
		fmt.Println("HTTP server listening on :3000")
		http.ListenAndServe(":3000", nil)
	}()

	for {
		select {
		case <-ticker.C:
			sendKeepAliveRequest()
		}
	}
}
