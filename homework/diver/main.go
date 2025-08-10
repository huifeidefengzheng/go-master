package main

import (
	"github.com/d/d05"
	_ "github.com/go-sql-driver/mysql" // MySQL驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// 如果使用其他数据库，导入相应的驱动
	// _ "github.com/lib/pq" // PostgreSQL
	// _ "github.com/mattn/go-sqlite3" // SQLite
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("failed to connect database")
	}

	////d01.Run(db)
	//d02.Run(db)
	d05.Run(db)

	// 请根据实际情况修改数据库连接信息
	//db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/testdb")
	//if err != nil {
	//	log.Fatal("数据库连接失败:", err)
	//}
	//defer db.Close()
	////d03.Run(db)
	//d04.Run(db)
}
