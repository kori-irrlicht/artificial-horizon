// +build js

package network

import (
	"net"

	"github.com/goxjs/websocket"
)

// Dial opens a udp, websocket or channel connection depending on the execution
// environment
// origin can be ignored on desktop
//
// Connects to 'ws:// + url + /ws', so only the base address is needed
func Dial(url, origin string) (net.Conn, error) {
	return websocket.Dial("ws://"+url+"/ws", origin)
}
