package discovery

import (
	"fmt"
	"net"
	"time"

	"github.com/pinfake/pes6go/network/blocks"
	"github.com/pinfake/pes6go/network/messages"
	"github.com/pinfake/pes6go/network/messages/discovery"
	"github.com/pinfake/pes6go/server"
)

var handleMap map[uint16]func(block blocks.Block) messages.Message = map[uint16]func(block blocks.Block) messages.Message{
	0x2008: HandleDiscoveryInit,
	0x2006: HandleServerTime,
}

type Server struct {
	server.Handler
}

type VoidMessage struct {
	messages.Message
}

func (m VoidMessage) getBlocks() []blocks.Block {
	return []blocks.Block{}
}

func HandleDiscoveryInit(_ blocks.Block) messages.Message {
	fmt.Println("I am handling discovery init")
	return discovery.Init{
		Time:  time.Date(2017, 1, 1, 12, 0, 0, 0, time.UTC),
		Title: "Hey, this is a title!",
		Text:  "Hey, this is the text, not so long!",
	}
}

func HandleServerTime(_ blocks.Block) messages.Message {
	fmt.Println("I am handling server time")
	return discovery.ServerTime{
		Time: time.Now(),
	}
}

func (s Server) HandleBlock(block blocks.Block) (messages.Message, error) {
	method, ok := handleMap[block.Header.Query]
	if !ok {
		return nil, fmt.Errorf("Unknown query!")
	}
	return method(block), nil
}

func (s Server) HandleConnection(conn net.Conn) {
	for i := 1; i < 6; i++ {
		conn.Write([]byte(fmt.Sprintf("%d\n", i)))
		time.Sleep(1 * time.Second)
	}
}

func Start() {
	fmt.Println("Here i am the s server!")
	s := Server{}
	server.Serve(10881, s)
}
