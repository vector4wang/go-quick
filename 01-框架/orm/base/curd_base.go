package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "xxxxx@tcp()/xxx?charset=utf8mb4")
	if err != nil {
		fmt.Println("connet err:", err)
		panic(err)
	}

	// 增
	// 方式一：
	result, err := db.Exec("INSERT INTO user_info (name, age, birthday) VALUES (?, ?, ?)", "zhangsan", "10", "2022-02-23")
	if err != nil {
		fmt.Println("exec err", err)
	}
	fmt.Println(result)

	// 方式二：
	stmt, err := db.Prepare("INSERT INTO user_info (name, age, birthday) VALUES (?, ?, ?)")
	result2, err := stmt.Exec("zhangsan", 2, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println("exec err", err)
	}
	fmt.Println(result2)

	// 删
	// 方式一：
	result3, err := db.Exec("delete from user_info where id=?", 1)

	fmt.Println(result3)

	// 方式二：
	stmt, err = db.Prepare("delete from user_info where id=?")
	if err != nil {
		fmt.Println("exec err", err)
	}
	result4, err := stmt.Exec(2)
	fmt.Println(result4)

	// 改
	// 方式一：
	result, err = db.Exec("update user_info set name=? where id=?", "lisi", 3)
	fmt.Println(result)

	// 方式二：
	stmt, err = db.Prepare("update user_info set name=? where id=?")

	result, err = stmt.Exec("lisis", 4)
	fmt.Println(result)

	// 查
	// 单条
	var username, jobname, createdAt string
	err = db.QueryRow("select name, age, birthday from user_info where id=?", 4).Scan(&username, &jobname, &createdAt)
	if err != nil {
		fmt.Println("QueryRow error :", err.Error())
	}
	fmt.Println("username: ", username, "jobname: ", jobname, "created_at: ", createdAt)

	// 多条
	rows, err := db.Query("select name, age, birthday from user_info where name=?", "zhangsan")
	if err != nil {
		fmt.Println("QueryRow error :", err.Error())
	}
	// 结构体
	type user_info struct {
		Name     string `json:"name"`
		Age      string `json:"age"`
		Birthday string `json:"birthday"`
	}

	// 初始化
	var user []user_info

	// 判断及遍历
	for rows.Next() {
		var username1, jobname1, createdAt1 string
		if err := rows.Scan(&username1, &jobname1, &createdAt1); err != nil {
			fmt.Println("Query error :", err.Error())
		}
		user = append(user, user_info{Name: username1, Age: jobname1, Birthday: createdAt1})
	}

	fmt.Println(user)
}
