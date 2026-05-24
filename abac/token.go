package abac

// import (
// 	"algorithm/mycode/RSA"
// 	"algorithm/mycode/sql"
// 	"crypto/sha256"
// 	SQL "database/sql"
// 	"encoding/json"
// 	"fmt"
// )

// // 访问令牌
// type MyToken struct {
// 	index     string
// 	OID       string `json:"oid"` //仅用写ID
// 	Op        string `json:"op"`  //操作，如果是域外则只能为read
// 	Time      int64  `json:"time"`
// 	Delegate  bool   // 是否发生了委派，默认为false
// 	Owner     []byte //监督者的签名
// 	Recipient []byte //接收者的签名
// 	// Sign string `json:"sign"`
// }

// func (token *MyToken) GetID() string {
// 	return token.index
// }
// func (token *MyToken) ToBytes() []byte {
// 	tokenAsBytes, err := json.Marshal(token)
// 	if err != nil {
// 		fmt.Println("token字节化出错")
// 		return nil
// 	}
// 	return tokenAsBytes
// }

// var DB = sql.InitDB()

// // 结构化令牌
// func paraToken(str string) (*MyToken, error) {
// 	var token *MyToken
// 	err := json.Unmarshal([]byte(str), &token)
// 	if err != nil {
// 		fmt.Println("参数化失败: ", err.Error())
// 		return nil, err
// 	}
// 	return token, nil
// }

// // 令牌字节化
// func TokenBytes(token *MyToken) ([]byte, error) {
// 	token_bytes, err := json.Marshal(*token)
// 	return token_bytes, err
// }

// // 验证令牌有效性
// func CheckToken(token *MyToken, req *ABACRequest) (bool, error) {
// 	// fmt.Println(req.Sub.UID)
// 	// fmt.Println(token.Owner)
// 	// fmt.Println(req.Sub.UID)
// 	if !RSA.RSAUsePublicKeyVerify([]byte(req.Sub.UID), token.Owner, req.Sub.UID) {
// 		return false, fmt.Errorf("令牌签名不匹配")
// 	}
// 	if token.Time < req.CurTime {
// 		DestroyToken(GetTokenID(req.Sub.UID, req.Obj, req.Op))
// 		return false, fmt.Errorf("超出有效期")
// 	}
// 	//发生了委派
// 	if token.Delegate {
// 		if !RSA.RSAUsePublicKeyVerify([]byte(req.Recipient), token.Recipient, req.Recipient) {
// 			return false, fmt.Errorf("接收者签名不匹配")
// 		}
// 		// 委派结束，立刻销毁
// 		// 仅销毁委派令牌？试试
// 		// fmt.Println("faseheng weipai ")
// 		// DestroyToken(GetTokenID(req.Sub.UID, req.Obj, req.Op))
// 	}
// 	return true, nil
// }

// func NewMyToken(req *ABACRequest) *MyToken {
// 	return &MyToken{
// 		index:    GetTokenID(req.Sub.UID, req.Obj, req.Op),
// 		OID:      req.Obj,
// 		Op:       req.Op,
// 		Time:     int64(req.CurTime + 86400), //增加一天的有效期
// 		Delegate: false,
// 		Owner:    RSA.RSAUsePrivateKeySign([]byte(req.Sub.UID), req.Sub.UID),
// 	}

// }

// // 增加token
// func CreateToken(req *ABACRequest) *MyToken {
// 	// fmt.Println(req)
// 	token := NewMyToken(req)
// 	tokenID := token.GetID()
// 	// DB := sql.InitDB()
// 	var str string
// 	// casenumber := "sjk"
// 	// fmt.Println("tokenID:", tokenID)
// 	SQLString := "select * from token where token_id='" + tokenID + "'"
// 	// SQLString := "select * from token where tokenID='" + tokenID + "'"
// 	err1 := DB.QueryRow(SQLString).Scan(&str)
// 	if err1 != SQL.ErrNoRows { //没有结果
// 		// fmt.Println("查询失败 is ", err1.Error())
// 		return nil
// 	}

// 	SQLString3 := "insert into token(token_id,content)values(?,?)"
// 	token_bytes, err2 := TokenBytes(token)
// 	if err2 != nil {
// 		fmt.Println("字节化失败：", err2.Error())
// 	}
// 	// fmt.Println("token is ", token)
// 	_, err := DB.Exec(SQLString3, tokenID, token_bytes)
// 	if err != nil {
// 		// fmt.Println("插入失败：", err.Error())
// 		return nil
// 	}
// 	// fmt.Println("jkandkj")
// 	// fmt.Println("aaaa", token.Owner)

// 	return token
// }

// func GetTokenID(uid, oid, op string) string {
// 	return fmt.Sprintf("%x", sha256.Sum256([]byte(uid+oid+op)))
// }

// // 查找token
// func GetToken(req *ABACRequest) (*MyToken, error) {
// 	// 公钥加密(uid,oid,op), 私钥解密==>如果给B了，就先解密后在用对方公钥加密==>然后加入担保人的信息
// 	tokenID := GetTokenID(req.Sub.UID, req.Obj, req.Op)
// 	// fmt.Println("token id is ", tokenID)

// 	// DB := sql.InitDB()
// 	SQLString := "select content from token where token_id='" + tokenID + "'"
// 	// SQLString := "select content from token where token_id='" + tokenID + "'"
// 	var str string
// 	err1 := DB.QueryRow(SQLString).Scan(&str)
// 	if err1 != nil {
// 		// fmt.Println("数据库查询失败：", err1)
// 		return nil, err1
// 	}

// 	token, err2 := paraToken(str)
// 	if err2 != nil {
// 		// fmt.Println("Error is ", err2.Error())
// 		return nil, err2
// 	}

// 	return token, nil
// }

// // 销毁令牌
// func DestroyToken(tokenID string) error {

// 	// DB := sql.InitDB()
// 	SQLString := "delete from token where token_id='" + tokenID + "'"
// 	_, err := DB.Exec(SQLString)
// 	if err != nil {
// 		// fmt.Println("删除失败: " + err.Error())
// 		return err
// 	}
// 	return nil
// }

// // 令牌的委派
// // A-->B
// func DelegateToken(token *MyToken, owner string, recipient string) (*MyToken, error) {
// 	// fmt.Println(owner)
// 	// fmt.Println(token)
// 	// fmt.Println("reci is ", recipient)

// 	if !RSA.RSAUsePublicKeyVerify([]byte(owner), token.Owner, owner) {
// 		return token, fmt.Errorf("签名不匹配")
// 	}
// 	token.Recipient = RSA.RSAUsePrivateKeySign([]byte(recipient), recipient)
// 	token.Delegate = true
// 	return token, nil
// }
