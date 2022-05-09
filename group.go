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

// SearchGroupByName searches an LDAP group by name and returns its DN
func (c *Client) SearchGroupByName(name, ou string, scope int) (string, error) {

	// Request name and description attributes
	req := ldap.NewSearchRequest(
		ou,
		scope,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=group)(cn=%s))", name),
		[]string{},
		[]ldap.Control{},
	)

	// Search for group
	sr, err := c.Conn.Search(req)
	if err != nil {
		return "", fmt.Errorf("searching group by name: %s", err)
	}

	// If no entries found, group doesn't exists, return error
	if len(sr.Entries) == 0 {
		return "", fmt.Errorf("group %q not found in OU %q", name, ou)
	}

	// If more than one entry, it's an error
	if len(sr.Entries) > 1 {
		return "", fmt.Errorf("more than one group found with name %q in OU %q", name, ou)
	}

	// Return group DN
	return sr.Entries[0].DN, nil
}

// ReadGroup reads ldap group and return it's attributes on an error if the group donesn't exist
func (c *Client) ReadGroup(dn string, memberPageSize int) (attributes map[string][]string, err error) {

	// Request name and description attributes
	req := ldap.NewSearchRequest(
		dn,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=group)",
		[]string{"name", "description"},
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

	// Page member attribute
	page := 0
	attributes["member"] = []string{}
	for {
		fmt.Printf("member;range=%d-%d\n", memberPageSize*page, (memberPageSize-1)+memberPageSize*page)

		// Request member attribute page
		req := ldap.NewSearchRequest(
			dn,
			ldap.ScopeBaseObject,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			"(objectclass=group)",
			[]string{fmt.Sprintf("member;range=%d-%d", memberPageSize*page, (memberPageSize-1)+memberPageSize*page)},
			[]ldap.Control{},
		)

		// Search
		sr, err := c.Conn.Search(req)
		if err != nil {
			return attributes, err
		}

		// If no attributes found, member doesn't exists, break
		if len(sr.Entries) == 0 {
			break
		}

		// If more than one entry, it's an error
		if len(sr.Entries) > 1 {
			return attributes, fmt.Errorf("more than one entry found retrieving group members")
		}

		// If no attributes found, member doesn't exists, break
		if len(sr.Entries[0].Attributes) == 0 {
			break
		}

		// Append member attribute
		attributes["member"] = append(attributes["member"], sr.Entries[0].Attributes[0].Values...)

		// If no more pages, break
		if len(sr.Entries[0].Attributes[0].Values) < memberPageSize {
			break
		}

		// Next page
		page++
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
