package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (p *Person) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(p.Name + ", ")
	if p.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(p.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}

func (p *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte(" Person has three fileds."))
	} else {
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

func (p *Person) GoString() string {
	return "&Person{Name is " + p.Name + ", Age is " + strconv.Itoa(p.Age) + ", Sex is " + strconv.Itoa(p.Sex) + "}"
}

func main() {
	p := &Person{"polaris", 28, 0}
	fmt.Println(p)
	// fmt.Printf("%L", p)
	fmt.Printf("%#v", p)

	var (
		name string
		age  int
	)
	n, _ := fmt.Sscan("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscan("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
}
