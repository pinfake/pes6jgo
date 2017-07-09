package accounting

import (
	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
)

type AccountingServer struct {
}

var handlers = map[uint16]server.Handler{
	0x3001: Init,
	0x3003: Login,
	0x3010: Profiles,
	0x3060: QueryPlayerId,
	0x0005: KeepAlive,
	0x0003: Disconnect,
}

func (s AccountingServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func QueryPlayerId(_ block.Block, _ *server.Connection) message.Message {
	return message.NewPlayerIdResponseMessage(block.PlayerIdOk)
}

func Profiles(_ block.Block, _ *server.Connection) message.Message {

	return message.NewAccountPlayersMessage(
		block.AccountPlayers{
			Players: [3]block.AccountPlayer{
				{
					Position:      0,
					Id:            12345,
					Name:          "PadreJohn",
					TimePlayed:    1000,
					Division:      2,
					Points:        0,
					Category:      500,
					MatchesPlayed: 20,
				},
				{
					Position:      1,
					Id:            0,
					Name:          "",
					TimePlayed:    0,
					Division:      2,
					Points:        0,
					Category:      500,
					MatchesPlayed: 0,
				},
				{
					Position:      2,
					Id:            0,
					Name:          "",
					TimePlayed:    0,
					Division:      2,
					Points:        0,
					Category:      500,
					MatchesPlayed: 0,
				},
			},
		})
}

func Login(b block.Block, _ *server.Connection) message.Message {
	auth := block.NewAthentication(b)
	fmt.Println("I am handling login")
	fmt.Printf("key: % x\n", auth.Key)
	fmt.Printf("password: % x\n", auth.Password)
	fmt.Printf("unknown: % x\n", auth.Unknown)
	fmt.Printf("roster: % x\n", auth.RosterHash)
	return message.LoginResponse{}
}

func Init(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling init")
	return message.AccountingInit{}
}

func KeepAlive(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("I am handling a keep alive")
	return message.KeepAlive{}
}

func Disconnect(_ block.Block, _ *server.Connection) message.Message {
	fmt.Println("Handling disconnect")
	return nil
}

func Start() {
	fmt.Println("Here i am the accounting server!")
	server.Serve(AccountingServer{}, 12881)
}
