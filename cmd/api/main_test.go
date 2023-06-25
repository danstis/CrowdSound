package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/danstis/CrowdSound/cmd/api"
	"github.com/gorilla/mux"
)

var a main.App

func TestMain(m *testing.M) {
	a.Router = mux.NewRouter()
	code := m.Run()
	os.Exit(code)
}

// executeRequest returns a httptest.ResponseRecorder by serving the request using
// the router.
//
// req *http.Request - The http.Request to be served.
// returns *httptest.ResponseRecorder - The ResponseRecorder containing the http.Response.
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Initialize()
	a.Router.ServeHTTP(rr, req)
	return rr
}
