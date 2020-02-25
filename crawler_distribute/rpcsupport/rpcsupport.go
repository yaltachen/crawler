package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// ServeRPC 开启一个rpc service
func ServeRPC(host string, service interface{}) error {
	var (
		listener net.Listener
		conn     net.Conn
		err      error
	)
	if err = rpc.Register(service); err != nil {
		log.Printf("register failed, err: %v\r\n", err)
		return err
	}

	if listener, err = net.Listen("tcp", host); err != nil {
		log.Printf("listen failed, err: %v\r\n", err)
		return err
	}
	for {
		if conn, err = listener.Accept(); err != nil {
			log.Printf("accept failed, err: %v\r\n", err)
			continue
		}
		log.Printf("accept conn")
		go jsonrpc.ServeConn(conn)
	}
}

// NewSupport 返回一个jsonrpc客户端
func NewSupport(host string) (*rpc.Client, error) {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", host); err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
