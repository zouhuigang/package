package txt2dat

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	outputFile         *os.File
	txtFile            []byte
	err                error
	first_index_offset int64
	HeaderData         io.Writer
	RecordData         io.Writer
	IndexData          io.Writer
	RecordMap          map[string]int64
)

type M_OutdatST struct{}

var M_Outdat = M_OutdatST{}

const (
	FILE_HEAD_LENGTH = 8 //头部8字节
)

func LoadTxtFile(phoneTxt string) M_OutdatST {
	var err error

	txtFile, err = ioutil.ReadFile(phoneTxt)
	if err != nil {

	}

	return M_Outdat

}

func (this *M_OutdatST) WriteDat(DatFile string, version string) error {

	/*text := []byte(`1300000|2|山东|济南|250000|0531
	1300001|2|江苏|常州|213000|0519
	1300002|2|安徽|巢湖|238000|0565
	1300003|2|四川|宜宾|644000|0831
	1300004|2|四川|自贡|643000|0813
	1300005|2|陕西|西安|710000|029
	1300006|2|江苏|南京|210000|025
	1300007|2|陕西|西安|710000|029
	1300008|2|湖北|武汉|430000|027
	1300009|2|陕西|西安|710000|029`)*/
	//创建Map，用于去重
	RecordMap = make(map[string]int64)
	a_txt := strings.Replace(string(txtFile), "\r\n", "\n", -1)
	//a_txt = strings.Trim(a_txt, "\n") //去除末尾
	arr_txt := strings.Split(a_txt, "\n")

	outputFile, err = os.OpenFile(DatFile, os.O_CREATE, 0600)
	defer outputFile.Close()

	//申明写入的数据
	HeaderData = M_Writer.NewWriter(outputFile, 0)                    //头部，<版本号,4字节>|<第一个索引区偏移量，4字节>
	RecordData = M_Writer.NewWriter(outputFile, FILE_HEAD_LENGTH)     //记录区,第8个字节开始写入,<省份>|<城市>|<邮编>|<长途区号>\000
	first_index_offset = FILE_HEAD_LENGTH + this.RecordTotal(arr_txt) //写入记录区数据，并返回索引区第一个偏移量
	IndexData = M_Writer.NewWriter(outputFile, first_index_offset)    //索引区，记录区之后开始写入,<手机号前7位,4字节><记录区对应数据偏移量，4字节><卡类型，1字节>
	//开始写入数据
	HeaderData.Write([]byte(version)) //1704,2017年04月
	HeaderData.Write(this.Bget4(int32(first_index_offset)))

	for _, value := range arr_txt { //fmt.Printf("arr[%d]=%d \n", index, value)
		if len(value) == 0 { //空行
			continue
		}
		recordByte := []byte(value) //1300000|2|山东|济南|250000|0531
		data := bytes.Split(recordByte, []byte("|"))
		record_string := string(data[2]) + "|" + string(data[3]) + "|" + string(data[4]) + "|" + string(data[5]) + "\000"

		recordOffset, ok := RecordMap[record_string]
		if !ok { //不存在
			//fmt.Println("")
			return errors.New("记录区数据出错，请检查数据...")
		}
		this.HandlerEachIndex(recordOffset, recordByte)

	}

	return nil
}

//记录区数据 <省份>|<城市>|<邮编>|<长途区号>\000
func (this *M_OutdatST) HandleEachRecord(b []byte) int64 {
	//data := bytes.Split(b, []byte("|"))
	//record_string := string(data[2]) + "|" + string(data[3]) + "|" + string(data[4]) + "|" + string(data[5]) + "\000"
	length, _ := RecordData.Write(b)
	return int64(length)
}

//索引区数据
func (this *M_OutdatST) HandlerEachIndex(offset int64, b []byte) {
	//Index_first_offset, _ := outputFile.Seek(0, os.SEEK_CUR) //fmt.Printf("Index_first_offset  is %d,Record_3_offset记录长度是: %d\n", Index_first_offset, len(Record_1))

	//end_offset := int32(bytes.Index(b, []byte("\000")))
	//data := bytes.Split(b[:end_offset], []byte("|"))
	data := bytes.Split(b, []byte("|"))

	phone7, _ := strconv.ParseInt(string(data[0]), 10, 64) //string转int64

	cardType, _ := strconv.ParseInt(string(data[1]), 10, 8)
	//fmt.Println(phone7)            //最后会有一个0
	cardTybeByte := byte(cardType) //byte(2)

	//fmt.Println(data, string(data[0]))
	//outputFile.Write(Index_offset_1)
	//outputFile.Write([]byte{offset})
	echoIndex := make([][]byte, 3)
	echoIndex[0] = this.Bget4(int32(phone7)) //手机号前7位
	echoIndex[1] = this.Bget4(int32(offset)) //索引对应的数据在记录区的偏移量
	echoIndex[2] = []byte{cardTybeByte}      //echoIndex[2] = []byte{1}    //卡类型,1,2,3,4,5,6

	bindex := bytes.Join(echoIndex, []byte(""))
	IndexData.Write(bindex)
	//length1, err1 := IndexData.Write(bindex)
	//fmt.Println(length1, err1)
}

//记录区去重及计算偏移量和记录区总长度
func (this *M_OutdatST) RecordTotal(arr []string) int64 {
	var total int64 = 0
	var rOffset int64 = FILE_HEAD_LENGTH //第一个数据偏移量
	for _, value := range arr {
		if len(value) == 0 { //空行
			continue
		}
		recordByte := []byte(value)
		data := bytes.Split(recordByte, []byte("|"))
		record_string := string(data[2]) + "|" + string(data[3]) + "|" + string(data[4]) + "|" + string(data[5]) + "\000"

		_, ok := RecordMap[record_string]
		if !ok { //不存在
			RecordMap[record_string] = rOffset                     //偏移量,存入map方便后面查询
			length := this.HandleEachRecord([]byte(record_string)) //写入文件中..
			rOffset += length                                      //下一个记录区偏移量
			total += length                                        //计算总长度，这个单位其实可以去掉
		}

	}

	return total
}

//将int类型数据转换为4字节
func (this *M_OutdatST) Bget4(i int32) []byte {

	byteI := byte(i)
	byte2 := byte(i >> 8)
	byte3 := byte(i >> 16)
	byte4 := byte(i >> 24)
	//fmt.Printf("%v (%T) \n", byteI, byteI)
	y := []byte{byteI, byte2, byte3, byte4}
	return y

}
