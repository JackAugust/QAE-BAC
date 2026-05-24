package abac

// import (
// 	"algorithm/mycode/PM"
// 	"encoding/json"
// 	"fmt"
// 	"time"
// )

// // 获取访问策略并对其参数化
// func DicPolicy(req *ABACRequest, attr_list []string) (*MyTree, []interface{}, *PM.Policy, error) {
// 	casenumber := req.Obj
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
// 	policy_str, err0 := PM.GetPolicy(casenumber)
// 	if err0 != nil {
// 		// fmt.Println("该策略不存在")
// 		return nil, nil, nil, err0
// 	}
// 	var policy PM.Policy
// 	err := json.Unmarshal([]byte(policy_str), &policy)
// 	if err != nil {
// 		// fmt.Println("Policy参数化失败")
// 		return nil, nil, nil, fmt.Errorf("Policy调取失败")
// 	}
// 	ptree := PolicyToTree(&policy, attr_list)

// 	return ptree, tra, &policy, nil
// }

// //加入树形结构——判断是否可以获得权限
// func DicTreeABAC(ptree *MyTree, tra []interface{}, policy *PM.Policy, req *ABACRequest) (bool, *MyToken, error) {
// 	// casenumber := req.Obj
// 	// var attr_list = []string{"org", "role", "researce", "action"}
// 	// 根据attr_list进行排序
// 	// var tra []interface{}
// 	// for _, v := range attr_list {
// 	// 	if v == "org" {
// 	// 		tra = append(tra, req.Sub.Org)
// 	// 	} else if v == "role" {
// 	// 		tra = append(tra, req.Sub.Role)
// 	// 	} else if v == "researce" {
// 	// 		tra = append(tra, req.Sub.Desease)
// 	// 	} else if v == "action" {
// 	// 		tra = append(tra, req.Op)
// 	// 	}
// 	// }
// 	// fmt.Println("tra: ", tra)

// 	// ptree, creatTime, endTime, err := TreePool.Get(casenumber)
// 	// if err != nil {

// 	// policy_str, err0 := PM.GetPolicy(casenumber)
// 	// if err0 != nil {
// 	// 	// fmt.Println("该策略不存在")
// 	// 	return false, err0
// 	// }
// 	// var policy PM.Policy
// 	// err := json.Unmarshal([]byte(policy_str), &policy)
// 	// if err != nil {
// 	// 	// fmt.Println("Policy参数化失败")
// 	// 	return false, fmt.Errorf("Policy调取失败")
// 	// }
// 	// fmt.Println(policy)
// 	creatTime := policy.Env.CreatedTime
// 	endTime := policy.Env.EndTime
// 	// 开始进行权限匹配

// 	// TODO:这里单独拎出来了，感觉可以放进去，可以再写一个函数对比
// 	// 1. Check有效期
// 	CreateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", creatTime, time.Local)
// 	EndTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
// 	if err != nil {
// 		// fmt.Println("有效期不合法，请检查")
// 		return false, nil, fmt.Errorf("有效期不合法，请检查")
// 	}
// 	if req.CurTime > EndTime.Unix() || req.CurTime < CreateTime.Unix() {
// 		// fmt.Println("访问时间失效, 有效时间为：", CreateTime, "截止时间为：", EndTime, "而当前时间为：", req.CurTime)
// 		return false, nil, fmt.Errorf("不在可访问时间范围")
// 	}
// 	// 2. check allowOrg，暂时略过
// 	// TODO：后期应该会改成allowIP之类的 就是判断是否是可以迅速获取该数据的用户，如果是，就直接返回，如果不是，在进行全部的规则判决
// 	//3 check sub的属性，这里使用树形结构判断
// 	// ptree := PolicyToTree(&policy, attr_list)
// 	// 加入到treepool
// 	// TreePool.Push(ptree, casenumber, creatTime, endTime)
// 	// // // stree := m.SubRuleToTree(&user)
// 	// }
// 	// PreorderPrint(ptree.Root)
// 	// fmt.Println(tra)
// 	if !ptree.Search(tra) {
// 		return false, nil, fmt.Errorf("权限不足")
// 	}
// 	// req2 := *req
// 	// req2.Obj = "dic-" + req.Obj
// 	// return true, CreateToken(&req2), nil
// 	return true, nil, nil
// }
