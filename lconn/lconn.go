package lconn

// Interface interface
type Interface interface {
	Open()
	Ping()
	Reconnect()
}
