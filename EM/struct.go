package EM

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

var returnTimeType string

// 属性数据
type Date2DB struct {
	Groups       string
	SubjectMark  string
	Diseases     string
	Researcher   string
	Organization string
}

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
func GetCaseNumber(data0, data1, data2, data3, data4 string) string {
	attr := Date2DB{
		Groups:       data0,
		SubjectMark:  data1,
		Diseases:     data2,
		Researcher:   data3,
		Organization: data4,
	}
	t := ReturnTime(2) //含时分秒
	//t := "2022-12-18 14:02:35"      //8244d8d6ccc1fac9e9333b3466571efef3678a691e69f3f2a0b972dfaf1091b2
	id := HashSHA256(attr, t)[0:13] //按 六院 数据元 长度为13位
	return id
}
func HashSHA256(a Date2DB, t string) string {
	baseString := []string{a.Groups, a.SubjectMark, a.Diseases, a.Researcher, a.Organization, t}
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
