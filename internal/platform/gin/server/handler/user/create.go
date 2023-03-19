package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/romycode/phrased/internal/user"
)

type createUserRequest struct {
	ID string `json:"-,omitempty"`
}

func RegisterCreateUserHandler(r *gin.RouterGroup, cu *user.CreateUserService) {
	r.POST(
		"/users/:id",
		func(c *gin.Context) {
			ID := c.Param("id")
			req := &createUserRequest{ID}

			_, err := cu.Execute(req.ID)
			if err != nil {
				if errors.Is(err, user.AlreadyExists) {
					c.Status(http.StatusCreated)
				}
				return
			}

			c.Status(http.StatusCreated)
			return
		},
	)
}
