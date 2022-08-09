package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type userInfo struct {
	id, age int
	name    string
}

var db *sql.DB

func MysqlSolution() (err error) {
	sqlConnStr := "db_game_socia_w:Z2gHj4VDrhUzECC_AJCBGTBG6sbtK8FQ@tcp(10.225.118.70:3306)/db_game_social?charset=utf8mb4&parseTime=True"

	db, err = sql.Open("mysql", sqlConnStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Ping()
	sqlInject("xxx' or 1=1#")
	return nil
}

func sqlInject(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user_info where name='%s'", name)
	fmt.Printf("sql:%s\n", sqlStr)
	var u userInfo
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v \n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
