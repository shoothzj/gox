package testx

import (
	"encoding/hex"
	"github.com/shoothzj/gox/buffer"
	"github.com/stretchr/testify/require"
	"testing"
)

func Hex2Bytes(t *testing.T, str string) []byte {
	bytes, err := hex.DecodeString(str)
	require.NoError(t, err)
	return bytes
}

func Hex2Buffer(t *testing.T, str string) *buffer.Buffer {
	bytes, err := hex.DecodeString(str)
	require.NoError(t, err)
	return buffer.NewBufferFromBytes(bytes)
}
