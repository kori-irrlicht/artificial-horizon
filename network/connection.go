package network

import "net"

type Connection interface {
	Tcp() net.Conn
	Udp() net.Conn
}

type tcpUdpConnection struct {
	tcp net.Conn
	udp net.Conn
}

func NewConnection(address, tcpPort, udpPort string) (c Connection, err error) {
	conn := &tcpUdpConnection{}
	if conn.tcp, err = net.Dial("tcp", address+":"+tcpPort); err != nil {
		return nil, err
	}
	if conn.udp, err = net.Dial("udp", address+":"+udpPort); err != nil {
		return nil, err
	}

	return conn, err

}

func (c tcpUdpConnection) Tcp() net.Conn {
	return c.tcp
}
func (c tcpUdpConnection) Udp() net.Conn {
	return c.udp
}
