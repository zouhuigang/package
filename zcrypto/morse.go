/*
摩斯密码加密解密
*/

package zcrypto

import (
	"fmt"
	"strings"
)

//通过映射字母表来解码摩斯密码
func morseDecode(s string, alphabet map[string]string, letterSeparator string, wordSeparator string) (string, error) {

	res := ""

	for _, part := range strings.Split(s, letterSeparator) {
		found := false
		for key, val := range alphabet {
			if val == part {
				res += key
				found = true
				break
			}
		}
		if part == wordSeparator {
			res += " "
			found = true
		}
		if found == false {
			return res, fmt.Errorf("unknown character " + part)
		}
	}
	return res, nil
}

// Encode encodes clear text in `s` using `alphabet` mapping
func morseEncode(s string, alphabet map[string]string, letterSeparator string, wordSeparator string) string {

	res := ""

	for _, part := range s {
		p := string(part)
		if p == " " {
			if wordSeparator != "" {
				res += wordSeparator + letterSeparator
			}
		} else if morseITU[p] != "" {
			res += morseITU[p] + letterSeparator
		}
	}
	return strings.TrimSpace(res)
}

// DecodeITU translates international morse code (ITU) to text
func MorseDecodeITU(s string) (string, error) {
	return morseDecode(s, morseITU, " ", "/")
}

// EncodeITU translates text to international morse code (ITU)
func MorseEncodeITU(s string) string {
	return morseEncode(s, morseITU, " ", "/")
}

// LooksLikeMorse returns true if string seems to be a morse encoded string
func LooksLikeMorse(s string) bool {

	if len(s) < 1 {
		return false
	}
	for _, b := range s {
		if b != '-' && b != '.' && b != ' ' {
			return false
		}
	}
	return true
}

//官方摩斯电码，.代表0，-代表1,
var (
	morseITU = map[string]string{
		"A":  "01",     //.-
		"B":  "1000",   //-...
		"C":  "1010",   //-.-.
		"D":  "100",    //-..
		"E":  "0",      //.
		"F":  "0010",   //..-.
		"G":  "110",    //--.
		"H":  "0000",   //....
		"I":  "00",     //..
		"J":  "0111",   //.---
		"K":  "101",    //-.-
		"L":  "0100",   //.-..
		"M":  "11",     //--
		"N":  "10",     //-.
		"O":  "111",    //---
		"P":  "0110",   //.--.
		"Q":  "1101",   //--.-
		"R":  "010",    //.-.
		"S":  "000",    //...
		"T":  "1",      //-
		"U":  "001",    //..-
		"V":  "0001",   //...-
		"W":  "011",    //.--
		"X":  "1001",   //-..-
		"Y":  "1011",   //-.--
		"Z":  "1100",   //--..
		"ä":  "0101",   //.-.-
		"ö":  "1110",   //---.
		"ü":  "0011",   //..--
		"Ch": "1111",   //----
		"0":  "11111",  //-----
		"1":  "01111",  //.----
		"2":  "00111",  //..---
		"3":  "00011",  //...--
		"4":  "00001",  //....-
		"5":  "00000",  //.....
		"6":  "10000",  //-....
		"7":  "11000",  //--...
		"8":  "11100",  //---..
		"9":  "11110",  //----.
		".":  "010101", //.-.-.-
		",":  "110011", //--..--
		"?":  "001100", //..--..
		"!":  "00110",  //..--.
		":":  "111000", //---...
		`\`:  "10010",  //.-..-.
		"'":  "011110", //.----.
		"=":  "10001",  //-...-
		"-":  "100001", //-....-
		"/":  "10010",  //-..-.
		"(":  "10110",  //-.--.
		")":  "101101", //-.--.-
		"+":  "01010",  //.-.-.
		"×":  "1001",   //-..-
		"@":  "011010", //.--.-.
		`"`:  "010010", /* Quotation mark      */
		" ":  "/",
	}
)
