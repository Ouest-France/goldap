package goldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// CreateOrganizationalUnit creates ldap OrganizationalUnit
func (c *Client) CreateOrganizationalUnit(dn, description, managedBy string) error {

	req := ldap.NewAddRequest(dn, []ldap.Control{})
	req.Attribute("objectClass", []string{"organizationalUnit"})

	if description != "" {
		req.Attribute("description", []string{description})
	}

	if managedBy != "" {
		req.Attribute("managedBy", []string{managedBy})
	}

	return c.Conn.Add(req)
}

// ReadOrganizationalUnit reads an OrganizationalUnit
func (c *Client) ReadOrganizationalUnit(dn string) (entries map[string][]string, err error) {

	filter := "(objectClass=organizationalUnit)"

	req := ldap.NewSearchRequest(
		dn,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{"*"},
		[]ldap.Control{},
	)

	sr, err := c.Conn.Search(req)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) == 0 {
		return nil, ldap.NewError(ldap.LDAPResultNoSuchObject, fmt.Errorf("the filter '%s' doesn't match any OU: %s", filter, dn))
	}

	if len(sr.Entries) > 1 {
		return nil, ldap.NewError(ldap.LDAPResultOther, fmt.Errorf("the filter '%s' match more than one OU: %s", filter, dn))
	}

	entries = map[string][]string{}
	for _, entry := range sr.Entries {
		for _, attr := range entry.Attributes {
			if !isExcludedAttribute(attr.Name) {
				entries[attr.Name] = attr.Values
			}
		}
	}

	return entries, nil
}

// UpdateOrganizationalUnit updates ldap OrganizationalUnit description
func (c *Client) UpdateOrganizationalUnitDescription(dn string, description string) error {

	req := ldap.NewModifyRequest(dn, []ldap.Control{})

	if description == "" {
		req.Delete("description", []string{})
	} else {
		req.Replace("description", []string{description})
	}

	return c.Conn.Modify(req)
}

// UpdateOrganizationalUnitManagedBy updates ldap OrganizationalUnit managedBy
func (c *Client) UpdateOrganizationalUnitManagedBy(dn string, managedBy string) error {

	req := ldap.NewModifyRequest(dn, []ldap.Control{})

	if managedBy == "" {
		req.Delete("managedBy", []string{})
	} else {
		req.Replace("managedBy", []string{managedBy})
	}

	return c.Conn.Modify(req)
}

// DeleteOrganizationalUnit deletes the specify OrganizationalUnit
func (c *Client) DeleteOrganizationalUnit(dn string) error {

	req := ldap.NewDelRequest(dn, []ldap.Control{})

	return c.Conn.Del(req)
}
