package abac

import (
	"algorithm/mycode/model"
	"algorithm/mycode/mytools"
	"fmt"
)

// baseline 进行权限判决
// TODO:没有直接判断是owner的情况
func BaseABAC(policies []*model.Policy, req map[string]interface{}) (bool, *model.MyToken, error) {

	// 2. check allowOrg，暂时略过
	// TODO：后期应该会改成allowIP之类的 就是判断是否是可以迅速获取该数据的用户，如果是，就直接返回，如果不是，在进行全部的规则判决

	// 3. check subrules
	// subRules := policy.SubRules
	// user := req.Sub
	// fmt.Println("-----------")
	for _, policy := range policies {
		kvs := mytools.Struct2Map(*policy)
		// fmt.Println(kvs, req)
		num := 0
		// var nilvalue interface{} = ""
		for k, v := range kvs {
			// num++
			if k == "PID" || v == "" {
				num++
				// fmt.Println(k, v, "continue")
				// if num == len(kvs) {
				// 	return true, nil, nil
				// }
				continue
			}
			if req[k] != v {
				// fmt.Println(k, v, reflect.TypeOf(v), "break")
				break
			}
			num++

			// fmt.Println(num, k, v, len(kvs))
		}
		if num == len(kvs) {
			// fmt.Println("zzzz")
			return true, nil, nil
		}
	}

	return false, nil, fmt.Errorf("权限不足")
}
