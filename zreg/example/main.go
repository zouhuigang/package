package main

import (
	"fmt"
	"github.com/zouhuigang/zreg"
)

func main() {
	s := []string{"18505921256", "13489594009", "12759029321", "18504620000", "15936388868", "15936388868dsdaada34", "18505921256", "13489594009", "18759029321"}
	for _, v := range s {
		fmt.Printf("手机号%v : %v\n", v, zreg.Is_phone(v))
	}

	//验证邮箱
	s = []string{"18505921256", "zou4009@qq.com", "12759029321@.", "@.com18504620000", "15936388868@1634.com", "15936388868dsdaada34", "18505921256", "13489594009", "18759029321"}
	for _, v := range s {
		fmt.Printf("邮箱%v : %v\n", v, zreg.Is_email(v))
	}

	//验证邮箱
	s = []string{"18505921256", "zou4009@qq.com", "邹慧刚", "邹慧刚123abc", "zouhuigang"}
	for _, v := range s {
		fmt.Printf("中文%v : %v\n", v, zreg.Is_chinese(v))
	}

}
