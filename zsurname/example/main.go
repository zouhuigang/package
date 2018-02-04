package main

import (
	"fmt"
	"github.com/zouhuigang/package/zsurname"
)

func main() {

	su1 := "邹欧阳李慧刚樂"
	p, _ := zsurname.New()
	b, str := p.FindSurname(su1)

	fmt.Println(b, str)
}
