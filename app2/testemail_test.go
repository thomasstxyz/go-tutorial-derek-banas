package app2

import("testing")

func TestIsEmail(t *testing.T){
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}
	_, err := IsEmail("thomas@aueae.xyr")
	if err != nil {
		t.Error("thomas@aueae.xyr is a valid email")
	}
}