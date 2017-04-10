package network;

import (
    "testing"
    "bytes"
)

var clean = []byte {
0x03, 0x03, 0x84, 0x02, 0x30, 0x01, 0x46, 0x02, 0x61, 0x72, 0x7a, 0x6e, 0xb6, 0xc3, 0x69, 0x73,
0x63, 0x73, 0x2f, 0x68, 0x43, 0x50, 0x30, 0x2f, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x5f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
};

var mutated = []byte {
0xa5, 0x74, 0x11, 0x7e, 0x96, 0x76, 0xd3, 0x7e, 0xc7, 0x05, 0xef, 0x12, 0x10, 0xb4, 0xfc, 0x0f,
0xc5, 0x04, 0xba, 0x14, 0xe5, 0x27, 0xa5, 0x53, 0xa6, 0x44, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x96, 0x7c, 0xa6, 0x28, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
}

func TestRead(t *testing.T) {
    t.Run("Should return an error on short buffer", func(t *testing.T) {
        b := [] byte{};
        _, err := Read(b);
        if err == nil {
            t.Error("No error on short byte array");
        }
    });
    t.Run("Should return mutated header", func(t *testing.T) {
        message, _ := Read(mutated);
        if message.header.Query != 0x0303 {
            t.Error( "Unexpected query!")
        }
    });
}

func TestMutate(t *testing.T) {
    mutation := Mutate(clean);
    if !bytes.Equal(mutation, mutated) {
        t.Error( "Bad mutation!");
    }
}
