package main

import (
    "net/http"
    "os"
    "fmt"

    "github.com/joho/godotenv"

    "molinahuel/qgis/back/py"
    "molinahuel/qgis/back/mobile"
    "molinahuel/qgis/back/front"
)


func main(){
    err := godotenv.Load()
    if err != nil {
        fmt.Println("error loading environment")
    }
    mux := http.NewServeMux()
    mux.HandleFunc("wsmobile/", mobile.MobileSocket)
    mux.HandleFunc("wspyclient/", py.PySocket)


    mux.HandleFunc("home/", frontend.Home)
    
    port := ":" + os.Getenv("PORT")
    host := os.Getenv("HOST")

    fmt.Printf("listening on %s:%s", host, port)
    http.ListenAndServe(port, mux)
}

