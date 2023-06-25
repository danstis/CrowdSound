package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pingHandler(t *testing.T) {
	a := New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	a.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
