package abac

import (
	"algorithm/mycode/model"
	"fmt"
)

// var TreePool = NewTreePools(4000)

// func CalPolicy(req model.ABACRequest, attr_list []string) (*MyTree, []interface{}, string, string, error) {
// 	casenumber := req.Obj
// 	// var attr_list = []string{"org", "role", "researce", "action"}
// 	// 根据attr_list进行排序
// 	var tra []interface{}
// 	for _, v := range attr_list {
// 		if v == "org" {
// 			tra = append(tra, req.Sub.Org)
// 		} else if v == "role" {
// 			tra = append(tra, req.Sub.Role)
// 		} else if v == "researce" {
// 			tra = append(tra, req.Sub.Desease)
// 		} else if v == "action" {
// 			tra = append(tra, req.Op)
// 		}
// 	}

// 	ptree, creatTime, endTime, err := TreePool.Get(casenumber)
// 	// fmt.Println("casenumber is ", casenumber)
// 	// fmt.Println(*TreePool)
// 	// fmt.Println("creatTime: ", creatTime, " endTime: ", endTime, err)
// 	if err != nil {

// 		policy_str, err0 := PM.GetPolicy(casenumber)
// 		if err0 != nil {
// 			// fmt.Println("该策略不存在")
// 			return nil, nil, "", "", err0
// 		}
// 		var policy PM.Policy
// 		err := json.Unmarshal([]byte(policy_str), &policy)
// 		if err != nil {
// 			// fmt.Println("Policy参数化失败")
// 			return nil, nil, "", "", fmt.Errorf("Policy调取失败")
// 		}
// 		// fmt.Println("policy is ", policy)
// 		creatTime = policy.Env.CreatedTime
// 		endTime = policy.Env.EndTime
// 		// 开始进行权限匹配
// 		// fmt.Println(creatTime, endTime)

// 		//3 check sub的属性，这里使用树形结构判断
// 		ptree = PolicyToTree(&policy, attr_list)
// 		// 加入到treepool
// 		TreePool.Push(ptree, casenumber, creatTime, endTime)
// 		// // // stree := m.SubRuleToTree(&user)
// 	}
// 	// fmt.Println("===", creatTime, endTime)

// 	return ptree, tra, creatTime, endTime, nil
// }

// AC
func CalTreeABAC(ptree *MyTree, tra []interface{}) (bool, *model.MyToken, error) {

	if !ptree.Search(tra) {
		// fmt.Println("----")
		return false, nil, fmt.Errorf("权限不足")
	}
	// fmt.Println(tra, " true")
	// return true, CreateToken(req), nil
	return true, nil, nil
}

// //加入信息熵——判断是否可以获得权限
// func CalTreeABACold(ptree *MyTree, tra []interface{}, creatTime string, endTime string, req *model.ABACRequest) (bool, *MyToken, error) {
// 	// fmt.Println("CalTreeABAC")

// 	// TODO:这里单独拎出来了，感觉可以放进去，可以再写一个函数对比
// 	// 1. Check有效期
// 	CreateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", creatTime, time.Local)
// 	EndTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
// 	// fmt.Println(creatTime, endTime)
// 	if err != nil {
// 		fmt.Println("时间标准化时报，请检查: ", err.Error(), "+++ ", endTime)
// 		return false, nil, fmt.Errorf("有效期不合法，请检查: " + err.Error())
// 	}
// 	if req.CurTime > EndTime.Unix() || req.CurTime < CreateTime.Unix() {
// 		fmt.Println("访问时间失效, 有效时间为：", CreateTime, "截止时间为：", EndTime, "而当前时间为：", req.CurTime)
// 		return false, nil, fmt.Errorf("不在可访问时间范围")
// 	}
// 	// 2. check allowOrg，暂时略过
// 	// TODO：后期应该会改成allowIP之类的 就是判断是否是可以迅速获取该数据的用户，如果是，就直接返回，如果不是，在进行全部的规则判决

// 	// PreorderPrint(ptree.Root)
// 	// fmt.Println(tra)
// 	if !ptree.Search(tra) {
// 		// fmt.Println("----")
// 		return false, nil, fmt.Errorf("权限不足")
// 	}
// 	// fmt.Println("klsdk")
// 	// return true, CreateToken(req), nil
// 	return true, nil, nil
// }
