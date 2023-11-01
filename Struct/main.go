package main

import "fmt"

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

// 此处Stringer是一个接口，它有一个方法String() string
type Stringer interface {
	String() string
}

// 此处使用person结构体实现Stringer接口
// 给结构体person定义了一个方法，这个方法(名称，参数返回值)与接口里方法签名一样
// 这样结构体person就实现了Stringer接口，即Golang中接口是隐式实现的
// 此变量p作为函数printString的参数
func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("THe address province is %s, the address city is %s", addr.province, addr.city)
}

// 实现Stringer接口的函数
// 此函数接收一个Stringer接口类型的参数，打印出Stringer接口String方法返回的字符串
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	p := person{
		age:  23,
		name: "Richard",
		addr: address{
			province: "Guangdong",
			city:     "Shenzhen",
		},
	}
	fmt.Println(p.addr.city)

	printString(p)
	//结构体address也实现了Stringer接口
	printString(p.addr)
}

//面向接口编程: 定义和调用双方满足约定，就可以使用，不用考虑具体实现当
