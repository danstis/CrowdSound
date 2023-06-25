package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_respondWithJSON(t *testing.T) {
	type args struct {
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Version",
			args: args{
				code:    http.StatusOK,
				payload: map[string]string{"version": "0.1.0"},
			},
		},
		{
			name: "Simple string",
			args: args{
				code:    http.StatusOK,
				payload: "Test 123",
			},
		},
		{
			name: "Empty payload",
			args: args{
				code:    http.StatusOK,
				payload: nil,
			},
		},
		{
			name: "Crazy quotes",
			args: args{
				code:    http.StatusOK,
				payload: "`test`,'test',\"test\"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respondWithJSON(w, tt.args.code, tt.args.payload)

			er, _ := json.Marshal(tt.args.payload)
			assert.Equal(t, tt.args.code, w.Code)
			assert.Equal(t, string(er), w.Body.String())
		})
	}
}

func Test_respondWithError(t *testing.T) {
	type args struct {
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "StatusInternalServerError",
			args: args{
				code:    http.StatusInternalServerError,
				message: "Internal Server Error",
			},
		},
		{
			name: "StatusNotFound",
			args: args{
				code:    http.StatusNotFound,
				message: "Not Found",
			},
		},
		{
			name: "StatusUnauthorized",
			args: args{
				code:    http.StatusUnauthorized,
				message: "Unauthorized",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respondWithError(w, tt.args.code, tt.args.message)
			assert.Equal(t, tt.args.code, w.Code)
			assert.Equal(t, "{\"error\":\""+tt.args.message+"\"}", w.Body.String())
		})
	}
}
