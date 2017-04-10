package network;

import (
    "testing"
    "bytes"
)

var clean = []byte{
    0x84, 0x50, 0x04, 0x01, 0x00, 0x00, 0x86, 0x02, 0xe1, 0x5b, 0x1d, 0xaf, 0x4b, 0xc2, 0x39, 0x06,
    0x6b, 0x25, 0x17, 0xd1, 0xcd, 0xdc, 0x39, 0x05, 0x00, 0x00, 0x5b, 0x30, 0x64, 0x61, 0x6c, 0x65,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
};

var mutated = []byte{
    0x22, 0x27, 0x91, 0x7d, 0xa6, 0x77, 0x13, 0x7e, 0x47, 0x2c, 0x88, 0xd3, 0xed, 0xb5, 0xac, 0x7a,
    0xcd, 0x52, 0x82, 0xad, 0x6b, 0xab, 0xac, 0x79, 0xa6, 0x77, 0xce, 0x4c, 0xc2, 0x16, 0xf9, 0x19,
    0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c,
    0xa6, 0x77, 0x95, 0x7c, 0xa6, 0x77, 0x95, 0x7c, 0xa6 ,0x77, 0x95, 0x7c,
}

func TestRead(t *testing.T) {
    t.Run("Must return an error on short buffer", func(t *testing.T) {
        b := [] byte{};
        _, err := Read(b);
        if err == nil {
            t.Error("No error on short byte slice")
        }
    })
    t.Run("Must return mutated header", func(t *testing.T) {
        message, _ := Read(mutated);
        t.Run("Must return appropiate query", func(t *testing.T) {
            if message.header.Query != 0x5084 {
                t.Error("Unexpected query!")
            }
        })
        t.Run("Must return appropiate size", func(t *testing.T) {
            if message.header.Size != 0x0104 {
                t.Error( "Unexpected size!")
            }
        })
        t.Run("Must return appropiate sequence", func(t *testing.T) {
            if message.header.Sequence != 0x0286 {
                t.Error( "Unexpected sequence!")
            }
        })
    })
}

func TestMutate(t *testing.T) {
    t.Run("Should return a mutated byte slice", func(t *testing.T) {
        mutation := Mutate(clean);
        if !bytes.Equal(mutation, mutated) {
            t.Error("Bad mutation!")
        }
    })
}
