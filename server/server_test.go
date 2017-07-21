package server

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/storage"
)

var clean = []byte{
	0x84, 0x50, 0x04, 0x01, 0x00, 0x00, 0x86, 0x02, 0xe1, 0x5b, 0x1d, 0xaf, 0x4b, 0xc2, 0x39, 0x06,
	0x6b, 0x25, 0x17, 0xd1, 0xcd, 0xdc, 0x39, 0x05, 0x00, 0x00, 0x5b, 0x30, 0x64, 0x61, 0x6c, 0x65,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var mutated = []byte{
	0x22, 0x27, 0x91, 0x7d, 0xa6, 0x77, 0x13, 0x7e, 0x47, 0x2c, 0x88, 0xd3, 0xed, 0xb5, 0xac, 0x7a,
	0xcd, 0x52, 0x82, 0xad, 0x6b, 0xab, 0xac, 0x79, 0xa6, 0x77, 0xce, 0x4c, 0xc2, 0x16, 0xf9, 0x19,
	0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
	0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
}

type emptyServer struct {
	logger      *log.Logger
	connections Connections
}

func (s emptyServer) Logger() *log.Logger {
	return s.logger
}

func (s emptyServer) Connections() Connections {
	return s.connections
}

func (s emptyServer) Config() ServerConfig {
	return ServerConfig{}
}

func (s emptyServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s emptyServer) Handlers() map[uint16]Handler {
	return map[uint16]Handler{}
}

func NewEmptyServer() emptyServer {
	return emptyServer{
		logger:      log.New(ioutil.Discard, "empty: ", log.LstdFlags),
		connections: NewConnections(),
	}
}

func TestShouldConnect(t *testing.T) {
	s := NewEmptyServer()
	go Serve(s, 19770)
	t.Run("Should be able to connect", func(t *testing.T) {
		c := client.Client{}
		err := c.Connect("localhost", 19770)
		if err != nil {
			t.Error("Error connecting: %s", err.Error())
		}
	})
}

func TestSendInvalidData(t *testing.T) {
	s := NewEmptyServer()
	go Serve(s, 19770)
	t.Run("Should not crash on invalid data", func(t *testing.T) {

	})
}