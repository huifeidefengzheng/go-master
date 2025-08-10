package d03

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Employee 员工信息结构体
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 创建表Employee
func CreateTableEmployee(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS employees (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        department TEXT NOT NULL,
        salary REAL NOT NULL
    )`
	_, err := db.Exec(query)
	return err
}

// QueryEmployeesByDepartment 查询指定部门的所有员工
func QueryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, query, department)

	if err != nil {
		return nil, fmt.Errorf("查询部门员工信息失败: %v", err)
	}

	return employees, nil
}

// QueryHighestPaidEmployee 查询工资最高的员工
func QueryHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee

	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)

	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工信息失败: %v", err)
	}

	return &employee, nil
}

func Run(db *sqlx.DB) {
	// 初始化多条数据
	//_, err := db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "张三", "IT", 5000.0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	_, err := db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "mm", "IT", 6000.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, err1 := db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "李四", "财务", 3000.0)
	//if err1 != nil {
	//	fmt.Println(err)
	//	return
	//}

	employees, err := QueryEmployeesByDepartment(db, "IT")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("IT部门员工信息:")
	for _, employee := range employees {
		fmt.Printf("员工ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", employee.Id, employee.Name, employee.Department, employee.Salary)
	}
	highestPaidEmployee, err := QueryHighestPaidEmployee(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("工资最高的员工信息:")
	fmt.Printf("员工ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", highestPaidEmployee.Id, highestPaidEmployee.Name, highestPaidEmployee.Department, highestPaidEmployee.Salary)

}
