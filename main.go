package main

import (
 "github.com/gorilla/mux"
 "log"
 "net/http"
 "time"
 "io"
 "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, "Hello, world!\n")
}

func router() *mux.Router {
 r := mux.NewRouter()
 r.HandleFunc("/", handler)
 return r
}


func main() {
 fmt.Println("Hello, world!")
 router := router()
 srv := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:9100",
  WriteTimeout: 15 * time.Second,
  ReadTimeout: 15 * time.Second,
}

log.Fatal(srv.ListenAndServe())

}
