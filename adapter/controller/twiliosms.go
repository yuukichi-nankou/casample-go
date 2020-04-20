package controller

import (
	"os"
	"github.com/yuukichi-nankou/casample-go/usecase"
	"github.com/yuukichi-nankou/casample-go/adapter/gateway"
)

type TwilioSMS struct {
}

func (t *TwilioSMS) Run() {
	cs := gateway.Contacts{ Path: "./local.contacts.csv"}
	s := usecase.Smssender{
		Contacts: cs.FetchContacts(),
		Message : &gateway.Message{Path:"./local.msg.txt"},
		Client : &gateway.Client{
			Sid: os.Getenv("TWILIO_SID"),
			Token: os.Getenv("TWILIO_TOKEN"),
			From: os.Getenv("TWILIO_FROM"),
		},
	}
	s.Run()
}