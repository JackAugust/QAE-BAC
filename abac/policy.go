package abac

import (
	"algorithm/mycode/model"
	"algorithm/mycode/mytools"
	"fmt"
	"log"
	"strings"
)

// 将所有policy根据attr_list进行重构，整合成一棵树(这样就不用每次查policy的时候进行重构了)
func RestructPolicy(attr_list []string) *MyTree {
	policyDataSet := mytools.ReadCSV("./dataset/policy.csv")
	tree := NewTree()
	// fmt.Println(len(policyDataSet))

	for _, v := range policyDataSet {
		// 每条policy转换成json后变成policy结构
		onePolicy := `{` + strings.Join(v, ",") + `}`
		// fmt.Println(onePolicy)
		// policy, err := model.ToPolicy(onePolicy)
		// if err != nil {
		// 	log.Fatalf("Policy参数化失败: %v", err)
		// }
		//
		kv, err := mytools.ToKV(onePolicy)
		if err != nil {
			fmt.Println("参数化失败: ", onePolicy, err.Error())
		}

		var value []interface{}
		for _, v := range attr_list {
			value = append(value, kv[v])
			// fmt.Println("k is ", k, "个v", value)
		}
		// fmt.Println(value)
		tree.Append(value)
	}
	// PreorderPrint(tree.Root)
	// fmt.Println()
	return tree
}
func GetAllPolicies() []*model.Policy {
	policy := mytools.ReadCSV("./dataset/policy.csv")
	policies := []*model.Policy{}
	for _, v := range policy {
		// onePolicy := `{` + strings.Join(v, ",") + `}`
		one, err := model.ToPolicy(`{` + strings.Join(v, ",") + `}`)
		if err != nil {
			log.Fatalf("TestDicAC 参数化policy失败:%v", err)
		}
		policies = append(policies, one)
	}
	return policies
}
