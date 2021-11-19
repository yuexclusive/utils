package es

import elastic "github.com/olivere/elastic/v7"

type Cao struct {
	Client *elastic.Client
}

func NewCao() *Cao {
	return &Cao{Client: Client()}
}
