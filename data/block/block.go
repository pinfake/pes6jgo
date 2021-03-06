package block

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"crypto/md5"

	"github.com/pinfake/pes6go/crypt"
)

const dtLayout = "2006-01-15 15:04:05"
const headerSize = 24

type Header struct {
	Query    uint16
	Size     uint16
	Sequence uint32
	Hash     [16]byte // Sixserver says this is the md5 of Data
}

type Body interface {
	GetBytes() []byte
}

type Block struct {
	Header *Header
	Body   Body
}

type GenericBody struct {
	Data []byte
}

func (body GenericBody) GetBytes() []byte {
	return body.Data
}

func newHeader(query uint16, size uint16) *Header {
	return &Header{Query: query, Size: size}
}

func NewBlock(query uint16, body Body) *Block {
	return &Block{newHeader(
		query,
		uint16(len(body.GetBytes())),
	), body}
}

func (b *Block) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, b.Header)
	buf.Write(b.Body.GetBytes())
	return buf.Bytes()
}

func (b *Block) String() string {
	return fmt.Sprintf("%x %x", b.Header, b.Body)
}

func (b *Block) getHash() [16]byte {
	raw := make([]byte, headerSize-16+b.Header.Size)
	copy(raw[:headerSize-16], b.GetBytes()[:headerSize-16])
	copy(raw[headerSize-16:], b.GetBytes()[headerSize:])
	return md5.Sum(raw)
}

func (b *Block) Sign(sequence uint32) {
	b.Header.Sequence = sequence
	b.Header.Hash = b.getHash()
}

func ReadBlock(data []byte) (*Block, error) {
	if len(data) < headerSize {
		return nil, errors.New("No Header found")
	}
	decoded := crypt.ApplyMask(data)
	var buf = bytes.NewBuffer(decoded[0:headerSize])
	header := Header{}
	err := binary.Read(buf, binary.BigEndian, &header)

	if err != nil {
		return nil, fmt.Errorf("unable to read: %s", err)
	}

	if len(decoded) < int(headerSize+header.Size) {
		return nil, fmt.Errorf(
			"Smaller body than header said, header: %d, body: %d",
			header.Size, len(decoded),
		)
	}

	b := Block{
		&header,
		GenericBody{decoded[headerSize : headerSize+header.Size]},
	}

	if b.getHash() != header.Hash {
		return nil, fmt.Errorf(
			"invalid getHash, expected: %x got: %x",
			b.getHash(), header.Hash,
		)
	}

	return &b, nil
}
