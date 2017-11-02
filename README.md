### golang的一些常用的包

### 使用方法：

根据自己需要的包，按照下面的方法，导入自己的项目，即可使用。

	import(
		"github.com/zouhuigang/package/zcore"
		"github.com/zouhuigang/package/ztime"
		"github.com/zouhuigang/package/zreg"
	)

然后获取代码：

	go get github.com/zouhuigang/package/zcore

或者使用第三方包管理工具，例如：gvt等

	添加:
	gvt fetch github.com/zouhuigang/package/zcore
	
	更新:
	gvt update github.com/zouhuigang/package/zcore

	删除：

	gvt delete github.com/zouhuigang/package/zcore

### 功能介绍

zreg包

>验证表单，输入输出,具体看详情页


ztime包

>时间及日历操作,具体看详情页


zphone包

>查找和写入手机归属地包


zprint包

>接入云打印机（易联云k4）


zbucket包

>令牌桶，用来限制流量，限制接口调用次数等


	