package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type WorldTimeAPIResponse struct {
	Datetime string `json:"datetime"`
}

var (
	timeRequests int
	mutex        sync.Mutex
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	timeRequests++
	mutex.Unlock()

	resp, err := http.Get("http://worldtimeapi.org/api/timezone/Europe/Moscow")
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при запросе времени: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Не удалось получить время: status code %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	var apiResponse WorldTimeAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании ответа: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(apiResponse.Datetime))
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
