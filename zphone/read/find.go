/*
数据更新，爬取http://www.ip138.com/数据
*/
package read

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	CMCC               byte = iota + 0x01 //中国移动
	CUCC                                  //中国联通
	CTCC                                  //中国电信
	CTCC_v                                //电信虚拟运营商
	CUCC_v                                //联通虚拟运营商
	CMCC_v                                //移动虚拟运营商
	INT_LEN            = 4
	CHAR_LEN           = 1
	HEAD_LENGTH        = 8
	PHONE_INDEX_LENGTH = 9
	PHONE_DAT          = "phone.dat"
)

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

var (
	content     []byte
	CardTypemap = map[byte]string{
		CMCC:   "中国移动",      //1
		CUCC:   "中国联通",      //2
		CTCC:   "中国电信",      //3
		CTCC_v: "中国电信虚拟运营商", //4
		CUCC_v: "中国联通虚拟运营商", //5
		CMCC_v: "中国移动虚拟运营商", //6
	}
	total_len, firstoffset int32
)

type PhoneST struct{}

var M_Phone = PhoneST{}

//初始化，加载dat文件,phoneDat为文件绝对路径
func LoadDatFile(phoneDat string) PhoneST {
	var err error
	content, err = ioutil.ReadFile(phoneDat)
	if err != nil {
		panic(err)
	}
	total_len = int32(len(content))
	firstoffset = M_Phone.get4(content[INT_LEN : INT_LEN*2])

	total := M_Phone.totalRecord()
	ver := M_Phone.Version2()
	fmt.Printf("LOAD PHONE DATA SUCCESS!! SIZE:%d,TOTAL:%d,VERSION:%s\n", total_len, total, ver)
	return M_Phone
}

//得到当前路径
func (this PhoneST) getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (this PhoneST) Debug() {
	fmt.Println(this.version())
	fmt.Println(this.totalRecord())
	fmt.Println(this.firstRecordOffset())
}

func (pr PhoneRecord) String() string {
	return fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
}

/*
|符号，为二进制，按位或运算符,先把a,b转化成二进制,然后位相或,有1出1,无1出0
简单的来说，a|b，就是把a,b 2个数转换为二进制，然后相同位数，有1取1,2个都没有1就取0

*/
func (this PhoneST) get4(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

//将手机号前7位转化为uint32类型
func (this PhoneST) getN(s string) (uint32, error) {
	var n, cutoff, maxVal uint32
	i := 0
	base := 10
	cutoff = (1<<32-1)/10 + 1
	maxVal = 1<<uint(32) - 1
	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
			//fmt.Println("d大于等于0，小于等于9:",v,d,s[1])
			//fmt.Println("0的ascii码:",'0',"9的ascii吗为：",'9')
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10 //如果为a-z，则v=10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10 //如果为A-Z，则v=10
		default:
			return 0, errors.New("当前字符有汉字或其他不符合字符")
		}
		if v >= byte(base) {
			return 0, errors.New("invalid syntax")
		}

		if n >= cutoff {
			// n*base overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n *= uint32(base)

		n1 := n + uint32(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n = n1
	}
	return n, nil
}

func (this PhoneST) version() string {
	return string(content[0:INT_LEN])
}

//返回版本号
func (this PhoneST) Version2() string {
	return string(content[0:INT_LEN]) //1701为2017年01月
}

//返回偏移量+0,里面存储的是索引开始的位置
func (this PhoneST) FirstOffset() string {
	return string(content[INT_LEN : INT_LEN*2])
}

//返回文本长度
func (this PhoneST) Length() (int32, int32) {
	t_len := (int32(len(content)) - this.firstRecordOffset()) / PHONE_INDEX_LENGTH
	//t_len := this.get4(content[INT_LEN : INT_LEN*2])
	firstoffset = this.get4(content[INT_LEN : INT_LEN*2])
	return t_len, firstoffset
}

//返回切片
func (this PhoneST) ContentMinMax(min int32, max int32) string {
	return string(content[min:max])
}

func (this PhoneST) totalRecord() int32 {
	return (int32(len(content)) - this.firstRecordOffset()) / PHONE_INDEX_LENGTH
}

func (this PhoneST) firstRecordOffset() int32 {
	return this.get4(content[INT_LEN : INT_LEN*2])
}

// 二分法查询phone数据
func (this PhoneST) Find(phone_num string) (pr *PhoneRecord, err error) {
	if len(phone_num) < 7 || len(phone_num) > 11 {
		return nil, errors.New("illegal phone length")
	}

	var left int32
	phone_seven_int, err := this.getN(phone_num[0:7])
	if err != nil {
		return nil, errors.New("illegal phone number")
	}
	phone_seven_int32 := int32(phone_seven_int)
	right := (total_len - firstoffset) / PHONE_INDEX_LENGTH
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		offset := firstoffset + mid*PHONE_INDEX_LENGTH //索引区，中间位置，相对于文件的原始位置
		if offset >= total_len {
			break
		}
		cur_phone := this.get4(content[offset : offset+INT_LEN])               //1851657 [9 65 28 0]
		record_offset := this.get4(content[offset+INT_LEN : offset+INT_LEN*2]) //315  [59 1 0 0]
		card_type := content[offset+INT_LEN*2 : offset+INT_LEN*2+CHAR_LEN][0]  //2
		//fmt.Println("当前手机号：", cur_phone, record_offset, card_type, content[offset:offset+INT_LEN], content[offset+INT_LEN:offset+INT_LEN*2])
		switch {
		case cur_phone > phone_seven_int32:
			right = mid - 1
		case cur_phone < phone_seven_int32:
			left = mid + 1
		default:
			cbyte := content[record_offset:]
			end_offset := int32(bytes.Index(cbyte, []byte("\000")))
			data := bytes.Split(cbyte[:end_offset], []byte("|"))
			card_str, ok := CardTypemap[card_type]
			//fmt.Println("卡类型：", card_type, card_str)
			if !ok {
				card_str = "未知电信运营商"
			}
			pr = &PhoneRecord{
				PhoneNum: phone_num,
				Province: string(data[0]),
				City:     string(data[1]),
				ZipCode:  string(data[2]),
				AreaZone: string(data[3]),
				CardType: card_str,
			}
			return
		}
	}
	return nil, errors.New("phone's data not found")
}
