package authx

import (
	"encoding/base64"
	"strings"
)

// ParseBasicAuth parses a Basic Authentication header and returns the username, password, and a boolean indicating success.
func ParseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// BasicAuth returns a Basic Authentication header string for the given username and password.
func BasicAuth(username, passwd string) string {
	auth := username + ":" + passwd
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
