package PM

// 属性数据
type Date2DB struct {
	Groups       string
	SubjectMark  string
	Diseases     string
	Researcher   string
	Organization string
}

// // 访问策略
// type Policy struct {
// 	Obj      string `json:"obj"`   //仅用写id
// 	Owner    string `json:"owner"` //
// 	Env      Env
// 	SubRules []string `json:"subRules"`
// }

// // 环境
// type Env struct {
// 	AllowOrg    string `json:"allowOrg"` // TODO:这里后期可以改为身份认证过程的东西，比如ip或身份标识符
// 	CreatedTime string `json:"createdTime"`
// 	EndTime     string `json:"endTime"` // 代表有效期
// }

// // Policy 的ID生成
// func (p *Policy) GetID() string {
// 	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.Obj)))
// }

// // Policy转码为json
// func (p *Policy) ToBytes() []byte {
// 	b, err := json.Marshal(*p)
// 	if err != nil {
// 		fmt.Println("Policy转码json字符串错误: ", err.Error())
// 		return nil
// 	}
// 	return b
// }
