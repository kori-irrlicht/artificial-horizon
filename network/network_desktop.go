// +build !js

package network

import "net"

// Dial opens a tcp, websocket or channel connection depending on the execution
// environment
// origin can be ignored on desktop
func Dial(url, origin string) (net.Conn, error) {
	return net.Dial("tcp", url)

}
