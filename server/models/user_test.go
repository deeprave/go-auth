package models

import (
	"github.com/deeprave/go-testutils/test"
	"github.com/pquerna/otp/totp"
	"testing"
	"time"
)

func TestNewCredentialPlain(t *testing.T) {
	cm := NewCredentialManager("Acme, Inc.", "user@example.com")
	secret := "secret"
	c, err := cm.NewCredential(1, Plain, "secret")
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	test.ShouldNotBeNil(t, c, "credential is nil")
	expectedData := map[CredType]string{
		"plain": secret,
	}
	expected := &Credential{
		UserId: 1,
		Type:   "plain",
		Data:   expectedData,
	}
	test.ShouldBeEqual(t, c, expected)
	valid, err := c.Matches(secret)
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	test.ShouldBeTrue(t, valid, "password '%s' does not match", secret)
}

func TestNewCredentialHash(t *testing.T) {
	cm := NewCredentialManager("Acme, Inc.", "user@example.com")
	secret := "secret"
	c, err := cm.NewCredential(1, Hash, secret)
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	test.ShouldNotBeNil(t, c, "credential is nil")
	_, ok := c.Data["hash"]
	test.ShouldBeTrue(t, ok, "expected hash not found in credential")
	valid, err := c.Matches(secret)
	test.ShouldBeNoError(t, err, "unexpected: %v", err)
	test.ShouldBeEqual(t, c.Type, Hash)
	test.ShouldBeTrue(t, valid, "password '%s' does not match", secret)
}

func TestNewCredentialTOTP(t *testing.T) {
	cm := NewCredentialManager("Acme, Inc.", "user@example.com")
	c, err := cm.NewCredential(1, TOTP)
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	test.ShouldNotBeNil(t, c, "c is nil")
	secret, ok := c.Data["totp"]
	test.ShouldBeTrue(t, ok, "expected TOTP secret not found in credential")
	for i := 0; i < len(secret); i += 4 {
		print(secret[i:i+4], " ")
	}
	println()
	// generate a TOTP and check it
	code, err := totp.GenerateCode(secret, time.Now())
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	valid, err := c.Matches(code)
	test.ShouldBeNoError(t, err, "unexpected: %v", err)
	test.ShouldBeEqual(t, c.Type, TOTP)
	test.ShouldBeTrue(t, valid, "TOTP does not match")
}
