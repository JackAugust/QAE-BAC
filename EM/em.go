// EHR Management
package EM

// func insertCoop(casenumber string, coop []string) error {
// 	DB := sql.InitDB()
// 	for _, organization := range coop {
// 		sqlStr2 := "insert into insti_coop(_CaseNumber,insti_name)values(?,?)"
// 		_, err2 := DB.Exec(sqlStr2, casenumber, organization)
// 		if err2 != nil {
// 			fmt.Println("插入合作机构失败:", err2.Error())
// 			return err2
// 		}
// 	}
// 	return nil
// }

// // 增 数据上传
// func UploadEHR(uid string, subjectMark string, coop []string) (string, error) {
// 	DB := sql.InitDB()
// 	var _Organization, _Diseases, user_role string
// 	row := DB.QueryRow("select user_insti, user_desease,user_role from user where uid='" + uid + "'")
// 	row.Scan(&_Organization, &_Diseases, &user_role)
// 	if _Organization == "" {
// 		fmt.Println("没有该用户: ", uid)
// 		return "", fmt.Errorf("没有该用户")
// 	}
// 	// if user_role == "u3" {
// 	// 	fmt.Println("该用户为u3: ", uid)
// 	// 	return "", fmt.Errorf("该用户为u3，无上传权限！")
// 	// }
// 	_SubjectMark := subjectMark
// 	_Researcher := uid
// 	_Groups := _Diseases
// 	_GatherTime := time.Now().Format("2006-01-02 15:04:05")
// 	_CaseNumber := GetCaseNumber(_Groups, _SubjectMark, _Diseases, _Researcher, _Organization)

// 	sqlStr1 := `insert into base_info(_SubjectMark,_Researcher,_Organization,_Diseases,_CaseNumber,_GatherTime,_Groups) values(?,?,?,?,?,?,?)`
// 	_, err1 := DB.Exec(sqlStr1, _SubjectMark, _Researcher, _Organization, _Diseases, _CaseNumber, _GatherTime, _Groups)
// 	if err1 != nil {
// 		return "", fmt.Errorf("数据库插入不成功！")
// 	} else {
// 		// insertCoop(_CaseNumber, coop)
// 		// 自动生成访问策略
// 		PM.UploadPolicy(_CaseNumber, coop)
// 	}
// 	// fmt.Println(_SubjectMark, "上传成功")

// 	return _CaseNumber, nil
// }
