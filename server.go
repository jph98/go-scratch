package main

import ("fmt"
    "net/http"
)

func main() {
  http.HandleFunc("/",
        func(w http.ResponseWriter, req *http.Request) {
            fmt.Fprintln(w, "Hello World!")
        })
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}