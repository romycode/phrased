package system

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest(http.MethodGet, "/v1/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	engine := gin.New()
	RegisterStatusHandler(engine.Group("/v1"))
	engine.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("sut returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	// Check the response body is what we expect
	expected := `{"data":{"status":"Up","details":"Connectivity ok"}}`
	assert.JSONEq(t, expected, rr.Body.String(), fmt.Sprintf("sut returned unexpected body: got %v want %v", rr.Body.String(), expected))
}
