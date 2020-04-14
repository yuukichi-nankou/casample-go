package controller

import (
	"os"
	"github.com/yuukichi-nankou/casample-go/usecase"
	"github.com/yuukichi-nankou/casample-go/adapter/gateway"
)

type TwilioSMS struct {
}

func (t *TwilioSMS) Run() {
	s := usecase.Smssender{
		Contacts: &gateway.Contacts{ Path: "./test.list"},
		Message : &gateway.Message{Path:"./test.msg"},
		Client : &gateway.Client{
			Sid: os.Getenv("TWILIO_SID"),
			Token: os.Getenv("TWILIO_TOKEN"),
			From: os.Getenv("TWILIO_FROM"),
		},
	}
	s.Run()
}