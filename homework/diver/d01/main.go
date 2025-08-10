package d01

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	Id    uint
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	// 插入数据
	//db.Create(&Student{Name: "lili", Age: 13, Grade: "三年级"})
	//db.Create(&Student{Name: "xao", Age: 13, Grade: "三年级"})

	// 查询数据 年龄大于18
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	fmt.Println("查询年龄大于18结果：")
	for _, student := range students {
		fmt.Println(student.Name, student.Age, student.Grade)
	}

	// 更新数据 姓名为张三的 年级改为"四年级"
	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//	删除年龄小于15岁的
	db.Where("age < ?", 15).Delete(&Student{})

}
