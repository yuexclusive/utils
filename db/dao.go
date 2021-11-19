package db

type Dao interface {
	Init(c *Client)
}
