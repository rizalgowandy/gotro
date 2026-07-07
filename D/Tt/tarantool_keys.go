package Tt

import "github.com/vmihailenco/msgpack/v5"

// Uint64Key serializes a single unsigned integer key without narrowing it to
// Go's platform-dependent uint type.
type Uint64Key struct {
	I uint64
}

func (k Uint64Key) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(1); err != nil {
		return err
	}
	return enc.EncodeUint(k.I)
}

// Int64Key serializes a single signed integer key without narrowing it to
// Go's platform-dependent int type.
type Int64Key struct {
	I int64
}

func (k Int64Key) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(1); err != nil {
		return err
	}
	return enc.EncodeInt(k.I)
}
