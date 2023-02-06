package api

import (
	"fmt"
	"net/http"
	"time"
)

type Cookie struct {
	Prefix string `yaml:"prefix,omitempty"`
	Domain string `yaml:"domain,omitempty"`
	Path   string `yaml:"path,omitempty"`
	Name   string `yaml:"name,omitempty"`
}

func (c *Cookie) CookieName() string {
	prefix := "__Host-"
	if c.Prefix != "" {
		prefix = c.Prefix
	}
	return fmt.Sprintf("%s%s", prefix, c.Name)
}

func (c *Cookie) CookieDomain(local bool) string {
	if local {
		return ""
	}
	return c.Domain
}

func (c *Cookie) GetRefreshCookie(refreshToken string, refreshExpiry time.Duration, local bool) *http.Cookie {
	return &http.Cookie{
		Name:     c.CookieName(),
		Path:     c.Path,
		Value:    refreshToken,
		Expires:  time.Now().Add(refreshExpiry),
		MaxAge:   int(refreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   c.CookieDomain(local),
		HttpOnly: true,
		Secure:   true,
	}
}

func (c *Cookie) GetExpiredRefreshCookie(local bool) *http.Cookie {
	return &http.Cookie{
		Name:     c.CookieName(),
		Path:     c.Path,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   c.CookieDomain(local),
		HttpOnly: true,
		Secure:   true,
	}
}
