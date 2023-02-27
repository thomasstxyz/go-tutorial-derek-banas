package app2

import(
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	r, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if r.MatchString(s) {
		return "Valid Email", nil
	} else  {
		return "", fmt.Errorf("not a valid email")
	}
}