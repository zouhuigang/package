package main

import (
	"fmt"
	"github.com/zouhuigang/package/zqueue"
)

func main() {
	taskcount := 10000                                                //任务数
	solidcount := 5                                                   //多少个士兵去执行
	zqueue.NewTask(taskcount, solidcount, callbackFunction, "a", "b") //参数可省略
}

//回调的函数,args回调的参数
func callbackFunction(args ...interface{}) {
	fmt.Println("list done", args[0], args[1])
}
