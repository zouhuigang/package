package zreg

import (
	"fmt"
	"testing"
)

func Test_p(t *testing.T) {
	s1, _ := PraseIdCard("362202199401272524")
	s2, _ := PraseIdCard("362202199201016171")
	s3, _ := PraseIdCard("130503670401001")
	fmt.Printf("%v\n%v\n%v\n", s1, s2, s3)
}
