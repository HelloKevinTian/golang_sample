package utils

import (
	"bytes"
	"encoding/gob"
)

func Encode(params interface{}) bytes.Buffer {
	var buffer bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buffer) // Will write to network.
	enc.Encode(params)
	return buffer
}

func Decode(buffer bytes.Buffer, result interface{}) {
	dec := gob.NewDecoder(&buffer) // Will read from network.
	dec.Decode(result)
}
