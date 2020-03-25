package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"os"
)

var Db *gorm.DB

func init() {
	// db user pwd here!!!!!
	// 修改成自己的数据库连接信息
	d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gindemo?charset=utf8&parseTime=True&loc=Asia%2FShanghai")

	if err != nil {
		fmt.Println("database connect error", err)
		os.Exit(0)
	}
	Db = d
	fmt.Println("db ok!")
	Db.SingularTable(true) // 全局设置表名不可以为复数形式。
	Db.LogMode(true)
}
