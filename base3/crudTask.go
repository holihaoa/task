package base3

import (
	"errors"

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

/*题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id
转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。*/

type account struct {
	Id      int
	Balance float64
}

type transaction struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func CrudBase2(db *gorm.DB, accountId int, baccountId int, amount float64) {
	/*db.AutoMigrate(&account{})
	db.AutoMigrate(&transaction{})
	db.Create(&[]account{{Id: 1, Balance: 100}, {Id: 2, Balance: 100}, {Id: 3, Balance: 100}})*/
	db.Transaction(func(tx *gorm.DB) error {
		var a account
		if err := tx.Debug().Where("id = ?", accountId).First(&a).Error; err != nil {
			return err
		}
		if a.Balance < amount {
			return errors.New("余额不足")
		}
		if err := tx.Debug().Model(&account{}).Where("id = ?", accountId).Update("Balance", a.Balance-amount).Error; err != nil {
			return err
		}
		if err := tx.Debug().Model(&account{}).Where("id = ?", baccountId).Update("Balance", gorm.Expr("Balance+?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Create(&transaction{FromAccountId: accountId, ToAccountId: baccountId, Amount: amount}).Error; err != nil {
			return err
		}
		return nil
	})
}
