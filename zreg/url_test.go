package zreg

import (
	"fmt"
	"testing"
)

func TestIsHttpUrl(t *testing.T) {
	s1 := IsHttpUrl("https://c3.helendoctor.com/user/signup/one?inviter=4&role=1")
	s2 := IsHttpUrl("user/signup/one?inviter=4&role=1")
	s3 := IsHttpUrl("https://c3.helendoctor.com/")
	s4 := IsHttpUrl("https://c3.helendoctor.com/")
	fmt.Printf("%v\n%v\n%v\n%v\n", s1, s2, s3, s4)
}
