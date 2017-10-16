
### 加载数据时，输出的版本信息:
	
2017年1月份更新：LOAD PHONE DATA SUCCESS!! SIZE:3203029,TOTAL:354522,VERSION:1701

	归属地信息库文件大小：3,203,029 字节
	归属地信息库最后更新：2017年1月(1701即17年1月份）
	手机号记录条数：354522


2017年7月份更新:LOAD PHONE DATA SUCCESS!! SIZE:3293083,TOTAL:364528,VERSION:1707

	归属地信息库文件大小：3,293,083 字节
	归属地信息库最后更新：2017年7月
	手机号段记录条数：364528




### 使用

代码例子：

	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zphone/read"
	)
	
	func main() {
		zphoneRead := read.LoadDatFile(`D:\mnt\anooc_go\src\github.com\zouhuigang\package\zphone\read\example\phone.dat`)
		list := []string{`18516573852`, `13247084265`, `15202922979`, `17365793990`, `15248312860`}
	
		for _, v := range list {
			pr, err := zphoneRead.Find(v)
			if err != nil {
				//panic(err)
	
			}
			fmt.Printf("%s,%s,%s,%s,%s,%s\n", v, pr.Province, pr.City, pr.ZipCode, pr.AreaZone, pr.CardType)
		}
	
	}


输出:

	D:\mnt\anooc_go\src\github.com\zouhuigang\package\zphone\read\example>go run main.go
	LOAD PHONE DATA SUCCESS!! SIZE:3203029,TOTAL:354522,VERSION:1701
	18516573852,上海,上海,200000,021,中国联通
	13247084265,江西,南昌,330000,0791,中国联通
	15202922979,陕西,西安,710000,029,中国移动
	17365793990,湖南,怀化,418000,0745,中国电信
	15248312860,内蒙古,阿拉善盟,737300,0483,中国移动