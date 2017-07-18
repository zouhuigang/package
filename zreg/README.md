### golang验证包，验证手机号，电话号码，生日等



[http://dmdgeeker.com/goBook/docs/ch04/validation.html](http://dmdgeeker.com/goBook/docs/ch04/validation.html)



### 1.判断是否为空

	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zreg"
	)
	
	func main() {
	
		//为空
		s := []string{"", "zou4009@qq.com", " 邹慧刚"}
		for _, v := range s {
			fmt.Printf("是否为空%v : %v\n", v, zreg.IsNull(v))
		}
	
	}

### 输出

	是否为空 : true
	是否为空zou4009@qq.com : false
	是否为空 邹慧刚 : false
