package abac

type Date2DB struct {
	Groups       string
	SubjectMark  string
	Diseases     string
	Researcher   string
	Organization string
}
type ActionList struct {
	AdminAction []string `json:"adminAction"`
	U1Action    []string `json:"u1Action"`
	U2Action    []string `json:"u2Action"`
	U3Action    []string `json:"u3Action"`
}

// // 主体
// type Sub struct {
// 	UID     string `json:"uid"`     // 作为某个人的外键，进而可以指到某个具体的人，和那个样本标识符是一个样的 具有唯一性
// 	Role    string `json:"role"`    // role of the subject, 可取值为："Adminstor"，“u1”，“u2”，“u3”
// 	Desease string `json:"Desease"` // 疾病种类，也就算研究方向,eg：精神病 or 心脏病；
// 	Org     string `json:"org"`     //隶属机构, e.g.: 北大六院
// 	Token   MyToken
// }

// // 主体
// type Sub struct {
// 	UID   string
// 	A     string
// 	B     string
// 	C     string
// 	Token MyToken
// }

// // 客体
// type Obj struct {
// 	oid string
// 	X   string
// 	Y   string
// }

// Obj转码为json
// func (p *Obj) ToBytes() []byte {
// 	b, err := json.Marshal(*p)
// 	if err != nil {
// 		fmt.Println("Obj转码json字符串错误: ", err.Error())
// 		return nil
// 	}
// 	return b
// }

// // []byte => obj
// func NewResource(b []byte) (Obj, error) {
// 	r := Obj{}
// 	err := json.Unmarshal(b, &r)
// 	return r, err
// }

// 操作
//  TODO: 目前设计为单操作，所以没有设计结构体

// // 环境
// type Env struct {
// 	AllowOrg    string `json:"allowOrg"` // TODO:这里后期可以改为身份认证过程的东西，比如ip或身份标识符
// 	CreatedTime string `json:"createdTime"`
// 	EndTime     string `json:"endTime"` // 代表有效期
// }

// // 访问策略
// type Policy struct {
// 	Obj      string `json:"obj"`   //仅用写id
// 	Owner    string `json:"owner"` //
// 	Env      Env
// 	SubRules []string `json:"subRules"`
// }

// 请求
// type ABACRequest struct {
// 	Sub     Sub    //仅用写id
// 	Obj     string `json:"obj"` //仅用写id
// 	CurTime int64  `json:"curTime"`
// 	Op      string `json:"op"` //
// }
// type ABACRequest struct {
// 	Type      string `json:"type"` //是否跨域
// 	Obj       string `json:"obj"`  //sid，资源的唯一标识符
// 	Sub       Sub    //这里将用户的查询提前了，所以直接记了sub而不是uAddr
// 	Op        string `json:"op"` //
// 	Delegate  bool
// 	Recipient string `json:"recipient"`
// 	CurTime   int64  `json:"curTime"`
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
