package zip

//https://github.com/rchunping/ip2location-qqwry/blob/master/src/qrcode/main.go
import (
	"encoding/binary"
	//"fmt"
	"github.com/zouhuigang/mahonia"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	INDEX_LEN       = 7
	REDIRECT_MODE_1 = 0x01
	REDIRECT_MODE_2 = 0x02
)

type QQwry struct {
	Ip       string
	Country  string
	Area     string
	filepath string
	file     *os.File
}

func NewQQwry(file string) (qqwry *QQwry) {
	qqwry = &QQwry{filepath: file}
	return
}

//判断ip地址是否正确
func VerifyIp(ip string) bool {
	ipArr := strings.Split(ip, ".")

	if len(ipArr) < 4 {
		return false
	}

	for i := 0; i < len(ipArr); i++ {
		ip_i, err := strconv.Atoi(ipArr[i])
		if err != nil {
			return false
		}

		if ip_i < 0 || ip_i > 255 {
			return false
		}

	}

	return true

}

func (this *QQwry) Find(ip string) {

	/*if !VerifyIp(ip) {
		return
	}*/

	if this.filepath == "" {
		return
	}

	file, err := os.OpenFile(this.filepath, os.O_RDONLY, 0400)
	defer file.Close()
	if err != nil {
		return
	}
	this.file = file

	this.Ip = ip
	offset := this.searchIndex(binary.BigEndian.Uint32(net.ParseIP(ip).To4()))
	if offset <= 0 {
		return
	}

	var country []byte
	var area []byte

	mode := this.readMode(offset + 4)
	// log.Println("mode", mode)
	if mode == REDIRECT_MODE_1 {
		countryOffset := this.readUInt24()
		mode = this.readMode(countryOffset)
		// log.Println("1 - mode", mode)
		if mode == REDIRECT_MODE_2 {
			c := this.readUInt24()
			country = this.readString(c)
			countryOffset += 4
		} else {
			country = this.readString(countryOffset)
			countryOffset += uint32(len(country) + 1)
		}
		area = this.readArea(countryOffset)
	} else if mode == REDIRECT_MODE_2 {
		countryOffset := this.readUInt24()
		country = this.readString(countryOffset)
		area = this.readArea(offset + 8)
	} else {
		country = this.readString(offset + 4)
		area = this.readArea(offset + uint32(5+len(country)))
	}

	enc := mahonia.NewDecoder("gbk")
	this.Country = enc.ConvertString(string(country))
	this.Area = enc.ConvertString(string(area))

}

func (this *QQwry) readMode(offset uint32) byte {
	this.file.Seek(int64(offset), 0)
	mode := make([]byte, 1)
	this.file.Read(mode)
	return mode[0]
}

func (this *QQwry) readArea(offset uint32) []byte {
	mode := this.readMode(offset)
	if mode == REDIRECT_MODE_1 || mode == REDIRECT_MODE_2 {
		areaOffset := this.readUInt24()
		if areaOffset == 0 {
			return []byte("")
		} else {
			return this.readString(areaOffset)
		}
	} else {
		return this.readString(offset)
	}
	return []byte("")
}

func (this *QQwry) readString(offset uint32) []byte {
	this.file.Seek(int64(offset), 0)
	data := make([]byte, 0, 30)
	buf := make([]byte, 1)
	for {
		this.file.Read(buf)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

func (this *QQwry) searchIndex(ip uint32) uint32 {
	header := make([]byte, 8)
	this.file.Seek(0, 0)
	this.file.Read(header)

	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])

	// log.Printf("len info %v, %v ---- %v, %v", start, end, hex.EncodeToString(header[:4]), hex.EncodeToString(header[4:]))

	for {
		mid := this.getMiddleOffset(start, end)
		this.file.Seek(int64(mid), 0)
		buf := make([]byte, INDEX_LEN)
		this.file.Read(buf)
		_ip := binary.LittleEndian.Uint32(buf[:4])

		// log.Printf(">> %v, %v, %v -- %v", start, mid, end, hex.EncodeToString(buf[:4]))

		if end-start == INDEX_LEN {
			offset := byte3ToUInt32(buf[4:])
			this.file.Read(buf)
			if ip < binary.LittleEndian.Uint32(buf[:4]) {
				return offset
			} else {
				return 0
			}
		}

		// 找到的比较大，向前移
		if _ip > ip {
			end = mid
		} else if _ip < ip { // 找到的比较小，向后移
			start = mid
		} else if _ip == ip {
			return byte3ToUInt32(buf[4:])
		}

	}
	return 0
}

func (this *QQwry) readUInt24() uint32 {
	buf := make([]byte, 3)
	this.file.Read(buf)
	return byte3ToUInt32(buf)
}

func (this *QQwry) getMiddleOffset(start uint32, end uint32) uint32 {
	records := ((end - start) / INDEX_LEN) >> 1
	return start + records*INDEX_LEN
}

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
