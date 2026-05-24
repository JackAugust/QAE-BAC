package mytools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// string to map[]interface{}
func ToKV(str string) (map[string]interface{}, error) {
	kv := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &kv)
	if err != nil {
		fmt.Println(str, "参数化失败: ", err.Error())
		return nil, err
	}
	return kv, nil
}

// []string-->[]interface{}
func SliceToInterface(strSlice []string) []interface{} {
	var interfaceSlice []interface{}
	for _, str := range strSlice {
		interfaceSlice = append(interfaceSlice, str) // 将每个字符串添加到 []interface{}
	}
	return interfaceSlice
}

// 将 []interface{} 转换为 []string
func InterfaceToStringSlice(interfaceSlice []interface{}) []string {
	// 使用 slice 进行类型断言
	var stringSlice []string
	for _, v := range interfaceSlice {
		if str, ok := v.(string); ok {
			stringSlice = append(stringSlice, str)
		}
	}
	return stringSlice
}

// 将 map[string]float64 转为[][]string
func MapToStringSlice(data map[string][]float64) [][]string {
	result := [][]string{}
	for k, v := range data {
		result = append(result, []string{k, fmt.Sprintf("%v", v[0]), fmt.Sprintf("%v", v[1])})
	}
	// 排序，按照 key 中空格数量升序
	sort.Slice(result, func(i, j int) bool {
		// 计算 key 中空格的数量
		countSpaces := func(s string) int {
			return strings.Count(s, " ")
		}
		return countSpaces(result[i][0]) < countSpaces(result[j][0])
	})
	return result
}

func Struct2Map(stru interface{}) map[string]interface{} {
	t := reflect.TypeOf(stru)
	v := reflect.ValueOf(stru)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 把csv中的数据[][]string转换成图像中的(x,y)坐标
func Trans(data [][]string) map[float64]float64 {
	result := make(map[float64]float64)
	for _, value := range data {
		if len(value) > 3 {
			continue
		}
		// x:=value[0].stroc
		// a := value[len(value)-1]
		// b := strings.Split(a[1:len(a)-1], " ")
		x, _ := strconv.ParseFloat(value[1], 64)
		y, _ := strconv.ParseFloat(value[2], 64)
		// fmt.Println(value, x, y)
		result[x] = y
	}

	return result
}
