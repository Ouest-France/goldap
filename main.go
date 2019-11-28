package goldap

import (
	"gopkg.in/ldap.v2"
)

// Client represents an LDAP client instance
type Client struct {
	Client   *ldap.Conn
	Host     string
	Port     int
	User     string
	Password string
	TLS      bool
}
