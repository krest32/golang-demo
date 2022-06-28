package mapper

import (
	"demo-basic/blog/code/config"
	"fmt"
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

//Db数据库连接池
var DB *sql.DB



func GetCon() *sql.DB {

	userName:=config.GetDatasourceUserName()
	password :=config.GetDatasourcePassword()
	ip :=config.GetDatasourceIP()
	port :=config.GetDatasourcePort()
	dbName :=config.GetDatasourceDBName()

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/day7-mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	DB.SetMaxIdleConns(5) //设置最大空闲连接数
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return nil
	}
	fmt.Println("connnect success")
	return DB
}


