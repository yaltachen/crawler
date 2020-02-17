package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
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
		go rpc.ServeConn(conn)
	}
}
