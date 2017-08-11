### 任务队列，可用于爬虫等项目


使用说明：

	package main
	
	import (
		"fmt"
		"github.com/zouhuigang/package/zqueue"
	)
	
	func main() {
		taskcount := 10000000000 //任务数
		solidcount := 5          //多少个士兵去执行（线程）
		zqueue.NewTask(taskcount)
		zqueue.Soldiers(solidcount, callbackFunction, "a", "b") //参数可省略
	
	}
	
	//回调的函数,args回调的参数
	func callbackFunction(args ...interface{}) {
		fmt.Println("list done", args[0], args[1])
	}

注意：solidcount越多，执行越快，但是需要注意CPU也会越高。
