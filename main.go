package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
)

func main() {
    http.HandleFunc("/ping", PingServer)
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func PingServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "frontend: Pong")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://backend:8080")
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Fprintf(w, "backend: %s", body)
}
