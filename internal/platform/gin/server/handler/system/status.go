package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Data statusData `json:"data"`
}

type statusData struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

func newStatusResponse() *statusResponse {
	return &statusResponse{
		Data: struct {
			Status  string `json:"status"`
			Details string `json:"details"`
		}{
			Status:  "Up",
			Details: "Connectivity ok",
		},
	}
}

// RegisterStatusHandler register handler to show api status info.
func RegisterStatusHandler(r *gin.RouterGroup) {
	r.GET(
		"/status",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, newStatusResponse())
			return
		},
	)
}
