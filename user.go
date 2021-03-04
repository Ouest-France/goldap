package goldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// ReadUser reads ldap user and return it's attributes on an error if the user donesn't exist
func (c *Client) ReadUser(ou string, name string, sAMAccountName string, upn string) (entries map[string][]string, err error) {
	filter := ""

	if name != "" {
		filter = fmt.Sprintf("%s(name=%s)", filter, name)
	}

	if sAMAccountName != "" {
		filter = fmt.Sprintf("%s(sAMAccountName=%s)", filter, sAMAccountName)
	}

	if upn != "" {
		filter = fmt.Sprintf("%s(userPrincipalName=%s)", filter, upn)
	}

	filter = fmt.Sprintf("(&(objectCategory=person)(objectClass=user)%s)", filter)

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
