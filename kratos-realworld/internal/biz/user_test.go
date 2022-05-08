package biz

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "abc"
	s := hashPassword(password)
	t.Log(s)
}

func TestVerifyPassword(t *testing.T) {
	hashed := "$2a$10$NpBmlloqSDU8c6WzPGG7EOoDfwO/LlSMS1x36h7Dv2yB.QRQR5VwG"
	input1 := "abc"
	input2 := "def"
	assert.Equal(t, verifyPassword(hashed, input1), true)
	assert.Equal(t, verifyPassword(hashed, input2), false)
}
