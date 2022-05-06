package goldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// CreateGroup creates ldap group
func (c *Client) CreateGroup(dn, name string, description string, members []string) error {

	req := ldap.NewAddRequest(dn, []ldap.Control{})
	req.Attribute("objectClass", []string{"group"})
	req.Attribute("sAMAccountName", []string{name})

	if description != "" {
		req.Attribute("description", []string{description})
	}

	if len(members) > 0 {
		req.Attribute("member", members)
	}

	return c.Conn.Add(req)
}

// ReadGroup searchs for a ldap group and returns a map of its attributes. If not found, an empty map
func (c *Client) ReadGroup(baseDn string, groupName string, memberPageSize int) (attributes map[string][]string, err error) {

	var pageSize uint32 = 32
	pagingControl := ldap.NewControlPaging(pageSize)
	// Create request to search for group attributes
	req := ldap.NewSearchRequest(
		baseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0,
		false,
		fmt.Sprintf("(&(objectClass=group)(cn=%s))", groupName),
		[]string{},
		[]ldap.Control{pagingControl},
	)

	// Search group
	sr, err := c.Conn.Search(req)
	if err != nil {
		return attributes, err
	}

	// Fills attribute map with entry key/values
	attributes = map[string][]string{}
	for _, entry := range sr.Entries {
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
	}
	return attributes, nil
}

// UpdateGroupDescription updates ldap group description
func (c *Client) UpdateGroupDescription(dn string, description string) error {

	req := ldap.NewModifyRequest(dn, []ldap.Control{})

	if description == "" {
		req.Delete("description", []string{})
	} else {
		req.Replace("description", []string{description})
	}

	return c.Conn.Modify(req)
}

// UpdateGroupMembers updates ldap group members
func (c *Client) UpdateGroupMembers(dn string, members []string) error {

	req := ldap.NewModifyRequest(dn, []ldap.Control{})

	if len(members) == 0 {
		req.Delete("member", members)
	} else {
		req.Replace("member", members)
	}

	return c.Conn.Modify(req)
}

// UpdateGroupType updates ldap group type
func (c *Client) UpdateGroupType(dn string, groupType string) error {

	req := ldap.NewModifyRequest(dn, []ldap.Control{})

	req.Replace("groupType", []string{groupType})

	return c.Conn.Modify(req)
}

// DeleteGroup deletes the specify group
func (c *Client) DeleteGroup(dn string) error {

	req := ldap.NewDelRequest(dn, []ldap.Control{})

	return c.Conn.Del(req)
}
