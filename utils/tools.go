package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
)

/*

文件说明：

该文件内容为开发中需要的函数工具，（即封装好的算法），如加密算法，字符串处理等

*/



const (
	//BASE64字符表,不要有重复
	base64Table        = "<>:;',./?~!@#$CDVWX%^&*ABYZabcghijklmnopqrstuvwxyz01EFGHIJKLMNOPQRSTU2345678(def)_+|{}[]9/"
	hashFunctionHeader = "a.c.community"
	hashFunctionFooter = "09.O25.O20.78"
)

/* 对一个字符串进行不可逆MD5加密 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //需要加密的字符
	return hex.EncodeToString(h.Sum(nil))
}

/* 获取 SHA1 字符串 */
func GetSHA1String(s string) string {
	t := sha1.New()
	t.Write([]byte(s))
	return hex.EncodeToString(t.Sum(nil))
}

/* 获取一个Guid值 */
func GetGuid() string {
	g := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, g); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(g))
}

/*
base64 加、解密
*/

var coder = base64.NewEncoding(base64Table)

/* base64 加密 */
func Base64Encode(s string) string {
	src := []byte(hashFunctionHeader + s + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

/* base64 解密 */
func Base64Decode(s string) (string, error) {
	src := []byte(s)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}




/* 字符截取 */

func SubStrings(s string, pos, length int) string {
	runes := []rune(s)					// 将string对象转为一个字符数组
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
