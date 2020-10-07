package udp_server

import (
	"context"
	"log"
	"net"
)

func listen(ctx context.Context, receive chan []byte, connection *net.UDPConn) {
	defer connection.Close()
	defer close(receive)

	buffer := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		n, err := connection.Read(buffer)
		if err != nil {
			log.Fatal(err)
			return
		}
		receive <- buffer[:n]
	}
}

// Initializes a udp server listening at localhost:port.
// FIXME: Report F12020 packetheaders rather than []byte.
func Init(ctx context.Context, port string) <-chan []byte {
	address, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	connection, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	receive := make(chan []byte, 10)
	go listen(ctx, receive, connection)
	return receive
}
