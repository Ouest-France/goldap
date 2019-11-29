package goldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// Client represents an LDAP client instance
type Client struct {
	Conn         *ldap.Conn
	Host         string
	Port         int
	BindUser     string
	BindPassword string
	TLS          bool
}

func (c *Client) Connect() error {

	uri := fmt.Sprintf("%s:%d", c.Host, c.Port)

	l, err := ldap.Dial("tcp", uri)
	if err != nil {
		return fmt.Errorf("error dialing: %s", err)
	}

	err = l.Bind(c.BindUser, c.BindPassword)
	if err != nil {
		return fmt.Errorf("error binding: %s", err)
	}
	c.Conn = l

	return nil
}
