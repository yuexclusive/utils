package nats

import (
	"github.com/nats-io/nats.go"
)

func Publish(addr string, subj string, bytes []byte) error {
	nc, err := nats.Connect(addr, nats.UserInfo("", ""))
	if err != nil {
		return err
	}
	defer nc.Close()

	return nc.Publish(subj, bytes)
	// return nc.Flush()
}

func Subscribe(addr string, subj string, handler nats.MsgHandler) error {
	nc, err := nats.Connect(addr)
	if err != nil {
		return err
	}
	// defer nc.Close()
	nc.Subscribe(subj, handler)
	return nil
}
