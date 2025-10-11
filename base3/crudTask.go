package base3

import (
	"gorm.io/gorm"
)

/*题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。*/

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func CrudBase(db *gorm.DB) {
	//db.AutoMigrate(Student{})
	//db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	//db.Create(&Student{Name: "李四", Age: 22, Grade: "四年级"})
	//db.Create(&Student{Name: "王五", Age: 14, Grade: "一年级"})
	/*var students []Student
	db.Where("Age > ?", 18).Find(&students)
	fmt.Println(students)*/

	/*var student Student
	student.Id = 1
	//student.Name = "张三"
	//student.Age = 20
	student.Grade = "四年级"
	//db.Debug().Model(&student).Where("Name = ?", "张三").Update("Grade", "四年级")
	db.Debug().Select("Grade").Save(&student)*/
	//db.Debug().Where("Age < ?", 15).Delete(&Student{})
	db.Debug().Delete(&Student{}, "Age < ?", 21)
}
