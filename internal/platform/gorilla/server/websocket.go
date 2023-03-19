package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	mywebsocket "github.com/romycode/phrased/internal/platform/websocket"
)

// GorillaWebsocket is an implementation of Websocket using gorilla/websocket package
type GorillaWebsocket struct {
	ws *websocket.Conn
}

func (g *GorillaWebsocket) Read(data any) error {
	err := g.ws.ReadJSON(data)
	if err != nil {
		return fmt.Errorf("gorillaWebsocket error: %w", err)
	}

	return nil
}

func (g *GorillaWebsocket) Write(data any) error {
	err := g.ws.WriteJSON(data)
	if err != nil {
		return fmt.Errorf("gorillaWebsocker error: %w", err)
	}

	return nil
}

// GorillaWebsocketPromoter is an implementation of Promoter using gorilla/websocket package
type GorillaWebsocketPromoter struct {
	wu *websocket.Upgrader
}

func NewGorillaWebsocketPromoter() *GorillaWebsocketPromoter {
	return &GorillaWebsocketPromoter{
		&websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (g *GorillaWebsocketPromoter) Promote(w http.ResponseWriter, r *http.Request) (mywebsocket.Websocket, error) {
	ws, err := g.wu.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return &GorillaWebsocket{ws}, nil
}
