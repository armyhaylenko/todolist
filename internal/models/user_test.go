package models

import (
	"reflect"
	"testing"
)

func TestFromMap(t *testing.T) {
	m := map[string]string{
		"id":       "123",
		"email":    "foo@example.com",
		"password": "\x71\x77\x65\x72\x74\x79\x31\x32\x33\x34",
	}
	u := &User{}
	_ = u.FromMap(m)
	want := User{"123", "foo@example.com", []byte{113, 119, 101, 114, 116, 121, 49, 50, 51, 52}}
	if u.ID != want.ID {
		t.Errorf("ID got: %v, want: %v", u.ID, want.ID)
	}
	if u.Email != want.Email {
		t.Errorf("Email got: %v, want: %v", u.Email, want.Email)
	}
	if !reflect.DeepEqual(u.PasswordHash, want.PasswordHash) {
		t.Errorf("PasswordHash got: %v, want: %v", u.PasswordHash, want.PasswordHash)
	}
}
