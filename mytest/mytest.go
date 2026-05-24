package mytest

import (
	"algorithm/mycode/abac"
	"algorithm/mycode/anonymity"
	"algorithm/mycode/dataGen"
	"algorithm/mycode/mytools"
	"fmt"
	"log"
	"strconv"
	"time"
)

var dataNum = 1000
var userNum = 10000
var reqNum = 100000
var attrNum = len(abac.Attr_list)
var policyNum = 80
var attInrNum = abac.AttrinNum

// ///////////////////////////模拟生成数据集///////////////////////////////////////
func Simulate() {
	// abac.Attr_list=[]string{"A", "B", "E", "C", "D", "X", "Y", "O"}
	// Sub_Attr = []string{"A", "B", "C", "D"}
	// Obj_Attr = []string{"X", "Y"}

	// // group 1: 为保证一致性，policy没有重新生成，而是使用group2的，这样保证只有sub数量是变幻的
	// userNum = 5000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 2---baseline
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 3：只生成sub和req，policy也不动
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 4：只改变obj和req，policy等不动
	// userNum = 10000
	// dataNum = 5000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 5：只改变obj和req，policy等不动
	// userNum = 10000
	// dataNum = 15000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 6: 为保证一致性，不再生成req，直接取其中的500k
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 500000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 7: 为保证一致性，在原有1000k的基础上加500k
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 500000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 8:
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D", "E"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 9:
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 10: inValue改变，所有数据集更改
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// group 11: inValue改变，所有数据集更改--
	// 这里对比的是group3，因为6^4=1296===不对，还是2吧
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 6
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 12: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 50
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 13: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 150
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}
	// // group 14: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q"}

	// // group 14: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q"}

	// group 15: 只改变polcy个数，（连带着需要重新生成req）
	userNum = 10000
	dataNum = 10000
	reqNum = 1000000
	abac.Sub_Attr = []string{"A", "B", "C", "D"}
	abac.AttrinNum = 2
	policyNum = 100
	abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q", "R", "S"}

	old_attr_list := abac.Attr_list
	abac.PoolNum = reqNum / 10
	for i := 0; i < 1; i++ {
		abac.Attr_list = old_attr_list

		// // fmt.Println("///////////////////////////模拟生成数据集///////////////////////////////////////")
		// dataGen.CreateData(dataNum, "O1000")
		// dataGen.CreateUser(userNum, "S1000")
		// dataGen.CreatePolicy(2, 2)
		dataGen.CreateRequest(int(0.2*float64(reqNum)), int(0.8*float64(reqNum)))
		// // req := mytools.ReadCSV("./dataset/request.csv")
		// // mytools.WriteCSV("./dataset/request.csv", req[:reqNum])
		// fmt.Println("///////////////////////////数据集生成完毕///////////////////////////////////////")

		// 如果散点图不好，可能是策略的原因
		TestAnonmity()
		// TestCompareAC()

	}

}

// ///////////////////////////模拟评估匿名性指标///////////////////////////////////////

func TestAnonmity() {
	fmt.Println("///////////////////////////模拟评估匿名性指标///////////////////////////////////////")
	sub := mytools.ReadCSV("./dataset/subject.csv")
	req := mytools.ReadCSV("./dataset/request.csv")
	r, s := anonymity.CalSubAnonymity(sub, req)
	// fmt.Println(r)
	reqAno := [][]string{{"dataNum: " + strconv.Itoa(dataNum), "userNum: " + strconv.Itoa(userNum), "reqNum: " + strconv.Itoa(reqNum), "attrNum: " + strconv.Itoa(attrNum), "policyNum: " + strconv.Itoa(policyNum), "attInrNum: " + strconv.Itoa(attInrNum)}}
	reqAno = append(reqAno, mytools.MapToStringSlice(r)...)
	mytools.WriteCSV("result/anonmityReq.csv", reqAno)
	// mytools.AppendCSV("result/anonmityReq.csv", reqAno)

	subAno := [][]string{{"dataNum: " + strconv.Itoa(dataNum), "userNum: " + strconv.Itoa(userNum), "reqNum: " + strconv.Itoa(reqNum), "attrNum: " + strconv.Itoa(attrNum), "policyNum: " + strconv.Itoa(policyNum), "attInrNum: " + strconv.Itoa(attInrNum)}}
	subAno = append(subAno, mytools.MapToStringSlice(s)...)
	mytools.WriteCSV("result/anonmitySub.csv", subAno)
	// mytools.AppendCSV("result/anonmitySub.csv", subAno)

	// 画图，没用上，用了excel画
	data := mytools.Trans(mytools.ReadCSV("./result/anonmitySub.csv"))
	// fmt.Println(data)
	mytools.SavePic(data, "result/anonmity_sub_pro1.png")
	data = mytools.Trans(mytools.ReadCSV("./result/anonmityReq.csv"))
	// fmt.Println(data)
	mytools.SavePic(data, "result/anonmity_req_pro1.png")
}

// ///////////////////////////测试域内访问控制///////////////////////////////////////

func TestCompareAC() {
	fmt.Println("///////////////////////////测试域内访问控制///////////////////////////////////////")

	reqDataSet := mytools.ReadCSV("./dataset/request.csv")

	baseTime, basenum := TestBaseAC(reqDataSet)
	// fmt.Println(e, f)
	dicTime, dicnum := TestDicAC(reqDataSet)
	// fmt.Println(c, d)
	calTime, calnum := TestAnoAC(reqDataSet)
	// fmt.Println(a, b)
	// data:=[]string{}
	// 换算成s
	throughout := []string{fmt.Sprintf("%v", basenum/baseTime*1e9), fmt.Sprintf("%v", dicnum/dicTime*1e9), fmt.Sprintf("%v", calnum/calTime*1e9)}
	delay := []string{fmt.Sprintf("%v", baseTime/basenum), fmt.Sprintf("%v", dicTime/dicnum), fmt.Sprintf("%v", calTime/calnum)}
	online := []string{strconv.Itoa(userNum), strconv.Itoa(dataNum), strconv.Itoa(len(reqDataSet)), strconv.Itoa(len(abac.Sub_Attr)), strconv.Itoa(attInrNum), strconv.Itoa(policyNum), strconv.Itoa(attrNum)}
	online = append(online, throughout...)
	online = append(online, delay...)
	// online = append(online, )
	mytools.AppendCSV("result/introAC.csv", [][]string{online})
}

// attr_list定期更新
// TODO：必须要搭配 dataGen.Attr_list = abac.ChooseBestFeature(dataGen.MyHistoryPool)
func TestAnoAC(reqDataSet [][]string) (float64, float64) {
	// 用来记录req中tra的顺序，因为在这个方法中tra的属性顺序会改变
	old_attr_list := abac.Attr_list
	policy := abac.RestructPolicy(abac.Attr_list)
	sumTime := 0.0
	totalReq := float64(len(reqDataSet))
	for k, v := range reqDataSet {
		// req := mytools.SliceToInterface(v[:6])
		// str := `{"X":"` + v[0] + `,` + sub[1:len(sub)-1] + `,` + obj[1:len(obj)-1] + `}`
		kv := make(map[string]interface{})
		for i, attr := range old_attr_list {
			kv[attr] = v[i]
		}
		// fmt.Println(abac.Attr_list, kv)

		// kv["A"] = v[0]
		// kv["B"] = v[1]
		// kv["C"] = v[2]
		// kv["X"] = v[3]
		// kv["Y"] = v[4]
		// kv["D"] = v[5]
		tra := dataGen.GenTra(kv)
		// fmt.pln
		// 进行权限判断
		start := time.Now()
		result, _, _ := abac.CalTreeABAC(policy, tra)
		elapsed := time.Since(start).Nanoseconds()

		// 判断结果是否和我们所期望的一致，如果不一致则报错
		if r, _ := strconv.ParseBool(v[len(v)-3]); r != result {
			log.Fatalf("TestAnoAC %s和tra%s授权结果不正确: %t", v, tra, result)
		}
		sumTime += float64(elapsed)

		record := abac.NewHistoryRecord(result)
		oneTra := mytools.InterfaceToStringSlice(tra)

		// // TODO:改属性需要改这里：
		record.Attr1 = oneTra[0]
		record.Attr2 = oneTra[1]
		record.Attr3 = oneTra[2]
		record.Attr4 = oneTra[3]
		record.Attr5 = oneTra[4]
		record.Attr6 = oneTra[5]
		record.Attr7 = oneTra[6]
		record.Attr8 = oneTra[7]
		record.Attr9 = oneTra[8]
		record.Attr10 = oneTra[9]
		record.Attr11 = oneTra[10]
		record.Attr12 = oneTra[11]
		record.FeatureNum = len(abac.Attr_list)
		abac.MyHistoryPool.Push(record)
		// fmt.Println(record)
		// 每500个请求作为一个轮次更新属性
		if (k+1)%abac.PoolNum == 0 {
			// fmt.Println(k, abac.Attr_list)
			abac.Attr_list = abac.ChooseBestFeature(abac.MyHistoryPool)
			fmt.Println(k, abac.Attr_list)

			policy = abac.RestructPolicy(abac.Attr_list)

		}
	}
	return sumTime, totalReq
}

// 把attr_list固定，来证明匿名性的作用
func TestDicAC(reqDataSet [][]string) (float64, float64) {
	// attr_list := []string{"B", "D", "X", "C", "Y", "A"}
	// attr_list := []string{"D", "C", "A", "X", "B", "Y"}
	attr_list := abac.Attr_list
	// attr_list := []string{"X", "Y", "A", "B", "C", "D"}
	policy := abac.RestructPolicy(attr_list)
	sumTime := 0.0
	totalReq := float64(len(reqDataSet))

	for _, v := range reqDataSet {
		// req := mytools.SliceToInterface(v[:6])
		kv := make(map[string]interface{})
		for i, attr := range abac.Attr_list {
			kv[attr] = v[i]
		}
		// 这样不行，因为attr可能顺序不对，gentra是根据abac.Attr_list顺序生成的
		// kv["X"] = v[0]
		// kv["Y"] = v[1]
		// kv["A"] = v[2]
		// kv["B"] = v[3]
		// kv["C"] = v[4]
		// kv["D"] = v[5]
		// tra := dataGen.GenTra(kv)
		var tra []interface{}
		for _, v := range attr_list {
			tra = append(tra, kv[v])
		}
		// 进行权限判断
		start := time.Now()
		result, _, _ := abac.CalTreeABAC(policy, tra)
		elapsed := time.Since(start).Nanoseconds()

		// 判断结果是否和我们所期望的一致，如果不一致则报错
		if r, _ := strconv.ParseBool(v[len(v)-3]); r != result {
			log.Fatalf("TestDicAC %s授权结果不正确", v)
		}
		sumTime += float64(elapsed)
	}
	return sumTime, totalReq
}

func TestBaseAC(reqDataSet [][]string) (float64, float64) {
	policies := abac.GetAllPolicies()
	sumTime := 0.0
	totalReq := float64(len(reqDataSet))

	for _, v := range reqDataSet {
		req := make(map[string]interface{})
		for i, attr := range abac.Attr_list {
			req[attr] = v[i]
		}
		start := time.Now()
		result, _, _ := abac.BaseABAC(policies, req)
		elapsed := time.Since(start).Nanoseconds()

		if r, _ := strconv.ParseBool(v[len(v)-3]); r != result {
			fmt.Println(v[len(v)-3])
			log.Fatalf("TestBaseAC %s授权结果不正确", v)
		}
		sumTime += float64(elapsed)
	}

	// policy, err := abac.BasePolicy(request)
	return sumTime, totalReq
}
