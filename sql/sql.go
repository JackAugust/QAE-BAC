package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mytest"
)

// Db数据库连接池
var DB *sql.DB

func InitDB() *sql.DB {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := userName + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbName + "?allowNativePasswords=true"
	// fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	// fmt.Println(DB)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10000)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10000)
	DB.SetMaxOpenConns(1000)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		fmt.Println(err)
	}
	// fmt.Println("connnect success")
	return DB
}

//创建数据表（测试成功）（表名“Student”）
func CreateTable(db *sql.DB, sqlString string) {
	// sqlString := "create table if not exists Student(id BIGINT(20) NOT NULL AUTO_INCREMENT,name varchar(20),age int,primary key(id))"
	res, err := db.Exec(sqlString)
	if err != nil {
		fmt.Print("create table failed\n")
	}

	fmt.Print("creare table succeed\n")
	log.Printf("%v", res)

}
