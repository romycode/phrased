package websocket

import (
	"net/http"
)

// Websocket is an interface to wrap the implementation used
type Websocket interface {
	Read(data any) error
	Write(data any) error
}

// Promoter is an interface to wrap the
type Promoter interface {
	Promote(w http.ResponseWriter, r *http.Request) (Websocket, error)
}
