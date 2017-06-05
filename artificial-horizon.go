package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("resources"))))

	http.Handle("/echo", websocket.Handler(func(ws *websocket.Conn) {
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg[:n])

		m, err := ws.Write(msg[:n])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", msg[:m])

	}))
	http.ListenAndServe("127.0.0.1:42424", nil)
}
