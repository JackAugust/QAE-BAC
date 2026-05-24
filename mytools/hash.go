package mytools

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

// 定义返回的时间格式
var returnTimeType string

func ReturnTime(t int32) string {
	if t == 1 {
		returnTimeType = "2006-01-02"
	} else if t == 2 {
		returnTimeType = "2006-01-02 15:03:04"
	}
	now := time.Now()
	seconds := time.Unix(now.Unix(), 0)
	timeString := seconds.Format(returnTimeType)
	// 确定输出格式为string
	// fmt.Printf("%T\n", timeString)
	// 输出 时间戳，完整时间，最后返回内容
	// fmt.Println(now.Unix(), seconds, timeString)
	return timeString
}
func HashSHA256(baseString []string) string {
	// baseString := []string{a.Groups, a.SubjectMark, a.Diseases, a.Researcher, a.Organization, t}
	//fmt.Println(baseString)
	str := strings.Join(baseString, "")
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(str))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	//fmt.Println(hashCode)
	return hashCode
}
