package model

import (
	"encoding/json"
	"fmt"
	"log"
)

// 主体
type Sub struct {
	UID string
	A   string
	B   string
	C   string
	D   string
	// Token MyToken
}

func (s *Sub) ToBytes() []byte {
	b, err := json.Marshal(*s)
	if err != nil {
		fmt.Println("Policy转码json字符串错误: ", err.Error())
		return nil
	}
	return b
}

// 客体
type Obj struct {
	OID string
	X   string
	Y   string
}

func (o *Obj) ToBytes() []byte {
	b, err := json.Marshal(*o)
	if err != nil {
		fmt.Println("Policy转码json字符串错误: ", err.Error())
		return nil
	}
	return b
}

// request
type ABACRequest struct {
	Type string //是否跨域
	Obj  Obj    //sid，资源的唯一标识符
	Sub  Sub    //这里将用户的查询提前了，所以直接记了sub而不是uAddr
	Op   string //
	// Delegate  bool
	// Recipient string
	// CurTime   int64
}

// 访问策略
type Policy struct {
	PID string
	// subject
	A string
	B string
	C string
	D string
	// object
	X string
	Y string
	P string
	Q string
	R string
	S string
	// op
	O string
	// env
	E string
}

// Policy转码为json
func (p *Policy) ToBytes() []byte {
	b, err := json.Marshal(*p)
	if err != nil {
		fmt.Println("Policy转码json字符串错误: ", err.Error())
		return nil
	}
	return b
}

func ToPolicy(str string) (*Policy, error) {
	var policy Policy
	err := json.Unmarshal([]byte(str), &policy)
	if err != nil {
		log.Fatalf("Policy参数化失败: %v", err)
		return nil, err
	}
	return &policy, nil
}

// // Policy 的ID生成
// func (p *Policy) GetID() string {
// 	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.Obj)))
// }
// 访问令牌
type MyToken struct {
	index     string
	OID       string `json:"oid"` //仅用写ID
	Op        string `json:"op"`  //操作，如果是域外则只能为read
	Time      int64  `json:"time"`
	Delegate  bool   // 是否发生了委派，默认为false
	Owner     []byte //监督者的签名
	Recipient []byte //接收者的签名
	// Sign string `json:"sign"`
}
