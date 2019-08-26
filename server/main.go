package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type ChatMessage struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		name := so.Request().FormValue("name")
		if name == "" {
			log.Println("disconnect for empty name", so.Id())
			so.Disconnect()
			return
		}
		log.Println("on connection", so.Id(), name)
		so.Join("room")
		log.Println("join room", so.Id())

		so.On("chat", func(msg string) {
			log.Println("chat", so.Id(), name, msg)
			chatMessage := ChatMessage{
				ID:      so.Id(),
				Name:    name,
				Message: msg,
			}
			server.BroadcastTo("room", "chat", chatMessage)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect", so.Id())
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
