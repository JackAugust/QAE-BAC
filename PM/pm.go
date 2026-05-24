// policy management
package PM

// import (
// 	"algorithm/mycode/sql"
// 	SQL "database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"strings"
// 	"time"
// )

// var DB = sql.InitDB()

// // policy的基础模板
// func InitPolicyCreate(attr Date2DB, CaseNumber string, GatherTime string) string {
// 	// policy模板
// 	template := `{
// 		"CaseNum": {
// 			"role": {
// 				"admin": {
// 					"action": "r&w&d"
// 				},
// 				"u1": {
// 					"researce": "",
// 					"action": "r&w&d",
// 					"org": ""
// 				},
// 				"u2": {
// 					"researce": "",
// 					"action": "r&w",
// 					"org": ""
// 				},
// 				"u3": {
// 					"action": "r",
// 					"org": ""
// 				}
// 			},
// 			"owner": "",
// 			"allowOrg": "",
// 			"time": ""
// 		}

// 	}`
// 	var mpJson map[string]interface{}
// 	json.Unmarshal([]byte(template), &mpJson)

// 	casenumber := CaseNumber // _CaseNumber 数据ID
// 	owner := attr.Researcher //_Researcher 研究者
// 	ThisTime := GatherTime   //_GatherTime 时间戳

// 	org := attr.Organization
// 	Diseases := attr.Diseases //指研究方向or疾病

// 	policy := make(map[string]interface{})
// 	policy[casenumber] = mpJson["CaseNum"]
// 	policy[casenumber].(map[string]interface{})["owner"] = owner
// 	policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})["u1"].(map[string]interface{})["researce"] = Diseases
// 	policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})["u1"].(map[string]interface{})["org"] = org
// 	policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})["u2"].(map[string]interface{})["researce"] = Diseases
// 	policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})["u2"].(map[string]interface{})["org"] = org
// 	policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})["u3"].(map[string]interface{})["org"] = org
// 	// TODO：这里看看后期改成allowIP那些
// 	policy[casenumber].(map[string]interface{})["allowOrg"] = org
// 	policy[casenumber].(map[string]interface{})["time"] = ThisTime

// 	sub := policy[casenumber].(map[string]interface{})["role"].(map[string]interface{})
// 	var subrules []string
// 	for k, v := range sub {
// 		value, _ := json.Marshal(v.(map[string]interface{}))
// 		sv := string(value)
// 		// sv = strings.Replace(sv, "\"", "", -1)
// 		sv = strings.Replace(sv, "\\u0026", "&", -1)
// 		// fmt.Println(" the v is ", sv)
// 		subrules = append(subrules, `"role":"`+k+`",`+sv[1:len(sv)-1])
// 	}

// 	// fmt.Println("subrules length is ", len(subrules))

// 	// 把policy标准化
// 	p := Policy{}
// 	p.Obj = casenumber
// 	p.Owner = owner
// 	p.Env.AllowOrg = org
// 	p.Env.CreatedTime = ThisTime
// 	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", ThisTime, time.Local)
// 	ttemp := tt.AddDate(+1, 0, 0)
// 	p.Env.EndTime = time.Unix(ttemp.Unix(), 0).Format("2006-01-02 15:03:04")
// 	p.SubRules = subrules
// 	// fmt.Println(p.SubRules)
// 	by, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Printf("Map转化为byte数组失败,异常:%s\n", err)
// 		// return err
// 	}
// 	str := string(by)
// 	str = strings.Replace(str, "\\u0026", "&", -1)
// 	// fmt.Println(str)
// 	return (str)
// }

// //增 添加某条数据的访问策略
// func UploadPolicy(casenumber string, coop []string) string {
// 	// DB := sql.InitDB()
// 	// 1. 根据casenum查询出基础policy所需要的属性，生成基本policy
// 	var g, s, d, r, o, t string
// 	rows := DB.QueryRow("select _Groups,_SubjectMark,_Diseases,_Researcher,_Organization,_GatherTime FROM base_info where _CaseNumber='" + casenumber + "'")
// 	rows.Scan(&g, &s, &d, &r, &o, &t)
// 	attr := Date2DB{
// 		Groups:       g,
// 		SubjectMark:  s,
// 		Diseases:     d,
// 		Researcher:   r,
// 		Organization: o,
// 	}
// 	// fmt.Println("attr is ", attr)

// 	// 2. 生成基础策略
// 	policy_str := InitPolicyCreate(Date2DB(attr), casenumber, t)

// 	// 这里进行了更改，因为需要模拟不同机构，所以在初始化访问策略后直接模拟了其他机构的情况
// 	// 3. 加入组织机构
// 	var policy Policy
// 	err := json.Unmarshal([]byte(policy_str), &policy)
// 	if err != nil {
// 		fmt.Println("policy参数化失败")
// 	}
// 	for _, v := range coop {
// 		// var onerule string
// 		u1rule := `"role":"u1","researce":"` + d + `","org":"` + v + `","action":"r&w"`
// 		u2rule := `"role":"u2","researce":"` + d + `","org":"` + v + `","action":"r"`
// 		// u3无权限
// 		policy.SubRules = append(policy.SubRules, u1rule)
// 		policy.SubRules = append(policy.SubRules, u2rule)
// 	}
// 	policy_newbyte, err1 := json.Marshal(policy)
// 	policy_newstr := string(policy_newbyte)
// 	if err1 != nil {
// 		fmt.Println("policy转换字符串失败")
// 	}
// 	policy_newstr = strings.Replace(policy_newstr, "\\u0026", "&", -1)

// 	var str string

// 	// 4. 插入数据库，进行存储
// 	SQLString := "select * from policy where policy_id='" + casenumber + "'"
// 	err2 := DB.QueryRow(SQLString).Scan(&str)
// 	if err2 == SQL.ErrNoRows { //没有结果
// 		SQLString3 := "insert into policy(policy_id,policy_data)values(?,?)"
// 		_, err := DB.Exec(SQLString3, casenumber, policy_newstr)
// 		if err != nil {
// 			fmt.Println("err")
// 		}
// 	}
// 	return policy_newstr
// }

// // 删 从数据库中删除Policy
// func DeletePolicy(casenumber string) error {
// 	// DB := sql.InitDB()
// 	SQLString := "delete from policy where policy_id='" + casenumber + "'"
// 	_, err := DB.Exec(SQLString)
// 	if err != nil {
// 		fmt.Println("删除失败: " + err.Error())
// 		return err
// 	}
// 	return nil
// }

// //查 从数据库中调取Policy
// func GetPolicy(casenumber string) (string, error) {
// 	// DB := sql.InitDB()
// 	var str string
// 	SQLString := "select policy_data from policy where policy_id='" + casenumber + "'"
// 	err := DB.QueryRow(SQLString).Scan(&str)
// 	if err != nil {
// 		// fmt.Println("该策略不存在: " + casenumber + err.Error())
// 		return "", err
// 	}
// 	// fmt.Println("the policy is ", str)
// 	return str, nil
// }
