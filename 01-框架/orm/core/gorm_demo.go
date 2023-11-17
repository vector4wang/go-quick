package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/**
https://www.cnblogs.com/kylin5201314/p/15484473.html
*/

// 定义数据模型
type UserInfo struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Name       string    `gorm:"size:255" json:"name"`
	Age        uint      `json:age`
	Birthday   time.Time `json:birthday`
	CreateTime time.Time `gorm:"-" json:createTime`
	UpdateTime time.Time `gorm:"-" json:updateTime`
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second, // 慢 SQL 阈值
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // 禁用彩色打印
	},
)

func DbTest() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DC.MysqlConfig.UserName, DC.MysqlConfig.Pwd, DC.MysqlConfig.Url, DC.MysqlConfig.Database)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 账号密码地址端口
	}), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 表名后不加s
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 直接通过上面连接得到的db，调用DB()函数返回即可设置
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println(db)

	// 自动迁移数据库结构
	db.AutoMigrate(&UserInfo{})

	// 创建一条数据
	bd, err := time.Parse(time.DateTime, "2023-11-16 19:00:35")
	fmt.Println(bd)
	db.Create(&UserInfo{Name: "面包123", Age: 3, Birthday: bd})

	// 查询所有数据
	var userInfos []UserInfo
	db.Find(&userInfos)
	fmt.Println("所有用户：", userInfos)
	// 更新一条数据
	db.Model(&UserInfo{}).Where("name = ?", "面包").Update("age", 4)

	// 删除一条数据
	db.Where("name = ?", "面包1").Delete(&UserInfo{})

}
