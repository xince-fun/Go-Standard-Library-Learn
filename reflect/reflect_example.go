package main

import (
	"fmt"
	"reflect"
)

// 定一个 Enum 类型
type Enum int

const (
	Zero Enum = 0
)

type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

func main() {
	/*
		var a int = 3
		t := reflect.TypeOf(a)          // 一个 reflect.Type
		fmt.Println(t.String())         // int
		fmt.Println(t)                  // int
		fmt.Println(t.Name(), t.Kind()) // int
	*/

	/*
		var w io.Writer = os.Stdout
		fmt.Println(reflect.TypeOf(w)) // *os.File
	*/

	/*
		var a int = 3
		v := reflect.ValueOf(a) // 一个 reflect.Value
		fmt.Println(v)          // 3
		fmt.Println(v.String()) // <int Value>
	*/

	/*
		t := v.Type()           // reflect.Type
		fmt.Println(t.String()) // int
	*/

	/*
		v := reflect.ValueOf(3) // reflect.Value
		x := v.Interface()      // interface{}
		i := x.(int)            // int
		fmt.Printf("%d\n", i)   // "3"
	*/

	/*
		type cat struct{}
		typeOfCat := reflect.TypeOf(cat{})
		fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // cat struct
		typeOfA := reflect.TypeOf(Zero)
		fmt.Println(typeOfA.Name(), typeOfA.Kind()) // Enum int
	*/

	/*
		type cat struct{}
		ins := &cat{}
		typeOfCat := reflect.TypeOf(ins)
		fmt.Printf("name: '%v' kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name: '' kind: 'ptr'
		typeOfCat = typeOfCat.Elem()
		fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // element name: 'cat', element kind: 'struct'
	*/

	/*
		type cat struct {
			Name string
			Type int `json:"type" id:"100"`
		}

		ins := cat{Name: "mini", Type: 1}
		typeOfCat := reflect.TypeOf(ins)
		for i := 0; i < typeOfCat.NumField(); i++ {
			fieldType := typeOfCat.Field(i) // 获取每个成员的结构体字段类型
			fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
		}
		if catType, ok := typeOfCat.FieldByName("Type"); ok {
			fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		}
	*/

	/*
		type cat struct {
			Name string
			Type int `json: "type" id:"100"`
		}
		typeOfCat := reflect.TypeOf(cat{})
		if catType, ok := typeOfCat.FieldByName("Type"); ok {
			fmt.Println(catType.Tag.Get("json"))
		}
	*/

	/*
		var x float64 = 3.4
		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())
		fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
		fmt.Println("value:", v.Float())
	*/

	/*
		var x uint8 = 'x'
		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())                            // uint8.
		fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true
		x = uint8(v.Uint())                                       // v.Uint returns a uint64.
	*/

	/*
		var x float64 = 3.4
		v := reflect.ValueOf(x)
		fmt.Println("settability of v:", v.CanSet())
	*/

	/*
		var x float64 = 3.4
		p := reflect.ValueOf(&x) // Note: take the address of x.
		fmt.Println("type of p: ", p.Type())
		fmt.Println("settability of p: ", p.CanSet())
	*/

	/*
		var x float64 = 3.4
		p := reflect.ValueOf(&x) // Note: take the address of x.
		v := p.Elem()
		fmt.Println("settability of v:", v.CanSet())
	*/

	/*
		var x float64 = 3.4
		p := reflect.ValueOf(&x) // Note: take the address of x.
		v := p.Elem()
		v.SetFloat(7.1)
		fmt.Println(v.Interface())
		fmt.Println(x)
	*/

	/*
		type T struct {
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
	*/

	/*
		type T struct {
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		s.Field(0).SetInt(77)
		s.Field(1).SetString("Sunset Strip")
		fmt.Println("t is now", t)
	*/

	/*
		var a int = 1024
		valueOfA := reflect.ValueOf(a)
		// 获取 interface{}类型的值，再通过断言转换
		var getA int = valueOfA.Interface().(int)
		// 获取64位的值，强转为int类型
		var getA2 int = int(valueOfA.Int())
		fmt.Println(getA, getA2)
	*/

	/*
		d := reflect.ValueOf(dummy{
			next: &dummy{},
		})

		// 获取字段数量
		fmt.Println("NumField", d.NumField())

		// 获取索引为2的字段(float32字段)
		floatField := d.Field(2)
		// 输出字段类型
		fmt.Println("Field", floatField.Type())

		// 根据名字查找字段
		fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())

		// 根据索引查找值， next字段的int字段的值
		fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
	*/

	/*
		// *int的空指针
		var a *int
		fmt.Println("var a *int: ", reflect.ValueOf(a).IsNil())

		// nil值
		fmt.Println("nil: ", reflect.ValueOf(nil).IsValid())

		// *int类型的空指针
		fmt.Println("(*int)(nil): ", reflect.ValueOf((*int)(nil)).Elem().IsValid())

		// 实例化一个结构体
		s := struct{}{}
		// 尝试从结构体中查找一个不存在的字段
		fmt.Println("不存在的结构体成员: ", reflect.ValueOf(s).FieldByName("").IsValid())
		// 尝试从结构体中查找一个不存在的方法
		fmt.Println("不存在的结构体方法: ", reflect.ValueOf(s).MethodByName("").IsValid())

		// 实例化一个map
		m := map[int]int{}
		// 尝试从map中查找一个不存在的键
		fmt.Println("不存在的键: ", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
	*/

	/*
		x := 2                   // value type variable?
		a := reflect.ValueOf(2)  // 2 int no
		b := reflect.ValueOf(x)  // 2 int no
		c := reflect.ValueOf(&x) // &x *int no
		d := c.Elem()            // 2 int yes

		fmt.Println(a.CanAddr()) // false
		fmt.Println(b.CanAddr()) // false
		fmt.Println(c.CanAddr()) // false
		fmt.Println(d.CanAddr()) // true
	*/

	/*
		// 声明整数变量a并赋值
		var a int = 1024

		// 获取变量a的反射值对象(a的地址)
		valueOfA := reflect.ValueOf(&a)

		// 取出a地址的元素(a的值)
		valueOfA = valueOfA.Elem()

		// 修改a的值为1
		valueOfA.SetInt(1)

		fmt.Println(valueOfA.Int())
	*/

	/*
		type dog struct {
			LegCount int
		}

		// 获取dog实例地址的反射对象
		valueOfDog := reflect.ValueOf(&dog{})

		// 取出dog示例地址的元素
		valueOfDog = valueOfDog.Elem()

		// 获取legCount字段的值
		vLegCount := valueOfDog.FieldByName("LegCount")

		// 尝试设置legCount的值为4
		vLegCount.SetInt(4)

		fmt.Println(vLegCount.Int())
	*/

	var a int

	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)

	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)

	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())
}
