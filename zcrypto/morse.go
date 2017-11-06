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

var (
	morseITU = map[string]string{
		"A":  ".-",
		"B":  "-...",
		"C":  "-.-.",
		"D":  "-..",
		"E":  ".",
		"F":  "..-.",
		"G":  "--.",
		"H":  "....",
		"I":  "..",
		"J":  ".---",
		"K":  "-.-",
		"L":  ".-..",
		"M":  "--",
		"N":  "-.",
		"O":  "---",
		"P":  ".--.",
		"Q":  "--.-",
		"R":  ".-.",
		"S":  "...",
		"T":  "-",
		"U":  "..-",
		"V":  "...-",
		"W":  ".--",
		"X":  "-..-",
		"Y":  "-.--",
		"Z":  "--..",
		"ä":  ".-.-",
		"ö":  "---.",
		"ü":  "..--",
		"Ch": "----",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		"?":  "..--..",
		"!":  "..--.",
		":":  "---...",
		`\`:  ".-..-.",
		"'":  ".----.",
		"=":  "-...-",
		"-":  "-....-",
		"/":  "-..-.",
		"(":  "-.--.",
		")":  "-.--.-",
		"+":  ".-.-.",
		"×":  "-..-",
		"@":  ".--.-.",
		" ":  "/",
	}
)
