package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader    = websocket.Upgrader{}
	connections = make(map[*websocket.Conn]struct{})
)

func main() {
	http.HandleFunc("/push-list", pushListHandler)
	http.HandleFunc("/reject", rejectHandler)

	go broadcastLoop()

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func pushListHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	connections[conn] = struct{}{}
}

func rejectHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			break
		}

		// 处理驳回逻辑
		// ...

		// 发送响应给客户端
		err = conn.WriteMessage(websocket.TextMessage, []byte("Rejected"))
		if err != nil {
			log.Println("WriteMessage error:", err)
			break
		}
	}
}

type Data struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func broadcastLoop() {
	var data []*Data
	for i := 0; i < 1000; i++ {
		content := strconv.Itoa(i)
		data = append(data, &Data{
			Id:      int(time.Now().UnixNano()) + i,
			Content: "conte1" + content,
		})
	}
	// data := &Data{Id: 123, Content: "contente111"}
	listData, _ := json.Marshal(data)
	for {
		// for _, v := range data {

		for conn := range connections {
			err := conn.WriteMessage(websocket.TextMessage, listData)
			if err != nil {
				log.Println("WriteMessage error:", err)
				delete(connections, conn)
			}
			// }
		}
	}
}
