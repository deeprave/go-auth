package api

import (
	"errors"
	"fmt"
	"github.com/deeprave/go-auth/models"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	Issuer        string        `yaml:"iss,omitempty"`
	Audience      string        `yaml:"aud,omitempty"`
	Secret        string        `yaml:"secret,omitempty"`
	TokenExpiry   time.Duration `yaml:"exp,omitempty"`
	RefreshExpiry time.Duration `yaml:"rxp,omitempty"`
	Cookie        Cookie        `yaml:"cookie,omitempty"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const (
	DefaultTokenExpiry  = "30m"
	DefaultTokenRefresh = "15m"
)

type Claims struct {
	jwt.RegisteredClaims
}

func ParseDuration(value string) time.Duration {
	dur, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("duration: %s - %v", value, err)
	}
	return dur
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	b := make([]byte, n)
	m := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(m)]
	}
	return string(b)
}

func (j *Auth) GenerateTokenPair(jwtUser *models.JWTUser) (TokenPairs, error) {
	// create a token
	token := jwt.New(jwt.SigningMethodHS256)
	// set the token claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = jwtUser.LongName()
	claims["sub"] = fmt.Sprint(jwtUser.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().Unix()
	claims["typ"] = "JWT"
	// set the expiry
	claims["exp"] = time.Now().Add(j.TokenExpiry).Unix()

	// create signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create refresh token & set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = fmt.Sprint(jwtUser.ID)
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["typ"] = "JWT"

	// set expiry for the refresh
	refreshClaims["exp"] = time.Now().Add(j.RefreshExpiry).Unix()

	// create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create token pairs and populate signed tokens & return
	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string, local bool) *http.Cookie {
	return j.Cookie.GetRefreshCookie(refreshToken, j.RefreshExpiry, local)
}

func (j *Auth) GetExpiredRefreshCookie(local bool) *http.Cookie {
	return j.Cookie.GetExpiredRefreshCookie(local)
}

func (j *Auth) GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	w.Header().Add("Vary", "Authorization")

	// get auth header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader != "" {
		// split the header on spaces
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) == 2 {
			// check to see if we have the word Bearer
			if headerParts[0] == "Bearer" {
				token := headerParts[1]
				// declare an empty claims
				claims := &Claims{}

				// parse the token
				_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(j.Secret), nil
				})

				if err != nil {
					if strings.HasPrefix(err.Error(), "token is expired by") {
						return "", nil, errors.New("expired token")
					}
					return "", nil, err
				}
				if claims.Issuer == j.Issuer {
					return token, claims, nil
				}
			}
		}
	}
	return "", nil, errors.New("invalid/missing auth header")
}
