package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type internalServerError struct {
	Error string `json:"error"`
}

// RegisterRecovery register a gin.HandlerFunc to recover panics that logs the recovered
// panic and aborts the HTTP request returning an Internal Server Error (500).
func RegisterRecovery(engine *gin.Engine) {
	engine.Use(
		func(c *gin.Context) {
			// Recover from panic
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf(
						"[Middleware] %s panic recovered:\n%s\n",
						time.Now().Format("2006/01/02 - 15:04:05"),
						err,
					)

					c.JSON(http.StatusInternalServerError, internalServerError{"unexpected error"})
					return
				}
			}()

			c.Next()
		},
	)
}
