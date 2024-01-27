package main

import (
	"Toshkent_Metro_System/Handlers"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("server started at: ", time.Now())

	http.HandleFunc("/", Handlers.Handle)
	http.ListenAndServe(":8080", nil)
}
