# Reflect - 反射包

Go语言提供了一种机制在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在编译时并不知道这些变量的具体类型，这种机制被称为反射。反射也可以让我们将类型本身作为第一类的值类型处理。

Go语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 Type 和 Value 任意接口值在反射中都可以理解为由 reflect.Type 和 reflect.Value 两部分组成，并且 reflect 包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的 Value 和 Type。

### reflect.Type 类型

Type 表示 Go 语言的一个类型，它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分，比如一个结构的各个字段或者一个函数的各个参数。

Type 是可以比较的， 比如使用 `==` 操作符等，所以它们可以用作为 `map` 的 `key`。如何两个 Type 代表的类型相同，则他们相同。

```go
var a int = 3
t := reflect.TypeOf(a)          // 一个 reflect.Type
fmt.Println(t.String())         // int
fmt.Println(t)                  // int
fmt.Println(t.Name(), t.Kind()) // int 
```

`reflect.TypeOf` 返回一个接口值对应的动态类型，所以它总是返回具体类型（而不是接口类型）。

```go
var w io.Writer = os.Stdout
fmt.Println(reflect.TypeOf(w)) // *os.File
```

### reflect.Value 类型

`reflect` 包的另一个类型就是 `Value`。 `reflect.Value` 可以包含一个任意类型的值。

`reflect.ValueOf` 接口任意的 `interface{}` 并将接口的动态值以 `reflect.Value` 返回。 `reflect.ValueOf` 的返回值也都是具体值，不过 `reflect.Value` 也可以包含一个接口值。

```go
var a int = 3
v := reflect.ValueOf(a) // 一个 reflect.Value
fmt.Println(v)          // 3
fmt.Println(v.String()) // <int Value>
```

调用 `Value` 的 `Type` 方法会把它的类型以 `reflect.Type` 方式返回

```go
t := v.Type()           // reflect.Type
fmt.Println(t.String()) // int
```

`reflect.ValueOf` 的逆操作是 `reflect.Value.Interface` 方法。 它返回一个 `interface{}` 接口值

```45go
v := reflect.ValueOf(3) // reflect.Value
x := v.Interface()      // interface{}
i := x.(int)            // int
fmt.Printf("%d\n", i)   // "3"
```

### 反射种类（Kind）的定义

Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

```go
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)
```

Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

### 从类型对象中获取类型名称和种类

Go语言中的类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串；类型归属的种类（Kind）使用的是 reflect.Type 中的 Kind() 方法，返回 reflect.Kind 类型的常量。

```go
// 定义一个Enum类型
type Enum int
const (
    Zero Enum = 0
)

type cat struct{}
typeOfCat := reflect.TypeOf(cat{})
fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // cat struct
typeOfA := reflect.TypeOf(Zero)
fmt.Println(typeOfA.Name(), typeOfA.Kind()) // Enum int
```

### 指针与指针指向的元素

Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作，代码如下：

```go
type cat struct{}
ins := &cat{}
typeOfCat := reflect.TypeOf(ins)
fmt.Printf("name: '%v' kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) 
typeOfCat = typeOfCat.Elem()
fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
```

输出:
```
name: '' kind: 'ptr'
element name: 'cat', element kind: 'struct'
```

Go语言的反射中对所有指针变量的种类都是 Ptr，但需要注意的是，指针变量的类型名称是空，不是 *cat。

### 使用反射获取结构体的成员类型

任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息。

|                             方法                             |                             说明                             |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
|                  `Field(i int) StructField`                  | 根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机 |
|                       `NumField() int`                       | 返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机 |
|        `FieldByName(name string) (StructField, bool)`        | 根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
|           `FieldByIndex(index []int) StructField`            | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机 |
| `FieldByNameFunc(match func(string) bool) (StructField,bool)` | 根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机 |

```go
type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

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
```

输出:

```
NumField 5
Field float32
FieldByName("b").Type string
FieldByIndex([]int{4, 0}).Type() int
```

#### 结构体字段类型

`reflect.Type` 的 `Field()` 方法返回 `StructField` 结构，这个结构描述结构体的成员信息，通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以通过 `StructField` 的 `Type` 字段进一步获取结构体成员的类型信息。

```go
type StructField struct {
    Name string          // 字段名
    PkgPath string       // 字段路径
    Type      Type       // 字段反射类型对象
    Tag       StructTag  // 字段的结构体标签
    Offset    uintptr    // 字段在结构体中的相对偏移
    Index     []int      // Type.FieldByIndex中的返回的索引值
    Anonymous bool       // 是否为匿名字段
}
```

#### 获取成员反射信息

下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。

```go
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
```

输出:
```
name: Name  tag: ''
name: Type  tag: 'json:"type" id:"100"'
type 100
```

### 结构体标签（Struct Tag）

通过 `reflect.Type` 获取结构体成员信息 `reflect.StructField` 结构中的 `Tag` 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

JSON、BSON 等格式进行序列化及对象关系映射（`Object Relational Mapping`，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。

#### 结构体标签的格式

```go
`key1:"value1" key2:"value2"`
```

结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用一个空格分隔。

#### 从结构体标签中获取值

StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：
- `func (tag StructTag) Get(key string) string`：根据 Tag 中的键获取对应的值，例如`key1:"value1" key2:"value2"`的 Tag 中，可以传入“key1”获得“value1”。
- `func (tag StructTag) Lookup(key string) (value string, ok bool)`：根据 Tag 中的键，查询值是否存在。

#### 结构体标签格式错误导致的问题

编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，示例代码如下：

```go
type cat struct {
    Name string
    Type int `json: "type" id:"100"`
}
typeOfCat := reflect.TypeOf(cat{})
if catType, ok := typeOfCat.FieldByName("Type"); ok {
    fmt.Println(catType.Tag.Get("json"))
}
```

在 json: 和 "type" 之间增加了一个空格，这种写法没有遵守结构体标签的规则，因此无法通过 Tag.Get 获取到正确的 json 对应的值。这个错误在开发中非常容易被疏忽，造成难以察觉的错误。

### 通过反射获取值信息

一个 `reflect.Value` 值的 `CanSet` 方法将返回此 `reflect.Value` 值代表的 Go 值是否可以被修改（可以被赋值）。如果一个 Go 值可以被修改，则我们可以调用对应的 `reflect.Value` 值的 Set 方法来修改此 Go 值。注意：`reflect.ValueOf` 函数直接返回的 `reflect.Value` 值都是不可修改的。

反射不仅可以获取值的类型信息，还可以动态地获取或者设置变量的值。Go语言中使用 `reflect.Value` 获取和设置变量的值。


```go
value := reflect.ValueOf(rawValue)
```

reflect.ValueOf 返回 reflect.Value 类型，包含有 rawValue 的值信息。reflect.Value 与原值间可以通过值包装和值获取互相转化。reflect.Value 是一些反射操作的重要类型，如反射调用函数。


#### 从反射值对象(reflect.Value)中获取值的方法

|          方法           |                             说明                             |
| :---------------------: | :----------------------------------------------------------: |
| `Interface() interface` | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型  |
|      `Int() int64`      |     将值以 int 类型返回，所有有符号整型均可以此方式返回      |
|     `Uint() uint64`     |     将值以 uint 类型返回，所有无符号整型均可以此方式返回     |
|    `Float() float64`    | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
|      `Bool() bool`      |                     将值以 bool 类型返回                     |
|    `Bytes() []bytes`    |               将值以字节数组 []bytes 类型返回                |
|    `String() string`    |                     将值以字符串类型返回                     |


#### 从反射值对象(reflect.Value)中获取值的例子

```go
var a int = 1024
valueOfA := reflect.ValueOf(a)
// 获取 interface{}类型的值，再通过断言转换
var getA int = valueOfA.Interface().(int)
// 获取64位的值，强转为int类型
var getA2 int = int(valueOfA.Int())
fmt.Println(getA, getA2)
```

输出:

```
1024 1024
```


### 判断反射值的空和有效性

|       方法       |                             说明                             |
| :--------------: | :----------------------------------------------------------: |
|  `IsNil() bool`  | 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的`v== nil`操作 |
| `IsValid() bool` | 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。 |



下面的例子将会对各种方式的空指针进行 IsNil() 和 IsValid() 的返回值判定检测。同时对结构体成员及方法查找 map 键值对的返回值进行 IsValid() 判定，参考下面的代码。

```go
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
```

`(*int)(nil)` 的含义是将 nil 转换为 `*int`，也就是`*int` 类型的空指针。此行将 nil 转换为 `*int` 类型，并取指针指向元素。由于 nil 不指向任何元素，`*int` 类型的 nil 也不能指向任何元素，值不是有效的。因此这个反射值使用 `Isvalid()` 判断时返回 false。

`IsNil()` 常被用于判断指针是否为空；`IsValid()` 常被用于判定返回值是否有效。

### 通过反射修改变量的值

Go语言中类似 `x`、`x.f[1]` 和 `*p` 形式的表达式都可以表示变量，但是其它如 `x + 1` 和 `f(2)` 则不是变量。一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新。

对于 `reflect.Value` 也有类似的区别。有一些 `reflect.Value` 是可取地址的；其它一些则不可以。考虑以下的声明语句：

```go
x := 2                   // value type variable?
a := reflect.ValueOf(2)  // 2 int no
b := reflect.ValueOf(x)  // 2 int no
c := reflect.ValueOf(&x) // &x *int no
d := c.Elem()            // 2 int yes
```

其中 `a` 对应的变量则不可取地址。因为 `a` 中的值仅仅是整数 `2` 的拷贝副本。`b` 中的值也同样不可取地址。`c` 中的值还是不可取地址，它只是一个指针 `&x` 的拷贝。实际上，所有通过 `reflect.ValueOf(x)` 返回的 `reflect.Value` 都是不可取地址的。但是对于 `d`，它是 `c` 的解引用方式生成的，指向另一个变量，因此是可取地址的。我们可以通过调用 `reflect.ValueOf(&x).Elem()`，来获取任意变量`x`对应的可取地址的 `Value`。

我们可以通过调用 `reflect.Value` 的 `CanAddr` 方法来判断其是否可以被取地址：

```
fmt.Println(a.CanAddr()) // false
fmt.Println(b.CanAddr()) // false
fmt.Println(c.CanAddr()) // false
fmt.Println(d.CanAddr()) // true
```

每当我们通过指针间接地获取的 `reflect.Value` 都是可取地址的，即使开始的是一个不可取地址的 `Value`。在反射机制中，所有关于是否支持取地址的规则都是类似的。例如，`slice` 的索引表达式 `e[i]` 将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。

以此类推，`reflect.ValueOf(e).Index(i)` 对于的值也是可取地址的，即使原始的 `reflect.ValueOf(e)` 不支持也没有关系。

使用 `reflect.Value` 对包装的值进行修改时，需要遵循一些规则。如果没有按照规则进行代码设计和编写，轻则无法修改对象值，重则程序在运行时会发生宕机。

**使用 reflect.Value 取元素、取地址及修改值的属性**

|       方法       |                             说明                             |
| :--------------: | :----------------------------------------------------------: |
|  `Elem() Value`  | 取值指向的元素值，类似于语言层`*`操作。当值类型不是指针或接口时发生宕 机，空指针时返回 nil 的 Value |
|  `Addr() Value`  | 对可寻址的值返回其地址，类似于语言层`&`操作。当值不可寻址时发生宕机 |
| `CanAddr() bool` |                       表示值是否可寻址                       |
| `CanSet() bool`  |         返回值能否被修改。要求值可寻址且是导出的字段         |

**使用 reflect.Value 修改值的相关方法**

|     Set(x Value)      |                将值设置为传入的反射值对象的值                |
| :-------------------: | :----------------------------------------------------------: |
|   `SetInt(x int64)`   | 使用 int64 设置值。当值的类型不是 int、int8、int16、 int32、int64 时会发生宕机 |
|  `SetUint(x uint64)`  | 使用 uint64 设置值。当值的类型不是 uint、uint8、uint16、uint32、uint64 时会发生宕机 |
| `SetFloat(x float64)` | 使用 float64 设置值。当值的类型不是 float32、float64 时会发生宕机 |
|   `SetBool(x bool)`   |      使用 bool 设置值。当值的类型不是 bod 时会发生宕机       |
| `SetBytes(x []byte)`  |  设置字节数组 []bytes值。当值的类型不是 []byte 时会发生宕机  |
| `SetString(x string)` |       设置字符串值。当值的类型不是 string 时会发生宕机       |

```go
// 声明整数变量a并赋值
var a int = 1024

// 获取变量a的反射值对象(a的地址)
valueOfA := reflect.ValueOf(&a)

// 取出a地址的元素(a的值)
valueOfA = valueOfA.Elem()

// 修改a的值为1
valueOfA.SetInt(1)

fmt.Println(valueOfA.Int())
```

#### 值可修改条件之一： 被导出

```go
type dog struct {
	legCount int
}

// 获取dog实例的反射对象
valueOfDog := reflect.ValueOf(dog{})

// 获取legCount字段的值
vLegCount := valueOfDog.FieldByName("letCount")

// 尝试设置legCount的值为4
vLegCount.SetInt(4)
```

会报错

```
panic: reflect: reflect.Value.SetInt using value obtained using unexported field
```

修改为可导出
```go
type dog struct {
	LegCount int
}

// 获取dog实例的反射对象
valueOfDog := reflect.ValueOf(dog{})

// 获取legCount字段的值
vLegCount := valueOfDog.FieldByName("LegCount")

// 尝试设置legCount的值为4
vLegCount.SetInt(4)
```

仍然报错

```
panic: reflect: reflect.Value.SetInt using unaddressable value
```

```go
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

fmt.Println(vLegCount.Int()) // 4
```

### 通过类型信息创建实例

```go
var a int

// 取变量a的反射类型对象
typeOfA := reflect.TypeOf(a)

// 根据反射类型对象创建类型实例
aIns := reflect.New(typeOfA)

// 输出Value的类型和种类
fmt.Println(aIns.Type(), aIns.Kind()) // *int ptr
```

## The Law of Reflection

Go 是静态类型的，每个变量都有一个静态类型，即在编译时已知且固定的一种类型: int、 float32、 *MyType、[]Type等。

```go
type MyInt int

var i int
var j MyInt
```

变量 i 的类型是 int，变量 j 的类型是 MyInt，虽然它们有着相同的基本类型，但静态类型却不一样，在没有类型转换的情况下，它们之间无法互相赋值。

接口是一个重要的类型，它意味着一个确定的方法集合，一个接口变量可以存储任何实现了接口的方法的具体值（除了接口本身），例如 io.Reader 和 io.Writer：

```go
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

如果一个类型声明实现了 Reader（或 Writer）方法，那么它便实现了 io.Reader（或 io.Writer），这意味着一个 io.Reader 的变量可以持有任何一个实现了 Read 方法的的类型的值。

```go
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```

必须要弄清楚的一点是，不管变量 r 中的具体值是什么，r 的类型永远是 io.Reader，由于Go语言是静态类型的，r 的静态类型就是 io.Reader。

在接口类型中有一个极为重要的例子——空接口：

```go
interface{}
```

它表示了一个空的方法集，一切值都可以满足它，因为它们都有零值或方法。

有人说Go语言的接口是动态类型，这是错误的，它们都是静态类型，虽然在运行时中，接口变量存储的值也许会变，但接口变量的类型是不会变的。我们必须精确地了解这些，因为反射与接口是密切相关的。

### 反射第一定律： 定律可以将“接口类型变量“转换为“反射类型对象”

> 这里反射类型指 `reflect.Type` 和 `reflect.Value` 

在最开始，我们先了解下 reflect 包的两种类型 Type 和 Value，这两种类型使访问接口内的数据成为可能，它们对应两个简单的方法，分别是 reflect.TypeOf 和 reflect.ValueOf，分别用来读取接口变量的 reflect.Type 和 reflect.Value 部分。

当然，从 reflect.Value 也很容易获取到 reflect.Type，目前我们先将它们分开。

```go
var x float64 = 3.4
fmt.Println("type:", reflect.TypeOf(x))
```

输出:

```
type: float64
```

大家可能会疑惑，为什么没看到接口？这段代码看起来只是把一个 float64 类型的变量 x 传递给 reflect.TypeOf 并没有传递接口。其实在 reflect.TypeOf 的函数签名里包含一个空接口：

```go
// TypeOf returns the reflection Type of the value in the interface{}.
func TypeOf(i interface{}) Type
```

我们调用 `reflect.TypeOf(x)` 时，x 被存储在一个空接口变量中被传递过去，然后 `reflect.TypeOf` 对空接口变量进行拆解，恢复其类型信息。

```go
var x float64 = 3.4
fmt.Println("value:", reflect.ValueOf(x))
```

输出:
```
value: 3.4
```

类型 reflect.Value 有一个方法 Type()，它会返回一个 reflect.Type 类型的对象。

Type 和 Value 都有一个名为 Kind 的方法，它会返回一个常量，表示底层数据的类型，常见值有：Uint、Float64、Slice 等

Value 类型也有一些类似于 Int、Float 的方法，用来提取底层的数据：
- `Int` 方法用来提取 `int64`
- `Float` 方法用来提取 `float64`，

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```

输出:
```
type: float64
kind is float64: true
value: 3.4
```

还有一些用来修改数据的方法，比如 SetInt、SetFloat。在介绍它们之前，我们要先理解“可修改性”（settability），这一特性会在下面进行详细说明。

首先是介绍下 `Value` 的 `getter` 和 `setter` 方法，为了保证 API 的精简，这两个方法操作的是某一组类型范围最大的那个。比如，处理任何含符号整型数，都使用 `int64`，也就是说 `Value` 类型的 Int 方法返回值为 `int64` 类型，`SetInt` 方法接收的参数类型也是 `int64` 类型。实际使用时，可能需要转化为实际的类型：


```go
var x uint8 = 'x'
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())                            // uint8.
fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true
x = uint8(v.Uint())                                       // v.Uint returns a uint64.
```

### 反射第二定律：反射可以将“反射类型对象“转换为“接口类型变量“

根据一个 reflect.Value 类型的变量，我们可以使用 Interface 方法恢复其接口类型的值。事实上，这个方法会把 type 和 value 信息打包并填充到一个接口变量中，然后返回。

其函数声明如下：

```go
// Interface returns v's value as an interface{}.
func (v Value) Interface() interface{}
```

然后，我们可以通过断言，恢复底层的具体值：

```go
y := v.Interface().(float64) // y will have type float64.
fmt.Println(y)
```

### 反射第三定律：如果要修改“反射类型对象“其值必须是“可写的”

下面这段代码虽然不能正常工作，但是非常值得研究：

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1) // Error: will panic
```

如果运行这段代码，它会抛出一个奇怪的异常：

```
panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
```

这里问题不在于值`7.1`不能被寻址，而是因为变量 v 是“不可写的”，“可写性”是反射类型变量的一个属性，但不是所有的反射类型变量都拥有这个属性。

我们可以通过 CanSet 方法检查一个 reflect.Value 类型变量的“可写性”，

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("settability of v:", v.CanSet())
```

```
settability of v: false
```

对于一个不具有“可写性”的 Value 类型变量，调用 Set 方法会报出错误。

首先我们要弄清楚什么是“可写性”，“可写性”有些类似于寻址能力，但是更严格，它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。“可写性”最终是由一个反射对象是否存储了原始值而决定的。


示例代码如下：

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
```

这里我们传递给 `reflect.ValueOf` 函数的是变量 `x` 的一个拷贝，而非 `x` 本身，想象一下如果下面这行代码能够成功执行：

```go
v.SetFloat(7.1)
```

如果这行代码能够成功执行，它不会更新 `x`，虽然看起来变量 `v` 是根据 `x` 创建的，相反它会更新 `x` 存在于反射对象 `v` 内部的一个拷贝，而变量 `x` 本身完全不受影响。这会造成迷惑，并且没有任何意义，所以是不合法的。“可写性”就是为了避免这个问题而设计的。

这看起来很诡异，事实上并非如此，而且类似的情况很常见。考虑下面这行代码：

```go
f(x)
```

代码中，我们把变量 `x` 的一个拷贝传递给函数，因此不期望它会改变 `x` 的值。如果期望函数 `f` 能够修改变量 `x`，我们必须传递 `x` 的地址（即指向 `x` 的指针）给函数 `f`，如下所示：

```go
f(&x)
```

反射的工作机制与此相同，如果想通过反射修改变量 x，就要把想要修改的变量的指针传递给反射库。

```go
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
fmt.Println("type of p: ", p.Type())
fmt.Println("settability of p: ", p.CanSet())
```

```
type of p:  *float64
settability of p:  false
```

反射对象 `p` 是不可写的，但是我们也不想修改 `p`，事实上我们要修改的是 `*p`。为了得到 `p` 指向的数据，可以调用 `Value` 类型的 Elem 方法。`Elem` 方法能够对指针进行“解引用”，然后将结果存储到反射 `Value` 类型对象 `v` 中：

```go
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
v := p.Elem()
fmt.Println("settability of v:", v.CanSet())
```

```
settability of v: true
```

由于变量 `v` 代表 `x`， 因此我们可以使用 `v.SetFloat` 修改 `x` 的值：

```go
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
v := p.Elem()
v.SetFloat(7.1)
fmt.Println(v.Interface())
fmt.Println(x)
```

输出:
```
7.1
7.1
```

反射不太容易理解，`reflect.Type` 和 `reflect.Value` 会混淆正在执行的程序，但是它做的事情正是编程语言做的事情。只需要记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。

### 结构体

我们一般使用反射修改结构体的字段，只要有结构体的指针，我们就可以修改它的字段。

下面是一个解析结构体变量 t 的例子，用结构体的地址创建反射变量，再修改它。然后我们对它的类型设置了 `typeOfT`，并用调用简单的方法迭代字段。

需要注意的是，我们从结构体的类型中提取了字段的名字，但每个字段本身是正常的 `reflect.Value` 对象。

```go
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
```

```
0: A int = 23
1: B string = skidoo
```

T 字段名之所以大写，是因为结构体中只有可导出的字段是“可设置”的。

```go
type T struct {
    A int
    B string
}
t := T{23, "skidoo"}
s := reflect.ValueOf(&t).Elem()
s.Field(0).SetInt(77)
s.Field(1).SetString("Sunset Strip")
fmt.Println("t is now", t)
```

```
t is now {77 Sunset Strip}
```

如果我们修改了程序让 s 由 t（而不是 &t）创建，程序就会在调用 SetInt 和 SetString 的地方失败，因为 t 的字段是不可设置的。

### 总结

反射规则可以总结为如下几条：
- 反射可以将“接口类型变量”转换为“反射类型对象”；
- 反射可以将“反射类型对象”转换为“接口类型变量”；
- 如果要修改“反射类型对象”，其值必须是“可写的”。

