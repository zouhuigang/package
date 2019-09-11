package zreg

import (
	"fmt"
	"testing"
)

func Test_p(t *testing.T) {
	s1, _ := PraseIdCard("362202199401272524")
	s2, _ := PraseIdCard("362202199201016171")
	s3, _ := PraseIdCard("130503670401001")
	s4, _ := PraseIdCard("13050367b401001")
	fmt.Printf("%v\n%v\n%v\n%v\n", s1, s2, s3, s4)
}

func Test_r(t *testing.T) {
	s1 := Is_usercard("130503670401001")
	fmt.Println(s1)
	s2 := Is_usercard("362202199201016171")
	fmt.Println(s2)
	s3 := Is_usercard("3622021s9201016171")
	fmt.Println(s3)
	s4 := Is_usercard("13050367b401001")
	fmt.Println(s4)
	// s5 := Is_usercard("32483498311103682")
	// fmt.Println(s5)
}

func TestIsRealDate(t *testing.T) {
	fmt.Print(IsRealDate(2020, 2, 3))
}
