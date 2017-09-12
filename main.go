package main

import (
  "fmt"
  "net/http"
  "log"

  "github.com/gorilla/websocket"
)

func main() {
  http.HandleFunc("/ws", wsHandler)
  http.Handle("/", http.FileServer(http.Dir("dist")))
  log.Fatal(http.ListenAndServe(":8000", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
  if err != nil {
    http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
  }

  go echo(conn)
}

func echo(conn *websocket.Conn) {
  for {
    if err := conn.WriteJSON("Hello world"); err != nil {
      fmt.Println(err)
    }
  }
}
