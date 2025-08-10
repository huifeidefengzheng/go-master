package main

import (
	"gorm.io/driver/mysql"
	"log"

	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/weblog?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("无法连接数据库:", err)
	}

	log.Println("数据库连接成功")

	// 自动迁移模型
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("数据表迁移失败:", err)
	}

	log.Println("数据表创建成功")

	// 启动HTTP服务器
	err = setupRoutes(db)
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
