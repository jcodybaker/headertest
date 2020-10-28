package main

import (
  "log"
  "net/http"
)

func main() {
  s := &http.Server{
    Addr: ":8080",
    Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Header().Add("X-Total-Count", "123")
      w.WriteHeader(http.StatusNoContent)
    }),
  }
  log.Fatal(s.ListenAndServe())
}
