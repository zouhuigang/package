package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {

	client_id := `1096828699`
	client_secret := `1f66e825d98bc916afc084c59fe3c883`
	var push_time int64 = 1509523926
	str := fmt.Sprintf("%s%d%s", client_id, push_time, client_secret)
	s1 := zcrypto.Sign(str)
	s2 := zcrypto.SignPhpSha1(str)
	fmt.Printf("%s\n%s\n", s1, s2)
}
