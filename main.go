package main

import (
  "fmt"
  "net/http"
  "log"
)

func main()  {
	fmt.Println("Hello world")
  http.Handle("/", http.FileServer(http.Dir("dist")))
  log.Fatal(http.ListenAndServe(":8000", nil))
}
