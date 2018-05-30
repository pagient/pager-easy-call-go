package pagient

import (
	"net/http"
	"time"
	"fmt"
)

const (
	pathPagers = "%s/pagers"
)

// ClientAPI describes a client API
type ClientAPI interface {
	// SetCredentials sets authentication credentials
	SetCredentials(string, string)

	// PagerList returns a list of all pagers
	PagerList() ([]*Pager, error)
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

func (c *Default) PagerList() ([]*Pager, error) {
	var out []*Pager

	uri := fmt.Sprintf(pathPagers, c.base)
	err := c.get(uri, &out)

	return out, err
}