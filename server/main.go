package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection", so.Id())
		so.Join("room")
		log.Println("join room", so.Id())
		so.On("chat", func(msg string) {
			log.Println("chat", so.Id(), msg)
			server.BroadcastTo("room", "chat", msg)
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
