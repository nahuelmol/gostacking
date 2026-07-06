package main

import (
    "net/http"
    "os"
    "fmt"

    "github.com/joho/godotenv"
)

func main(){
    err := godotenv.Load()
    if err != nil {
        fmt.Println("error loading environment")
    }
    mux := http.NewServeMux()
    
    port := ":" + os.Getenv("PORT")
    host := os.Getenv("HOST")

    fmt.Printf("Listening on %s:%s", host, port)
    http.ListenAndServe(port, mux)
}

