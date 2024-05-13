package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	timeRequests int
	mutex        sync.Mutex
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	timeRequests++
	mutex.Unlock()
	currentTime := time.Now().Format(time.RFC1123)
	w.Write([]byte(currentTime))
}

func statisticsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count := timeRequests
	mutex.Unlock()
	w.Write([]byte(fmt.Sprintf("%d", count)))
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/statistics", statisticsHandler)
	http.ListenAndServe(":8080", nil)
}
