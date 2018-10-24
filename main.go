package main

import (
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, JX!"))
}

func main(){
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil{
        panic(err)
    }
}
