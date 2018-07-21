package models

import (
	"fmt"
	"genesis/utils"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketMessage WebsocketMessage
type WebsocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

var (
	clients    sync.Map
	broadcast  = make(chan WebsocketMessage)
	upgrader   = websocket.Upgrader{}
	pingPeriod = 10 * time.Second
	writeWait  = 5 * time.Second
)

func init() {
	go handleMessages()
}

func handleMessages() {
	for {
		msg := <-broadcast
		clients.Range(func(key, value interface{}) bool {
			client := key.(*websocket.Conn)
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				clients.Delete(client)
			}
			return true
		})
	}
}

// Upgrade Upgrade
func Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (err error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		fmt.Println(err)
	}
	clients.Store(ws, "ws")
	go receiveMessage(ws)
	go sendPingMessage(ws)
	ws.SetPongHandler(receivePongMessage)
	return
}

func receiveMessage(ws *websocket.Conn) {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			clients.Delete(ws)
			break
		}
		if messageType == websocket.PongMessage {
			log.Println("recv:", message)
		}
		log.Printf("recv: %s", message)
	}
}

func sendPingMessage(ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		log.Println("client conn close")
	}()
	for {
		select {
		case <-ticker.C:
			log.Println("send ping msg")
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func sendMessage(message string) {
	for {
		time.Sleep(time.Second * 3)
		msg := WebsocketMessage{Type: message, Data: message + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
	}
}

// SendWebsocketMessage SendWebsocketMessage
func SendWebsocketMessage(_type string, data interface{}) {
	utils.Println(data)
	msg := WebsocketMessage{Type: _type, Data: data}
	broadcast <- msg
}

func receivePongMessage(appData string) (err error) {
	log.Println("recv pong ", appData)
	return
}
