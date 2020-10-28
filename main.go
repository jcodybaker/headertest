package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  s := &http.Server{
    Addr: ":8080",
    Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Header().Add("X-Total-Count", "123")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(r.Method+"\n\n"))
      for k, vals := range r.Header {
        for _, v := range vals {
          fmt.Fprintf(w, "%s: %s\n", k, v)
        }
      }
      fmt.Println(r.Method)
    }),
  }
  log.Fatal(s.ListenAndServe())
}
