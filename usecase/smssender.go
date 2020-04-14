package usecase

import "fmt"

type Smssender struct {
	Contacts Contacts
	Message Message
	Client Client
}

type Contacts interface {
    FetchNumbers() ([]string)
}

type Message interface {
    Get() (string)
}

type Client interface {
    Send(string, string)
}

func (s *Smssender) Run() {
	ns := s.Contacts.FetchNumbers()
	m := s.Message.Get()

	for k, n := range(ns) {
		fmt.Println(k)
		fmt.Println(n)
		s.Client.Send(n, m)
	}
}