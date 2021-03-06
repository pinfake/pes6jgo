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
	config  ServerConfig
	storage storage.Storage
}

var discoveryHandlers = map[uint16]Handler{
	0x2008: DiscoveryInit,
	0x2006: ServerTime,
	0x2005: Servers,
	0x2200: RankUrls,
}

func NewDiscoveryServerHandler(stor storage.Storage) DiscoveryServer {
	return DiscoveryServer{
		storage: stor,
		config:  ServerConfig{},
	}
}

func (s DiscoveryServer) Config() ServerConfig {
	return s.config
}

func (s DiscoveryServer) Handlers() map[uint16]Handler {
	return discoveryHandlers
}

func (s DiscoveryServer) Storage() storage.Storage {
	return s.storage
}

func (s DiscoveryServer) Data() interface{} {
	return nil
}

func DiscoveryInit(s *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewServerNews(
		s.Storage().GetServerNews(),
	)
}

func Servers(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewServerList(
		[]*block.Server{
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

func RankUrls(s *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewRankUrlList(
		s.Storage().GetRankUrls(),
	)
}

func ServerTime(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewServerTime(&block.ServerTime{Time: time.Now()})
}

func StartDiscovery(stor storage.Storage) {
	fmt.Println("Discovery Server starting")
	s := NewServer(
		log.New(os.Stdout, "Discovery: ", log.LstdFlags),
		NewDiscoveryServerHandler(stor),
	)
	s.Serve(10881)
}
