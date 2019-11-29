package goldap

import (
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) CreateGroup(dn, name string, members []string) error {

	req := ldap.NewAddRequest(dn, []ldap.Control{})
	req.Attribute("objectClass", []string{"group"})
	req.Attribute("sAMAccountName", []string{name})

	if len(members) > 0 {
		req.Attribute("member", members)
	}

	return c.Conn.Add(req)
}

func (c *Client) ReadGroup(dn string) (attributes map[string][]string, err error) {

	req := ldap.NewSearchRequest(
		dn,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=group)",
		[]string{"*"},
		[]ldap.Control{},
	)

	sr, err := c.Conn.Search(req)
	if err != nil {
		return attributes, err
	}

	attributes = map[string][]string{}
	for _, entry := range sr.Entries {
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
	}

	return attributes, nil
}

func (c *Client) DeleteGroup(dn string) error {

	req := ldap.NewDelRequest(dn, []ldap.Control{})

	return c.Conn.Del(req)
}
