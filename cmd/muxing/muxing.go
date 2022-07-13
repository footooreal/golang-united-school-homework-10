package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	r := mux.NewRouter()
	r.HandleFunc("/name/{PARAM}", nameHandler).Methods(http.MethodGet)
	r.HandleFunc("/bad", badHandler).Methods(http.MethodGet)
	r.HandleFunc("/data", dataHandler).Methods(http.MethodPost)
	r.HandleFunc("/headers", headersHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
		log.Fatal(err)
	}
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	w.Write([]byte(fmt.Sprintf("Hello, %v!", v["PARAM"])))
	w.WriteHeader(http.StatusOK)
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	w.Write([]byte(fmt.Sprintf("I got message:%v", b)))
	w.WriteHeader(http.StatusOK)
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Add("a+b", strconv.Itoa(a+b))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
