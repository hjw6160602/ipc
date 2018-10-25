package main

import (
    "net/http"
    "fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handle hello")
    fmt.Fprintf(w, "Hello World! ")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handle login")
    fmt.Fprintf(w, "login ")
}

func main() {
    //http.HandleFunc("/", Hello)
    //http.HandleFunc("/login", login)
    //err := http.ListenAndServe(":8880", nil)
    //if err != nil {
    //    fmt.Println("http listen failed")
    //}

    http.HandleFunc("/", logPanics(Hello))
    http.HandleFunc("/login", logPanics(login))
    err := http.ListenAndServe(":8880", nil)
    if err != nil {
        fmt.Println("http listen failed")
    }

}


func logPanics(handle http.HandlerFunc) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        defer func() {
            if x := recover(); x != nil {
                fmt.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
            }
        }()
        handle(writer, request)
    }
}