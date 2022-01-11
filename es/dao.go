package es

import elastic "github.com/olivere/elastic/v7"

type Dao struct {
	Client *elastic.Client
}

func NewCao() *Dao {
	return &Dao{Client: GetClient()}
}
