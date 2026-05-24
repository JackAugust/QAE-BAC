// user management
package UM

import (
	"algorithm/mycode/RSA"
	"algorithm/mycode/sql"
	"fmt"
)

func RegisterUser(uid, uname, user_role, user_insti, user_desease string) error {
	DB := sql.InitDB()
	sqlStr1 := `insert into user(uid,uname,user_role,user_insti,user_desease) values(?,?,?,?,?)`
	_, err0 := DB.Exec(sqlStr1, uid, uname, user_role, user_insti, user_desease)
	if err0 != nil {
		fmt.Println(err0.Error())
		return err0
	}
	err := RSA.GenerateRSAKey(uid)
	if err != nil {
		fmt.Println("create key error: ", err)
		return err
	}
	return nil
}
