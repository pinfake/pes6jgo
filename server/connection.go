package server

import (
	"net"

	"sync"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/info"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/network"
)

type Connection struct {
	id        int
	conn      net.Conn
	seq       uint32
	AccountId uint32
	Player    *info.Player
}

func (c *Connection) readBlock() (*block.Block, error) {
	var data [4096]byte
	slice := data[:]

	n, err := c.conn.Read(slice)
	if err != nil {
		return nil, err
	}
	got, err := block.ReadBlock(slice[:n])
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (c *Connection) writeMessage(message message.Message) {
	for _, b := range message.GetBlocks() {
		c.seq++
		b.Header.Sequence = c.seq
		c.conn.Write(network.Mutate(b.GetBytes()))
	}
}

type Connections struct {
	mu          sync.RWMutex
	connections map[int]*Connection
}

func NewConnections() *Connections {
	return &Connections{
		connections: make(map[int]*Connection),
	}
}

func (conns *Connections) remove(id int) {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	conns.connections[id].conn.Close()
	delete(conns.connections, id)
}

func (conns *Connections) add(c net.Conn) *Connection {
	defer conns.mu.Unlock()
	conns.mu.Lock()
	connection := Connection{
		id:   conns.newId(),
		seq:  0,
		conn: c,
	}
	conns.connections[connection.id] = &connection
	return &connection
}

func (conns *Connections) newId() int {
	return len(conns.connections) + 1
}

func (conns *Connections) findByPlayerId(id uint32) *Connection {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}

func (conns *Connections) findByPlayerName(name string) *Connection {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}

func (conns *Connections) findInLobby(lobbyId byte) *Connections {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}

func (conns *Connections) findInRoom(lobbyId byte, roomId uint32) *Connections {
	defer conns.mu.RUnlock()
	conns.mu.RLock()
	return nil
}
