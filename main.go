package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/suraboy/go-example/serivices/customers"
	"log"
	"net/http"
	"strings"
	time "time"
)

func main() {

	r := mux.NewRouter()
	customersService    = customers.NewService()
	customersHandler     = customers.NewHandler(customersService)
	r.HandleFunc("/api/customers", customersHandler.GetCustomers)
	r.HandleFunc("/api/time", TimeHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:7777",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

type Response struct {
	Name        string `json:"name,omitempty"`
	CurrentTime string `json:"current_time,omitempty"`
}

func TimeHandler(writer http.ResponseWriter, request *http.Request) {
	var timezones = strings.Split(request.URL.Query().Get("tz"), ",")
	if len(timezones) <= 1 {
		var loc, _ = time.LoadLocation(request.URL.Query().Get("tz"))
		parsedData := &Response{
			CurrentTime: time.Now().In(loc).String(),
		}
		responseData, _ := json.Marshal(parsedData)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseData)
	} else {
		parsedData := make(map[string]string)
		for _, v := range timezones {
			var loc, _ = time.LoadLocation(v)
			parsedData[v] = time.Now().In(loc).String()
		}
		responseData, _ := json.Marshal(parsedData)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseData)
	}
}
