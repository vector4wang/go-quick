package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义数据模型
type UserInfo struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Name       string    `gorm:"size:255" json:"name"`
	Age        uint      `json:age`
	Birthday   time.Time `json:birthday`
	CreateTime time.Time `json:createDate`
	UpdateTime time.Time `json:updateDate`
}

func main() {
	// 建立数据库连接

	db, err := gorm.Open("mysql", "xxx@tcp()/xxx?charset=utf8&parseTime=True&loc=Local&serverTimezone=Asia/Shanghai")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 用表名加s方法
	db.SingularTable(true)
	// 自动迁移数据库结构
	db.AutoMigrate(&UserInfo{})

	// 创建一条数据
	bd, err := time.Parse(time.DateTime, "2023-11-16 19:00:35")
	fmt.Println(bd)
	db.Create(&UserInfo{Name: "面包", Age: 3, Birthday: bd})

	// 查询所有数据
	var userInfos []UserInfo
	db.Find(&userInfos)
	fmt.Println("所有用户：", userInfos)
	// 更新一条数据
	db.Model(&UserInfo{}).Where("name = ?", "面包").Update("age", 4)

	// 删除一条数据
	db.Where("name = ?", "面包1").Delete(&UserInfo{})
}
