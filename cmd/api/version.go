package main

import (
	"net/http"

	"github.com/danstis/CrowdSound/internal/version"
)

type (
	versionInfo struct {
		Version string `json:"version"`
	}
)

func (a *App) getVersion(w http.ResponseWriter, r *http.Request) {
	version := versionInfo{
		Version: version.Version,
	}
	respondWithJSON(w, http.StatusOK, version)
}
