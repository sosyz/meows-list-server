package crypto

import "testing"

func TestRandString(t *testing.T) {
	token, err := RandString(32)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestPassword(t *testing.T) {
	hash, err := EncPassword("abc123")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(hash))
	if !CheckPassword("abc123", hash) {
		t.Error("password not match")
	}
}
