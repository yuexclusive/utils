package email

import "gopkg.in/gomail.v2"

type Dialer struct {
	*gomail.Dialer
}

func (d *Dialer) Send(subject string, contentType string, body string, tos ...string) error {
	if c, err := d.Dial(); err != nil {
		return err
	} else {
		msg := gomail.NewMessage()

		if contentType == "" {
			contentType = "text/plain"
		}

		msg.SetHeader("From", d.Host)
		msg.SetHeader("To", tos...)
		msg.SetHeader("Subject", subject)
		msg.SetBody(contentType, body)
		return c.Send(d.Host, tos, msg)
	}
}

func NewDialer(host string, port int, username, password string) *Dialer {
	return &Dialer{gomail.NewDialer(host, port, username, password)}
}
