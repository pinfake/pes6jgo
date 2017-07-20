package server

import (
	"fmt"
	"time"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type DiscoveryServer struct {
	logger      *log.Logger
	connections Connections
}

var discoveryHandlers = map[uint16]Handler{
	0x2008: DiscoveryInit,
	0x2006: ServerTime,
	0x2005: Servers,
	0x2200: RankUrls,
}

func NewDiscoveryServer() DiscoveryServer {
	return DiscoveryServer{
		logger:      log.New(os.Stdout, "Discovery: ", log.LstdFlags),
		connections: NewConnections(),
	}
}

func (s DiscoveryServer) Config() ServerConfig {
	return ServerConfig{}
}

func (s DiscoveryServer) Connections() Connections {
	return s.connections
}

func (s DiscoveryServer) Handlers() map[uint16]Handler {
	return discoveryHandlers
}

func (s DiscoveryServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s DiscoveryServer) Logger() *log.Logger {
	return s.logger
}

func DiscoveryInit(s Server, _ block.Block, _ *Connection) message.Message {
	return message.NewServerNewsMessage(
		s.Storage().GetServerNews(),
	)
}

func Servers(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewServerListMessage(
		[]block.Server{
			{7, "GROUP-SP/", "127.0.0.1", 10887, 0},
			{6, "SCHE-SP/", "127.0.0.1", 10887, 0},
			{4, "QUICK0-SP/", "127.0.0.1", 10887, 0},
			{4, "QUICK1-SP/", "127.0.0.1", 10887, 0},
			{8, "MENU03-SP/", "127.0.0.1", 12882, 0},
			{3, "TurboLobas Inc.", "127.0.0.1", 10887, 50},
			{3, "TurboLobas Inc.", "127.0.0.1", 10888, 130},
			{2, "ACCT03-SP/", "127.0.0.1", 12881, 0},
			{1, "GATE-SP/", "127.0.0.1", 10887, 0},
		},
	)
}

func RankUrls(s Server, _ block.Block, _ *Connection) message.Message {
	return message.NewRankUrlListMessage(
		s.Storage().GetRankUrls(),
	)
}

func ServerTime(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.ServerTime{
		ServerTime: block.ServerTime{Time: time.Now()},
	}
}

func StartDiscovery() {
	fmt.Println("Discovery Server starting")
	Serve(NewDiscoveryServer(), 10881)
}