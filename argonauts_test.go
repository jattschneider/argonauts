package argonauts

import (
	"testing"
)

func TestSalt(t *testing.T) {
	salt, err := Salt()
	if err != nil {
		t.Error(err)
	}
	if salt == nil || len(salt) == 0 {
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	salt, err := Salt()
	if err != nil {
		t.Error(err)
	}
	opts := DefaultOptions(salt)

	passwd := "somerandompassword"
	hash := Hash(opts, []byte(passwd))
	if hash == nil || len(hash) == 0 {
		t.Fail()
	}
}

func TestCompare(t *testing.T) {
	salt, err := Salt()
	if err != nil {
		t.Error(err)
	}
	saltString := Sprint(salt)
	t.Log(saltString)

	opts := DefaultOptions(salt)
	passwd := "somerandompassword"
	hash := Hash(opts, []byte(passwd))

	hashString := Sprint(hash)
	h, err := ReadString(hashString)
	if err != nil {
		t.Error(err)
	}
	t.Log(hashString)
	t.Log(InnerString(hashString))

	match, err := Compare(opts, []byte(passwd), h)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Fail()
	}

	match, err = Compare(opts, []byte("someotherpassword"), h)
	if match {
		t.Fail()
	}
}
