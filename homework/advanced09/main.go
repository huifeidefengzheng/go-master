package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("员工编号: %d\n", e.EmployeeID)
}

// 也可以为 Person 实现方法
func (p Person) PrintPersonInfo() {
	fmt.Printf("个人信息:\n")
	fmt.Printf("  姓名: %s\n", p.Name)
	fmt.Printf("  年龄: %d\n", p.Age)
}
func main() {

	// 创建 Person 实例
	person := Person{
		Name: "张三",
		Age:  30,
	}

	// 创建 Employee 实例，通过组合 Person
	employee := Employee{
		Person: Person{
			Name: "李四",
			Age:  28,
		},
		EmployeeID: 1001,
	}

	person.PrintPersonInfo()
	employee.PrintInfo()

}
