package listener

import (
	"net"
	"os"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func NewListener(net, laddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func newUnixListener(laddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func isSocket(m os.FileMode) bool { _ = "STUB: not implemented"; return false }

func newTCPListener(laddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
