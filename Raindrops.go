package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// RainDrops checks if passed argument is divisible by 3
func RainDrops(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.ParseInt(r.URL.Path[1:], 10, 64)
	if err != nil {
		number = 0
	}
	rainDropResult := ""
	if number%3 == 0 {
		rainDropResult += "Pling"
	}
	if number%5 == 0 {
		rainDropResult += "Plang"
	}

	if number%7 == 0 {
		rainDropResult += "Plong"
	}

	fmt.Fprintf(w, rainDropResult)
}

func main() {
	http.HandleFunc("/", RainDrops)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
