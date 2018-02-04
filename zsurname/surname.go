package zsurname

import (
	//"fmt"
	"github.com/zouhuigang/package/zreg"
	"io/ioutil"
	"log"
	"strings"
)

const (
	configFile = "config.txt"
)

type configDat struct {
	//data []byte
	data map[string]byte
}

var sur *configDat = nil

func New() (configDat, error) {
	if sur == nil {
		var err error
		sur, err = loadConfigDat()
		if err != nil {
			log.Fatal("the config file loaded failed!")
			return *sur, err
		}
	}
	return *sur, nil
}

//加载配置
func loadConfigDat() (*configDat, error) {

	p := configDat{}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	//预处理成map数据
	st := string(data)
	s1 := strings.Split(st, "\r\n")
	tempMap := map[string]byte{} // 存放不重复主键
	for _, v := range s1 {
		if v == "" {
			continue
		}
		tempMap[v] = 0

	}

	p.data = tempMap
	return &p, nil
}

/*func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}*/

//分离出姓
func getSurname(realname string, n int) string {
	nameRune := []rune(realname)
	//fmt.Println("截取到的姓", string(nameRune[:n]))
	suname := string(nameRune[:n])
	return suname
}

//查找姓名,至多查找2个汉字
func (p configDat) FindSurname(realname string) (bool, string) {

	if zreg.IsNull(realname) { //如果为空，则返回false
		return false, "姓名不能为空"
	}

	if !zreg.Is_chinese(realname) { //不全为中文
		return false, "含有非中文字符"
	}

	if len(realname) < 2 { //姓名长度至少为2
		return false, "至少为2个汉字"
	}

	//查询单姓
	surnameOne := getSurname(realname, 1)
	_, ok := p.data[surnameOne]
	if ok {
		return true, surnameOne
	}

	//分割姓--首先查询复姓
	surnameTwo := getSurname(realname, 2)
	_, ok = p.data[surnameTwo]
	if ok {
		return true, surnameTwo
	}

	return false, "查找失败"
}
