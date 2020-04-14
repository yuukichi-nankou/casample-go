package gateway

import "io/ioutil"

type Message struct {
	Path string
}

func (m *Message) Get() string {
	bytes, err := ioutil.ReadFile(m.Path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
