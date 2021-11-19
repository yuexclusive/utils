package pulsar

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

// IClient IClient
type IClient interface {
	pulsar.Client
}

// Client Client
type Client struct {
	pulsar.Client
}

// NewClient NewClient
func NewClient(url string) (IClient, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})
	if err != nil {
		return nil, err
	}
	return &Client{Client: client}, nil
}
