package gateway

import "github.com/sfreiberg/gotwilio"
import "fmt"

type Client struct {
	Sid   string
	Token string
	From  string
}

func (t *Client) Send(to string, message string) {
	client := gotwilio.NewTwilioClient(t.Sid, t.Token)
	_, exp, err := client.SendSMS(t.From, to, message, "", "")
	fmt.Println(exp)
	fmt.Println(err)
}
