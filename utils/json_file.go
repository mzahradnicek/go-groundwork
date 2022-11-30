package utils

import (
	"bytes"
)

type JsonFile struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func (jf JsonFile) GetReader() *bytes.Reader {
	return bytes.NewReader(jf.Data)
}
