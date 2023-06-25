package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danstis/CrowdSound/internal/version"
	"github.com/stretchr/testify/assert"
)

func Test_versionHandler(t *testing.T) {
	a := New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	a.Router.ServeHTTP(w, req)

	expected, _ := json.Marshal(map[string]string{"version": version.Version})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
