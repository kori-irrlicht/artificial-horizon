// +build !js

package network

import (
	"net"
	"time"
)

/*
	// TODO: Implement later. currently not needed
*/

// channelConn is a channel-based connection
// It can only be used, if client and server are running in the same process
type channelConn struct{}

func (cc *channelConn) Close() error                     { return nil }
func (cc *channelConn) LocalAddr() net.Addr              { return nil }
func (cc *channelConn) RemoteAddr() net.Addr             { return nil }
func (cc *channelConn) Read([]byte) (int, error)         { return 0, nil }
func (cc *channelConn) Write([]byte) (int, error)        { return 0, nil }
func (cc *channelConn) SetDeadline(time.Time) error      { return nil }
func (cc *channelConn) SetReadDeadline(time.Time) error  { return nil }
func (cc *channelConn) SetWriteDeadline(time.Time) error { return nil }

// Check if channelConn implements net.Conn
var _ net.Conn = (*channelConn)(nil)

var newChannelConnections = make(chan *channelConn, 255)

// DialChannel creates a new channel-based connection
func DialChannel() (net.Conn, error) {
	return nil, nil
}

// ListenChannel listens, if a new channel-based connection has been created`
func ListenChannel() (net.Conn, error) {

	return nil, nil
}
