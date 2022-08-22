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
	data := []byte("123456")
	hash, err := EncPassword(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(hash))
	if !CheckPassword(data, hash) {
		t.Error("password not match")
	}
}
