package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(15)

	if len(s) != 15 {
		t.Error("wrong length random string returned")
	}
}
