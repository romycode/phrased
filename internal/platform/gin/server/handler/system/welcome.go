package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeResponse struct {
	Data string `json:"data"`
}

func newWelcomeResponse() *welcomeResponse {
	return &welcomeResponse{Data: "!~ Go(lang) powered API ~!"}
}

// RegisterWelcomeHandler register handler to show welcome info.
func RegisterWelcomeHandler(r *gin.RouterGroup) {
	r.GET(
		"/",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, newWelcomeResponse())
			return
		},
	)
}
