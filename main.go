package goldap

import (
	"crypto/tls"
	"crypto/x509"
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
	TLSCACert    string
	TLSInsecure  bool
}

// Connect Creates un ldap connection
func (c *Client) Connect() error {
	uri := fmt.Sprintf("%s:%d", c.Host, c.Port)

	if c.TLS {
		// Get the System Cert Pool
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			return fmt.Errorf("error tls: %s", err)
		}

		// Use the provided CA certificate
		if c.TLSCACert != "" {
			caCertPool = x509.NewCertPool()
			if ok := caCertPool.AppendCertsFromPEM([]byte(c.TLSCACert)); !ok {
				return fmt.Errorf("error tls: Can't add the CA certificate to certificate pool")
			}
		}

		conn, err := ldap.DialTLS("tcp", uri, &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: c.TLSInsecure,
		})
		if err != nil {
			return fmt.Errorf("error dialing: %s", err)
		}
		c.Conn = conn

	} else {
		conn, err := ldap.Dial("tcp", uri)
		if err != nil {
			return fmt.Errorf("error dialing: %s", err)
		}
		c.Conn = conn
	}

	err := c.Conn.Bind(c.BindUser, c.BindPassword)
	if err != nil {
		return fmt.Errorf("error binding: %s", err)
	}

	return nil
}
