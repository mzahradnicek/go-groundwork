package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type JsonMapper struct {
	reader  io.Reader
	payload []byte
}

func (jm *JsonMapper) Decode(v interface{}) (err error) {
	// check for first time encoding
	if len(jm.payload) == 0 {
		jm.payload, err = ioutil.ReadAll(jm.reader)
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(jm.payload, v)
	return
}

func NewJsonMapper(r io.Reader) *JsonMapper {
	return &JsonMapper{reader: r}
}
