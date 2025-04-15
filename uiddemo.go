// Package UIDDemo a UIDDemo plugin.
package uiddemo

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/rwxrob/uniq"
)

// Config the plugin configuration.
type Config struct {
	// ...
	HeaderName string
	Mystring string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		// ...
		HeaderName: "X-Traefik-UIDDemoID",
	}
}

// UIDDemo holds required components of this plugin.
type UIDDemo struct {
	next       http.Handler
	headerName string
	name       string
	mystring   string
	// ...
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...

	if len(config.HeaderName) == 0 {
		return nil, fmt.Errorf("no header name provided from dev branch")
	}
	return &UIDDemo{
		next:       next,
		headerName: config.HeaderName,
		name:       name,
		mystring:    config.Mystring,
	}, nil
}

func (u *UIDDemo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	uid := uniq.UUID()
	req.Header.Set(u.headerName, uid)
	req.Header.Set(u.mystring, "thi")
	rw.Header().Set(u.headerName, uid)

	rw.Header().Set(u.mystring, "this")
	// rw.Write([]byte("Hello you"))
	u.next.ServeHTTP(rw, req)
}
