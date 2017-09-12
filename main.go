package main

import (
  "fmt"
  "net/http"
  "log"

  "github.com/gorilla/websocket"
  "time"
  "math/rand"
  "strconv"
)

type PipelineStatus struct {
  Id      string
  Name    string
  Success bool
  Order   int
}

type PipelineMessage struct {
  NewBuildBroken   bool
  NewBuildFixed    bool
  PipelineStatuses []PipelineStatus
}

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

  pipelineMessageChan := make(chan PipelineMessage)
  go serveSocket(conn, pipelineMessageChan)
  go listenToConcourse(pipelineMessageChan)
}

func serveSocket(conn *websocket.Conn, incomingChan chan PipelineMessage) {
  for message := range incomingChan {
    if err := conn.WriteJSON(message); err != nil {
      fmt.Println(err)
    }
  }
}

func listenToConcourse(out chan PipelineMessage) {
  statuses := make(map[string]PipelineStatus)
  statuses["1"] = PipelineStatus{Id: "1", Name: "Pipeline 1", Success: true,}
  statuses["5"] = PipelineStatus{Id: "5", Name: "Pipeline 5", Success: true,}
  statuses["2"] = PipelineStatus{Id: "2", Name: "Pipeline 2", Success: false,}
  statuses["3"] = PipelineStatus{Id: "3", Name: "Pipeline 3", Success: true,}
  statuses["4"] = PipelineStatus{Id: "4", Name: "Pipeline 4", Success: true,}

  for {
    randId := strconv.Itoa(rand.Intn(5) + 1)
    success := rand.Intn(100)%4 == 0
    status := statuses[randId]
    status.Success = success
    statuses[randId] = status
    out <- PipelineMessage{
      PipelineStatuses: getStatusList(statuses),
    }
    time.Sleep(5 * time.Second)
  }
}

func getStatusList(statusMap map[string]PipelineStatus) []PipelineStatus {
  statuses := make([]PipelineStatus, 0)
  for _, val := range statusMap {
    statuses = append(statuses, val)
  }
  return statuses
}
