package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")

	if text == "" {
		http.Error(w, "misssing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number", http.StatusBadRequest)
		return
	}

	if _, err = fmt.Fprintln(w, v*2); err != nil {
		http.Error(w, "not a number", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/double", doubleHandler)
	log.Fatalln(http.ListenAndServe(":4000", nil))
}
