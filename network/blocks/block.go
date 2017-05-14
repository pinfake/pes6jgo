package blocks

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/pinfake/pes6go/network"
)

const headerSize = 24

type Header struct {
	Query    uint16
	Size     uint16
	Unknown1 uint16
	Sequence uint16
	Unknown2 [16]byte
}

type Body interface {
	GetBytes() []byte
}

type Block struct {
	Header Header
	body   Body
}

type GenericBody struct {
	data []byte
}

func (body GenericBody) GetBytes() []byte {
	return body.data
}

func NewHeader(query uint16, size uint16) Header {
	return Header{Query: query, Size: size}
}

func NewBlock(query uint16, body Body) Block {
	return Block{NewHeader(query, uint16(len(body.GetBytes()))), body}
}

func (b Block) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, b.Header)
	buf.Write(b.body.GetBytes())
	return buf.Bytes()
	//retb := [headerSize + b.Header.Size]byte{}
	//var buf = bytes.NewBuffer(retb[:])
	//binary.Write(buf, binary.BigEndian, &b.Header)
	//buf.Write(b.body.GetBytes())
	//return retb[:]
}

func ReadBlock(data []byte) (Block, error) {
	if len(data) < headerSize {
		return Block{}, errors.New("No Header found")
	}
	decoded := network.Mutate(data)
	var buf = bytes.NewBuffer(decoded[0:headerSize])
	header := Header{}
	err := binary.Read(buf, binary.BigEndian, &header)
	if err != nil {
		panic(err)
	}

	if len(decoded) < int(headerSize+header.Size) {
		fmt.Printf("%d headersize % x", header.Size, decoded)
		panic("hostion")
	}
	return Block{header, GenericBody{decoded[headerSize : headerSize+header.Size]}}, nil
}
