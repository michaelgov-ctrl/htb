package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)

  log.Println("Listening on :8000...")
  if err := http.ListenAndServe(":8000", nil); err != nil {
    log.Fatal(err)
  }
}
