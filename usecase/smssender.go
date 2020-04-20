package usecase

import (
	"fmt"
	"github.com/yuukichi-nankou/casample-go/domain"
)

type Smssender struct {
	Contacts []domain.Contact
	Message Message
	Client Client
}

type Message interface {
    Get() (string)
}

type Client interface {
    Send(string, string)
}

func (s *Smssender) Run() {
	m := s.Message.Get()

	for k, c := range(s.Contacts) {
		fmt.Println(k)
		fmt.Println(c.GetName())
		s.Client.Send(c.GetNumber(), m)
	}
}