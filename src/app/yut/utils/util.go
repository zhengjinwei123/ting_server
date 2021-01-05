package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

var (
	md5Hash = md5.New()
)

func LoadXmlConfig(filename string, v interface{}) error {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		if err = xml.Unmarshal(contents, v); err != nil {
			return err
		}
		return nil
	}
}

func LoadJsonConfig(filename string, v interface{}) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, v)
	if err != nil {
		return err
	}
	return nil
}

func MD5(str string) string {
	md5Hash.Reset()
	md5Hash.Write([]byte(str))
	return hex.EncodeToString(md5Hash.Sum(nil))
}

func Password(password string) string {
	md5_sign := MD5(password)
	return md5_sign
}

func CheckSign(data string, secret_key string, sign string) bool {
	sign_str := fmt.Sprintf("%s#%s", data, secret_key)

	md5_sign := MD5(sign_str)

	return md5_sign == sign
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func Addslashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}
