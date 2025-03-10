package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin (for testing).  For production, be
		// more restrictive.
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		log.Printf("Received: %s", string(p))

		if err := conn.WriteMessage(messageType, []byte("Server received: "+string(p))); err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	port := ":8080" // Or use an environment variable like : " + os.Getenv("PORT")
	fmt.Printf("WebSocket server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
