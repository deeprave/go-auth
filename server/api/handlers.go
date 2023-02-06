package api

import (
	"errors"
	"fmt"
	"github.com/deeprave/go-auth/lib"
	"github.com/deeprave/go-auth/models"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

func logIfError(err error) bool {
	pc, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%s(%d):%v", file, line, pc)
	if err != nil {
		log.Printf("%s: %v", prefix, err)
		return true
	}
	return false
}

func isLocalRequest(r *http.Request) bool {
	return r.RemoteAddr == "localhost"
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {
	auth := &app.cfg.Auth
	// read json payload for username, password and possibly totp
	payload := models.AuthRequest{}

	err := lib.JSONReadRequest(w, r, &payload)
	if err != nil {
		logIfError(lib.JSONErrorResponse(w, err, http.StatusBadRequest))
		return
	}

	// validate user & check credentials
	var (
		valid = true
		user  *models.User
	)
	if user, err = app.DB.GetUserByName(payload.Username); err == nil {
		var credentials []*models.Credential
		if credentials, err = app.DB.GetCredentialsForUser(user.Id); err == nil {
			valid, err = payload.Matches(credentials)
		}
	}
	if err != nil || !valid {
		logIfError(lib.JSONErrorResponse(w, errors.New("invalid credentials"), http.StatusForbidden))
		return
	}

	// generate the token
	var tokens TokenPairs
	if tokens, err = auth.GenerateTokenPair(user.ToJWTUser()); err != nil {
		logIfError(lib.JSONErrorResponse(w, err))
		return
	}

	refreshCookie := auth.GetRefreshCookie(tokens.RefreshToken, isLocalRequest(r))
	http.SetCookie(w, refreshCookie)

	logIfError(lib.JSONWriteResponse(w, &tokens, http.StatusOK))
}

func (app *App) logout(w http.ResponseWriter, r *http.Request) {
	auth := app.cfg.Auth
	http.SetCookie(w, auth.GetExpiredRefreshCookie(isLocalRequest(r)))
	w.WriteHeader(http.StatusAccepted)
}

func (app *App) refresh(w http.ResponseWriter, r *http.Request) {
	auth := app.cfg.Auth
	errorText := http.StatusText(http.StatusUnauthorized)
	for _, cookie := range r.Cookies() {
		if cookie.Name == auth.Cookie.Name {
			claims := &Claims{}
			refreshToken := cookie.Value
			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.Secret), nil
			})
			if err == nil {
				var userID int64
				if userID, err = strconv.ParseInt(claims.Subject, 0, 64); err == nil {
					var user *models.User
					if user, err = app.DB.GetUserById(userID); err == nil {
						jwtUser := user.ToJWTUser()
						var tokens TokenPairs
						if tokens, err = auth.GenerateTokenPair(jwtUser); err != nil {
							errorText = "error generating tokens"
						} else {
							http.SetCookie(w, auth.GetRefreshCookie(tokens.RefreshToken, isLocalRequest(r)))
							logIfError(lib.JSONWriteResponse(w, &tokens, http.StatusOK))
							return
						}
					}
				}
			}
			// found our cookie so stop iteration
			break
		}
	}
	logIfError(lib.JSONErrorResponse(w, errors.New(errorText), http.StatusUnauthorized))
}

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
