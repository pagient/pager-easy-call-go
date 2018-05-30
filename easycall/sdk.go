package easycall

import (
	"net/http"
	"time"
	"fmt"
)

const (
	pathSend = "%s/send"
)

// ClientAPI describes a client API
type ClientAPI interface {
	// SetCredentials sets authentication credentials
	SetCredentials(string, string)

	// PagerList returns a list of all pagers
	Send(*SendOptions) error
}

// Default implements the client interface
type Default struct {
	client      *http.Client
	base        string
	username    string
	password    string
}

func NewClient(uri string) ClientAPI {
	return &Default{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		base:   uri,
	}
}

func (c *Default) SetCredentials(username, password string) {
	c.username = username
	c.password = password
}

func (c *Default) Send(options *SendOptions) error {
	uri := fmt.Sprintf(pathSend, c.base)
	err := c.post(uri, options, nil)

	return err
}
