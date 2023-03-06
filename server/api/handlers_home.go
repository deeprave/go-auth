package api

import (
	"github.com/deeprave/go-auth/lib"
	"net/http"
)

func (app *App) Home(w http.ResponseWriter, _ *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		AppId   string `json:"app-id"`
		Version string `json:"version"`
	}{
		Status:  "active",
		AppId:   app.appname,
		Version: app.version,
	}

	logIfError(lib.JSONWriteResponse(w, payload, http.StatusOK))
}
