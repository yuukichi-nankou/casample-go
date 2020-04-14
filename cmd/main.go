package main

import (
	"github.com/yuukichi-nankou/casample-go/adapter/controller"
)

func main() {
	t := controller.TwilioSMS{}
	t.Run()
}