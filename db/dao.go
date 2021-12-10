package db

type IDao interface {
	Client() *Client
	Prefix(prefix string)
}
