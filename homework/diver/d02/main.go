package d02

import (
	"fmt"
	"gorm.io/gorm"
)

type Account struct {
	Id      uint
	Balance float64
}

type Transaction struct {
	Id            uint
	FromAccountId uint
	Amount        float64
}

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// TransferMoney 实现从账户A向账户B转账的事务操作
func TransferMoney(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	// 开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 检查转出账户余额是否足够
	var fromBalance float64
	err := tx.Raw("SELECT balance FROM accounts WHERE id = ?", fromAccountID).Scan(&fromBalance).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if fromBalance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足")
	}

	// 扣除转出账户余额
	err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 增加转入账户余额
	err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccountID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 记录交易信息
	transaction := Transaction{
		FromAccountId: fromAccountID,
		Amount:        amount,
	}
	err = tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{}, &Employee{})

	//db.Create(&Account{Id: 1, Balance: 100})
	//db.Create(&Account{Id: 2, Balance: 100})
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	// 假设账户1向账户2转账100元
	err := TransferMoney(db, 1, 2, 100.0)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功")
	}
}
