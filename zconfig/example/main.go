//zouhuigang 952750120@qq.com
//copyright anooc.com
package main

import (
	. "github.com/zouhuigang/package/zconfig"
	"fmt"
)

func main(){
	lv:=ConfigFile.MustValue("global", "log_level", "DEBUG")
	fmt.Println(lv)
}
