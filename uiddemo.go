// Package UIDDemo a UIDDemo plugin.
package uiddemo

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/rwxrob/uniq"
)

const defaultHeader = "X-Traefik-UIDDemoID"

// Config the plugin configuration.
type Config struct {
	// ...
	HeaderName string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		// ...
		HeaderName: defaultHeader,
	}
}

// UIDDemo holds required components of this plugin.
type UIDDemo struct {
	next       http.Handler
	headerName string
	name       string
	// ...
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...

	if len(config.HeaderName) == 0 {
		return nil, fmt.Errorf("no header name provided")
	}
	return &UIDDemo{
		next:       next,
		headerName: config.HeaderName,
		name:       name,
	}, nil
}

func (u *UIDDemo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	uid := uniq.UUID()
	req.Header.Set(u.headerName, uid)
	rw.Header().Set(u.headerName, uid)
	// rw.Write([]byte("Hello you"))
	u.next.ServeHTTP(rw, req)
}
