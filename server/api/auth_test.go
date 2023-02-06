package api

import (
	"fmt"
	"github.com/deeprave/go-auth/models"
	"github.com/deeprave/go-testutils/test"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
	"unicode"
)

func TestParseDuration(t *testing.T) {
	durations := []string{"10m", "2h", "34h17m36s"}
	durValues := []time.Duration{600000000000, 7200000000000, 123456000000000}
	for i, value := range durations {
		dur := ParseDuration(value)
		val := durValues[i]
		test.ShouldBeEqual(t, dur, val)
	}
}

func TestRandString(t *testing.T) {
	for _, l := range []int{12, 20, 16, 4} {
		s := RandString(l)
		test.ShouldBeEqual(t, l, len(s))
		for _, c := range s {
			result := unicode.IsDigit(c) || unicode.IsLetter(c)
			test.ShouldBeTrue(t, result, "unknown rune; %v", c)
		}
	}
}

func TestGenerateTokenPair(t *testing.T) {
	auth := Auth{
		Issuer:        "Acme, Inc",
		Audience:      "Acme, Inc",
		Secret:        RandString(32),
		TokenExpiry:   ParseDuration(DefaultTokenExpiry),
		RefreshExpiry: ParseDuration(DefaultTokenRefresh),
		Cookie:        Cookie{},
	}
	user := &models.JWTUser{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Given:    "Test",
		Surname:  "User",
	}
	claims := &Claims{}
	tokenPairs, err := auth.GenerateTokenPair(user)
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)

	testToken := func(token string) {
		test.ShouldNotBeEqual(t, token, "")
		v, err := jwt.ParseWithClaims(token, claims,
			func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				if !test.ShouldBeTrue(t, ok, err.Error()) {
					return nil, err
				}
				return []byte(auth.Secret), nil
			})
		test.ShouldBeNoError(t, err, "unexpected error: %v", err)
		test.ShouldBeEqual(t, v.Method.Alg(), "HS256")
		test.ShouldBeEqual(t, v.Header["alg"], "HS256")
		test.ShouldBeTrue(t, v.Valid, "jwt token is NOT valid")
	}
	testToken(tokenPairs.Token)
	testToken(tokenPairs.RefreshToken)
}

func TestGetTokenFromHeaderAndVerify(t *testing.T) {

}
