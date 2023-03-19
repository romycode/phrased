package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	mywebsocket "github.com/romycode/phrased/internal/platform/websocket"
)

var connections = make(map[string]mywebsocket.Websocket)

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Kind     string `json:"kind"`
	Content  string `json:"content"`
}

type EchoMessage struct {
	Content string `json:"content"`
}

type websocketPromotionError struct {
	Error string `json:"error"`
}

type missingUserHeaderError struct {
	Error string `json:"error"`
}

// RegisterWebsocketHandler register handler to manage websocket connections.
func RegisterWebsocketHandler(r *gin.RouterGroup, promoter mywebsocket.Promoter) {
	r.GET("/chat/ws", func(c *gin.Context) {
		userID := c.GetHeader("x-user-id")
		if "" == userID {
			c.JSON(http.StatusBadRequest, missingUserHeaderError{"missing x-user-id header"})
			return
		}

		ws, err := promoter.Promote(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, websocketPromotionError{"error creating websocket"})
			return
		}

		connections[userID] = ws

		go func() {
			log.Printf("User: %s, listening ws inbox", userID)
			for {
				message := &Message{}
				err := ws.Read(message)
				if err != nil {
					log.Printf("ERROR: %s", err)
					break
				}

				switch message.Kind {
				case "echo":
					_ = ws.Write(EchoMessage{Content: message.Content})
				case "message":
					receiver, ok := connections[message.Receiver]

					if ok {
						_ = receiver.Write(message)
					}
				}

				log.Printf("User: %s, received ws message: %#v", userID, message)
			}
		}()
	})
}
