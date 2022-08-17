package goldap

import (
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

// ReadUser reads ldap user and return it's attributes or an error if the user doesn't exist
func (c *Client) ReadUser(ou string, name string, sAMAccountName string, upn string) (entries map[string][]string, err error) {
	conditions := []string{"(objectCategory=person)", "(objectClass=user)"}

	if name != "" {
		conditions = append(conditions, fmt.Sprintf("(name=%s)", name))
	}

	if sAMAccountName != "" {
		conditions = append(conditions, fmt.Sprintf("(sAMAccountName=%s)", sAMAccountName))
	}

	if upn != "" {
		conditions = append(conditions, fmt.Sprintf("(userPrincipalName=%s)", upn))
	}

	filter := fmt.Sprintf("(&%s)", strings.Join(conditions, ""))

	return c.ReadUserByFilter(ou, filter)
}

// ReadUserByFilter reads ldap user and return it's attributes or an error if the user doesn't exist
func (c *Client) ReadUserByFilter(ou string, filter string) (entries map[string][]string, err error) {
	req := ldap.NewSearchRequest(
		ou,
		ldap.ScopeWholeSubtree,
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
		return nil, ldap.NewError(ldap.LDAPResultNoSuchObject, fmt.Errorf("The filter '%s' doesn't match any user in the OU: %s", filter, ou))
	}

	if len(sr.Entries) > 1 {
		return nil, ldap.NewError(ldap.LDAPResultOther, fmt.Errorf("The filter '%s' match more than one user in the OU: %s", filter, ou))
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

func isExcludedAttribute(attr string) bool {
	// We exclude sensitive attributes like password
	excludedAttrs := []string{"unicodePwd", "userPassword"}

	for _, a := range excludedAttrs {
		if attr == a {
			return true
		}
	}
	return false
}
