package dataGen

import (
	"algorithm/mycode/abac"
	"algorithm/mycode/mytools"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// func in(target string, str_array []string) bool {
// 	sort.Strings(str_array)
// 	index := sort.SearchStrings(str_array, target)
// 	if index < len(str_array) && str_array[index] == target {
// 		return true
// 	}
// 	return false
// }

// 属性类别：org，role，desease，op

// action属性
var A_list = []string{"r", "w"}

// subject属性
// b、c分别代表一个集合
var B_list = []string{"b1", "b2", "b3", "b4"}
var C_list = []string{"c1", "c2", "c3", "c4"}

// d代表数字(比如年龄)
// var D_list = []string{"d1", "d2"}

// object属性
var X_list = []string{"x1", "x2"}
var Y_list = []string{"y1", "y2"}

// env属性
// 访问时间、地点、设备类型、安全级别等
// var Env=[]string{}
var attrNum = abac.AttrinNum

// 生成subject数据集
func CreateUser(num int, dcID string) {
	fmt.Println("------CreateUser------")
	sujectDataSet := [][]string{}
	subMap := make(map[string]int)
	// 生成num个subject，存入数据集
	for i := 0; i < num; i++ {
		// 生成0～3随机数
		rand.Seed(time.Now().UnixNano())

		// 填充用户数据信息
		uid := dcID + strconv.Itoa(i)
		// uname := uid
		// sub_a := A_list[rand.Intn(len(A_list))]
		oneSubject := []string{uid}

		for _, v := range abac.Sub_Attr {
			oneattr := v + strconv.Itoa(rand.Intn(attrNum))
			oneSubject = append(oneSubject, oneattr)
		}
		// sub_b := B_list[rand.Intn(len(B_list))]
		// sub_c := C_list[rand.Intn(len(C_list))]
		// sub_d := strconv.Itoa(rand.Intn(4))

		sujectDataSet = append(sujectDataSet, oneSubject)
		str := strings.Join(oneSubject[1:], ",")
		subMap[str]++
	}
	mytools.WriteCSV("./dataset/subject.csv", sujectDataSet)
	fmt.Println("------CreateUser完毕------")
	fmt.Println(subMap)

}

// 生成object数据集
func CreateData(num int, dcID string) {
	fmt.Println("------CreateData------")

	objectDataSet := [][]string{}

	// 生成num个object，存入数据集
	for i := 0; i < num; i++ {
		rand.Seed(time.Now().UnixNano())
		oid := dcID + strconv.Itoa(i)
		// obj_X := X_list[rand.Intn(len(X_list))]
		// obj_Y := Y_list[rand.Intn(len(Y_list))]
		oneObject := []string{oid}
		for _, v := range abac.Obj_Attr {
			oneattr := v + strconv.Itoa(rand.Intn(attrNum))
			oneObject = append(oneObject, oneattr)
		}
		objectDataSet = append(objectDataSet, oneObject)
	}
	mytools.WriteCSV("./dataset/object.csv", objectDataSet)
	fmt.Println("------CreateData完毕------")
}

// 生成policy数据集
func CreatePolicy(base int, num int) {
	objects := mytools.ReadCSV("./dataset/object.csv") // 读取object数据集
	objectSet := make(map[string]bool)

	policyDataSet := [][]string{}
	// 使用map来避免重复策略
	policySet := make(map[string]bool)

	for i := 0; i < len(objects); i++ {
		object := objects[i]
		op := A_list[rand.Intn(len(A_list))]
		sli := []string{`"O":"` + op + `"`}
		for k, v := range abac.Obj_Attr {
			sli = append(sli, `"`+v+`":"`+object[k+1]+`"`)
		}
		str := strings.Join(sli, " ")
		// 每个object和对应的操作生成1-3条规则，已生成的不再生成
		if _, exists := objectSet[str]; !exists {
			objectSet[str] = true
			for i := 0; i < base+rand.Intn(num); i++ {

				// 开始生成策略
				rand.Seed(time.Now().UnixNano())

				onePolicy := sli

				// obj_X := object[1]
				// onePolicy = append(onePolicy, `"X":"`+obj_X+`"`)

				// obj_Y := object[2]
				// onePolicy = append(onePolicy, `"Y":"`+obj_Y+`"`)

				for _, v := range abac.Sub_Attr {
					// 0.8的概率进行属性赋值
					if rand.Intn(8)%8 != 0 {
						// sub_b := B_list[rand.Intn(len(B_list))]
						oneattr := v + strconv.Itoa(rand.Intn(attrNum))
						onePolicy = append(onePolicy, `"`+v+`":"`+oneattr+`"`)
					}
				}

				// if rand.Intn(8)%8 != 0 {
				// 	sub_c := C_list[rand.Intn(len(C_list))]
				// 	onePolicy = append(onePolicy, `"C":"`+sub_c+`"`)
				// }
				// if rand.Intn(8)%8 != 0 {
				// 	sub_d1 := strconv.Itoa(rand.Intn(10))
				// 	onePolicy = append(onePolicy, `"D":"`+sub_d1+`"`)
				// }
				// str := ""
				// for _, v := range onePolicy {
				// 	str += v
				// }
				if len(onePolicy) > len(abac.Obj_Attr)+1 {
					pid := mytools.HashSHA256(onePolicy)
					policy := []string{`"PID":"` + pid + `"`}
					policy = append(policy, onePolicy...)
					if _, exists := policySet[pid]; !exists {
						policySet[pid] = true
						policyDataSet = append(policyDataSet, policy)
					}
				}
			}
		}
	}
	mytools.WriteCSV("./dataset/policy.csv", policyDataSet)
}

// 按比例生成req数据集
func CreateRequest(succNum int, failNum int) {
	fmt.Println("------CreateRequest------")

	subjects := mytools.ReadCSV("./dataset/subject.csv") // 读取object数据集
	objects := mytools.ReadCSV("./dataset/object.csv")   // 读取object数据集
	// policy := RestructPolicy(Attr_list)
	policy := abac.RestructPolicy(abac.Attr_list)
	// fmt.Println(policy)
	// abac.PreorderPrint(policy.Root)
	// fmt.Println(succNum, failNum)
	requestDataSet := [][]string{}
	for a, b := 0, 0; a < succNum || b < failNum; {
		// oneReq := CreateOneRequest(subjects, objects)
		// kv := kv(*oneReq)
		req := createOneReq(subjects, objects)
		tra := GenTra(req)
		if len(tra) < 4 {
			continue
		}
		// AC(oneReq)
		// fmt.Println(tra)
		result, _, _ := abac.CalTreeABAC(policy, tra)
		if result && a < succNum {
			a++
			oneTra := mytools.InterfaceToStringSlice(tra)
			oneTra = append(oneTra, fmt.Sprintf("%t", result))
			oneTra = append(oneTra, req["UID"].(string))
			oneTra = append(oneTra, req["OID"].(string))
			// fmt.Println(a, oneTra)
			requestDataSet = append(requestDataSet, oneTra)
			if rand.Intn(5)%5 == 0 {
				a++
				requestDataSet = append(requestDataSet, oneTra)
			}
		} else if !result && b < failNum {
			b++
			oneTra := mytools.InterfaceToStringSlice(tra)
			oneTra = append(oneTra, fmt.Sprintf("%t", result))
			oneTra = append(oneTra, req["UID"].(string))
			oneTra = append(oneTra, req["OID"].(string))
			// fmt.Println(b, oneTra)
			requestDataSet = append(requestDataSet, oneTra)
			if rand.Intn(5)%5 == 0 {
				b++
				requestDataSet = append(requestDataSet, oneTra)

			}
		}
		fmt.Println(a, b)

	}
	rand.Shuffle(len(requestDataSet), func(i, j int) {
		requestDataSet[i], requestDataSet[j] = requestDataSet[j], requestDataSet[i]
	})

	mytools.WriteCSV("./dataset/request.csv", requestDataSet)
	// mytools.AppendCSV("./dataset/request.csv", requestDataSet)
	fmt.Println("------CreateRequest完毕------")

}

// TODO:每次改属性需要该这里和model
// 废弃：0105
// func CreateOneRequest(subjectSet [][]string, objectSet [][]string) *model.ABACRequest {
// 	rand.Seed(time.Now().UnixNano())

// 	var data model.Obj
// 	object := objectSet[rand.Intn(len(objectSet))]
// 	data.OID = object[0]
// 	// data.X = "x2"
// 	// data.Y = "y2"
// 	// 0.8的概率进行属性赋值
// 	data.X = object[1]
// 	data.Y = object[2]

// 	// subreq
// 	subject := subjectSet[rand.Intn(len(subjectSet))]
// 	var user model.Sub
// 	user.UID = subject[0]
// 	// user.B = "b4"
// 	// user.C = "c3"
// 	// user.D = "5"

// 	if rand.Intn(8)%8 != 0 {
// 		user.B = subject[1]
// 	}
// 	if rand.Intn(8)%8 != 0 {
// 		user.C = subject[2]
// 	}
// 	if rand.Intn(8)%8 != 0 {
// 		user.D = subject[3]
// 	}

// 	op := A_list[rand.Intn(len(A_list))]
// 	// 确定此条访问请求格式
// 	var request model.ABACRequest
// 	request.Sub = user
// 	request.Obj = data
// 	request.Op = op
// 	// request.Delegate = false
// 	// request.Recipient = ""
// 	// request.CurTime = time.Now().Unix()
// 	// oldRequest = &request
// 	return &request
// }

// func kv(oneReq model.ABACRequest) map[string]interface{} {
// 	sub := string(oneReq.Sub.ToBytes())
// 	obj := string(oneReq.Obj.ToBytes())
// 	op := `"Op":"` + oneReq.Op + `"`

// 	str := `{` + op + `,` + sub[1:len(sub)-1] + `,` + obj[1:len(obj)-1] + `}`
// 	kv, err := mytools.ToKV(str)
// 	// fmt.Println(str, kv)

// 	if err != nil {
// 		log.Fatalf("GenTra: Policy转化kv失败: %v", err)
// 		return nil
// 	}
// 	return kv
// }

func createOneReq(subjectSet [][]string, objectSet [][]string) map[string]interface{} {
	rand.Seed(time.Now().UnixNano())
	kv := make(map[string]interface{})
	// op := `"O":"` + A_list[rand.Intn(len(A_list))] + `"`
	kv["O"] = A_list[rand.Intn(len(A_list))]
	subject := subjectSet[rand.Intn(len(subjectSet))]
	kv["UID"] = subject[0]
	// sub := string(oneReq.Sub.ToBytes())
	for k, v := range abac.Sub_Attr {
		// var key interface{} = nil
		kv[v] = nil
		// 0.8的概率进行属性赋值
		if rand.Intn(8)%8 != 0 {
			// sub_b := B_list[rand.Intn(len(B_list))]
			// oneattr := v + strconv.Itoa(rand.Intn(attrNum))
			kv[v] = subject[k+1]
			// onePolicy = append(onePolicy, `"`+v+`":"`+oneattr+`"`)
		}
	}
	object := objectSet[rand.Intn(len(objectSet))]
	kv["OID"] = object[0]
	for k, v := range abac.Obj_Attr {
		// sub_b := B_list[rand.Intn(len(B_list))]
		// oneattr := v + strconv.Itoa(rand.Intn(attrNum))
		kv[v] = object[k+1]
		// onePolicy = append(onePolicy, `"`+v+`":"`+oneattr+`"`)
	}
	// fmt.Println(kv, abac.Attr_list)

	// for _, v := range abac.Attr_list {
	// 	if _, exist := kv[v]; !exist {
	// 		kv[v] = strconv.Itoa(rand.Intn(abac.AttrinNum))
	// 	}
	// }

	return kv
}
func GenTra(kv map[string]interface{}) []interface{} {

	var tra []interface{}
	for _, v := range abac.Attr_list {
		if kv[v] == nil {
			tra = append(tra, "")

		} else {
			tra = append(tra, kv[v])

		}
	}
	// fmt.Println(kv, tra, abac.Attr_list)
	return tra
}
