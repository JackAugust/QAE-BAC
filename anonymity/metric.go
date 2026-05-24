package anonymity

import (
	"algorithm/mycode/abac"
	"algorithm/mycode/mytools"
	"math"
	"strings"
)

// 计算每个sub的匿名性
func CalSubAnonymity(subjectDataSet [][]string, requestDataSet [][]string) (map[string][]float64, map[string][]float64) {
	// 计算所有可发送的属性排列组合
	allReqGroup := subsets(len(subjectDataSet[0]) - 1)
	// allReqGroup := [][]int{{1, 4}}
	// fmt.Println(allReqGroup)

	allReqAnonymity := make(map[string][]float64)
	// 对于每个组合
	z := 0.0
	for _, v := range allReqGroup {
		// req:dataset
		// 生成当前属性中可以的所有取值
		labelMap := GenlabelMap(subjectDataSet, v)
		// fmt.Println("-----", labelMap)
		for req, subset := range labelMap {
			reqSet := GenReqGroup(requestDataSet, req)
			// fmt.Println(req, subset, reqSet)
			allSet := append(subset, reqSet...)

			// TODO：subset的len即为（r，t）匿名中的r，req为t，当r为1时需要进行扩充
			// 但是当subset很大时，r通常不为1
			// r := len(subset)
			t := len(strings.Split(req, " "))
			w := len(reqSet)
			AA := CalReqAnonymity(allSet)
			// AA := CalReqAnonymity(subset)
			// AA := CalReqAnonymity(reqSet)

			reqPro := float64(w)
			for z := 0; z < t; z++ {
				reqPro /= 4.000
			}
			// y := float64(len(requestDataSet))
			z += reqPro
			// fmt.Println(reqPro, w, z)
			allReqAnonymity[req] = []float64{AA, reqPro}
		}
	}
	for _, v := range allReqAnonymity {
		v[1] /= z
	}
	// fmt.Println(z)
	allSubAnonymity := make(map[string][]float64)
	for _, sub := range subjectDataSet {
		subAnonmity := 0.0
		subPro := 0.0
		// 对于每个组合
		for _, v := range allReqGroup {
			var selectedStrings []string
			for _, i := range v {
				selectedStrings = append(selectedStrings, sub[i])
			}
			canReq := strings.Join(selectedStrings, " ")
			subAnonmity += allReqAnonymity[canReq][0] * allReqAnonymity[canReq][1]
			// fmt.Println(sub, canReq, allReqAnonymity[canReq], subAnonmity)
			subPro += allReqAnonymity[canReq][1]
		}
		allSubAnonymity[sub[0]] = []float64{subAnonmity, subPro}
	}
	// fmt.Println(allSubAnonymity)
	return allReqAnonymity, allSubAnonymity
}

// 用来计算所有可以构成的req组合，比如对于四个属性abcd，可以组合出abcd、abc、abd、acd等
func subsets(n int) [][]int {
	result := [][]int{}

	// 共2^n个子集
	for mask := 0; mask < (1 << n); mask++ {
		currentSubset := []int{}

		// 检查每一位是否为1，若为1则选择当前元素
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				currentSubset = append(currentSubset, i+1)
			}
		}
		if len(currentSubset) != 0 {
			result = append(result, currentSubset)
		}
	}
	return result
}

// 用来生成t元组中的labelMap，比如属性AB，（a1，b1）为key，【】string{s1,s2,s3,s1}为value，记录主体出现的次数
func GenlabelMap(dataSet [][]string, index []int) map[string][]string {
	result := make(map[string][]string)
	for _, row := range dataSet {
		var selectedStrings []string
		for _, i := range index {
			selectedStrings = append(selectedStrings, row[i])
		}
		str := strings.Join(selectedStrings, " ")
		result[str] = append(result[str], row[0])
	}
	return result
}

// 根据reqdataset补充完整数据集，数据集大小即为（r-t）匿名中的r，req的大小为t，然后根据这个计算reqGroup的信息熵
func GenReqGroup(reqDataSet [][]string, reqstr string) []string {
	// req := strings.Split(reqstr, " ")
	result := []string{}

	for _, row := range reqDataSet {
		var selectedStrings []string
		for _, v := range abac.Sub_Attr {
			index := mytools.Index(v, abac.Attr_list)
			if row[index] != "" {
				selectedStrings = append(selectedStrings, row[index])
			}
		}
		str := strings.Join(selectedStrings, " ")
		// fmt.Println(row, "+", str, " + ", reqstr)
		// str := strings.Join(row[3:6], " ")
		if str == reqstr {
			// fmt.Println(str, " + ", reqstr, "+", row[len(row)-2])
			result = append(result, row[len(row)-2])
		}
	}
	return result
}

// // 根据reqdataset补充完整数据集，数据集大小即为（r-t）匿名中的r，req的大小为t，然后根据这个计算reqGroup的信息熵
// func GenReqGroup2(reqDataSet [][]string, reqstr string, dataset []string) []string {
// 	// req := strings.Split(reqstr, " ")
// 	result := []string{}

// 	for _, row := range reqDataSet {
// 		var selectedStrings []string
// 		for i := 3; i < 6; i++ {
// 			if row[i] != "" {
// 				selectedStrings = append(selectedStrings, row[i])
// 			}
// 		}
// 		str := strings.Join(selectedStrings, " ")
// 		// str := strings.Join(row[3:6], " ")
// 		if str == reqstr {
// 			fmt.Println(str, " + ", reqstr)

// 			result = append(result, row[len(row)-2])
// 		}
// 	}
// 	result = append(result, dataset...)
// 	return result
// }

// 计算req可以由哪些sub生成，并据此计算出信息熵
func CalReqAnonymity(dataSet []string) float64 {
	total := len(dataSet)

	// 计算每个元组的概率
	labelMap := make(map[string]int)
	for _, v := range dataSet {
		labelMap[v]++
	}

	// 计算整个数据空间的信息熵
	// 信息熵公式
	entropy := 0.0
	for _, count := range labelMap {
		// 计算该元组的概率
		prob := float64(count) / float64(total)
		// 计算信息熵的贡献
		entropy -= prob * math.Log2(prob)
		// fmt.Println(prob, math.Log2(prob))
	}
	return entropy
}
