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
