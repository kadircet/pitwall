package udp_server

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestUdpServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	port := ":27002"
	receive_chan := Init(ctx, port)

	server_addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		t.Error(err)
	}
	connection, err := net.DialUDP("udp", nil, server_addr)
	if err != nil {
		t.Error(err)
	}

	expected := []byte("test")
	connection.Write(expected)
	connection.Close()

	got := <-receive_chan
	if !bytes.Equal(got, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
	cancel()
}
