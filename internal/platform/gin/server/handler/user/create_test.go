package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/romycode/phrased/internal/user"
)

func TestCreateUserHandler(t *testing.T) {
	type args struct {
		ID string
		ur user.Repository
	}
	cases := []struct {
		name               string
		args               args
		expectedStatusCode uint
	}{
		{
			name: "itShouldReturnStatusCreated",
			args: args{
				ID: "666f7ece-cdb7-4392-a914-7ff0ffc988a7",
				ur: user.NewInMemoryRepository(map[string]*user.User{}),
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "itShouldReturnCreatedIfTheUserAlreadyExists",
			args: args{
				ID: "666f7ece-cdb7-4392-a914-7ff0ffc988a7",
				ur: user.NewInMemoryRepository(map[string]*user.User{"666f7ece-cdb7-4392-a914-7ff0ffc988a7": {ID: "666f7ece-cdb7-4392-a914-7ff0ffc988a7"}}),
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request to pass to our handler
			req, err := http.NewRequest(http.MethodPost, "/v1/users/"+tc.args.ID, strings.NewReader(""))
			assert.NoError(t, err)
			// We create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			engine := gin.New()
			cu := user.NewCreateUserService(tc.args.ur)
			RegisterCreateUserHandler(engine.Group("/v1"), cu)
			engine.ServeHTTP(rr, req)

			// Check the status code is what we expect
			assert.Equal(
				t,
				http.StatusCreated,
				rr.Code,
				fmt.Sprintf("sut returned wrong status code: got %v want %v", rr.Code, http.StatusCreated),
			)
		})
	}
}
