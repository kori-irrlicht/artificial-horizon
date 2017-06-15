// +build !js

package network

import "net"

// Dial opens a udp, websocket or channel connection depending on the execution
// environment
// origin can be ignored on desktop
//
// Connects to 'url + /udp', so only the base address is needed
func Dial(url, origin string) (net.Conn, error) {
	return net.Dial("udp", url+"/udp")

}
