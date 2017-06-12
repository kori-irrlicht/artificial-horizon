// +build js

package network

import (
	"net"

	"github.com/goxjs/websocket"
)

// Dial opens a tcp, websocket or channel connection depending on the execution
// environment
// origin can be ignored on desktop
func Dial(url, origin string) (net.Conn, error) {
	return websocket.Dial(url, origin)
}
