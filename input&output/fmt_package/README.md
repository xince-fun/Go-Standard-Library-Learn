# fmt - 格式化IO

fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf. 格式“占位符”衍生自C，但比C更简单。

fmt 包的官方文档对 Printing 和 Scanning 有很详细的说明。这里就直接引用文档进行说明，同时附上额外的说明或例子，之后再介绍具体的函数使用。

## Printing

### Sample

```go

type user struct {
	name string
}

func main() {
	u := user{"tang"}
	//Printf 格式化输出
	fmt.Printf("% + v\n", u)     //格式化输出结构
	fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
	fmt.Printf("%p\n", 0x123)    //0x开头的十六进制数表示
}
```

### 占位符

**普通占位符**
	
	占位符						说明						举例										输出
	%v		相应值的默认格式。								Printf("%v", site)，Printf("%+v", site)	{studygolang}，{Name:studygolang}
			在打印结构体时，“加号”标记（%+v）会添加字段名
	%#v		相应值的Go语法表示							Printf("#v", site)						main.Website{Name:"studygolang"}
	%T		相应值的类型的Go语法表示						Printf("%T", site)						main.Website
	%%		字面上的百分号，并非值的占位符					Printf("%%")							%

**布尔占位符**

	占位符						说明						举例										输出
	%t		单词 true 或 false。							Printf("%t", true)						true

**整数占位符**

	占位符						说明						举例									输出
	%b		二进制表示									Printf("%b", 5)						101
	%c		相应Unicode码点所表示的字符					Printf("%c", 0x4E2D)				中
	%d		十进制表示									Printf("%d", 0x12)					18
	%o		八进制表示									Printf("%o", 10)					12
	%q		单引号围绕的字符字面值，由Go语法安全地转义		Printf("%q", 0x4E2D)				'中'
	%x		十六进制表示，字母形式为小写 a-f				Printf("%x", 13)					d
	%X		十六进制表示，字母形式为大写 A-F				Printf("%x", 13)					D
	%U		Unicode格式：U+1234，等同于 "U+%04X"			Printf("%U", 0x4E2D)				U+4E2D

**浮点数和复数的组成部分（实部和虚部）**

	占位符						说明												举例									输出
	%b		无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat	
			的 'b' 转换格式一致。例如 -123456p-78
	%e		科学计数法，例如 -1234.456e+78									Printf("%e", 10.2)							1.020000e+01
	%E		科学计数法，例如 -1234.456E+78									Printf("%e", 10.2)							1.020000E+01
	%f		有小数点而无指数，例如 123.456									Printf("%f", 10.2)							10.200000
	%g		根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%g", 10.20)							10.2
	%G		根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%G", 10.20+2i)						(10.2+2i)

**字符串与字节切片**

	占位符						说明												举例									输出
	%s		输出字符串表示（string类型或[]byte)							Printf("%s", []byte("Go语言中文网"))		Go语言中文网
	%q		双引号围绕的字符串，由Go语法安全地转义							Printf("%q", "Go语言中文网")				"Go语言中文网"
	%x		十六进制，小写字母，每字节两个字符								Printf("%x", "golang")						676f6c616e67
	%X		十六进制，大写字母，每字节两个字符								Printf("%X", "golang")						676F6C616E67

**指针**

	占位符						说明												举例									输出
	%p		十六进制表示，前缀 0x											Printf("%p", &site)							0x4f57f0
	
这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地，这里也不需要指定操作数的大小（int8，int64）。

宽度与精度的控制格式以 Unicode 码点为单位。（这点与C的 printf 不同，它以字节数为单位）二者或其中之一均可用字符 '*' 表示，此时它们的值会从下一个操作数中获取，该操作数的类型必须为 int。

对数值而言，宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。 但对于 %g/%G 而言，精度为所有数字的总数。例如，对于123.45，格式 %6.2f 会打印123.45，而 %.4g 会打印123.5。%e 和 %f 的默认精度为6；但对于 %g 而言，它的默认精度为确定该值所必须的最小位数。

对大多数的值而言，宽度为输出的最小字符数，如果必要的话会为已格式化的形式填充空格。对字符串而言，精度为输出的最大字符数，如果必要的话会直接截断。

### Scanning

一组类似的函数通过扫描已格式化的文本来产生值。 Scan、Scanf 和 Scanln 从 os.Stdin 中读取； Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 中读取； Sscan、Sscanf 和 Sscanln 从实参字符串中读取。 Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，且需要条目紧随换行符之后； Scanf、Fscanf 和 Sscanf 需要输入换行符来匹配格式中的换行符；其它函数则将换行符视为空格。

Scanf、Fscanf 和 Sscanf 根据格式字符串解析实参，类似于 Printf。例如，%x 会将一个整数扫描为十六进制数，而 %v 则会扫描该值的默认表现格式。

格式化行为类似于 Printf，但也有如下例外：

```
%p 没有实现
%T 没有实现
%e %E %f %F %g %G 都完全等价，且可扫描任何浮点数或复数数值
%s 和 %v 在扫描字符串时会将其中的空格作为分隔符
标记 # 和 + 没有实现
```

在使用 %v 占位符扫描整数时，可接受友好的进制前缀0（八进制）和0x（十六进制）。

宽度被解释为输入的文本（%5s 意为最多从输入中读取5个 rune 来扫描成字符串），而扫描函数则没有精度的语法（没有 %5.2f，只有 %5f）。

当以某种格式进行扫描时，无论在格式中还是在输入中，所有非空的连续空白字符 （除换行符外）都等价于单个空格。由于这种限制，格式字符串文本必须匹配输入的文本，如果不匹配，扫描过程就会停止，并返回已扫描的实参数。

在所有的扫描参数中，若一个操作数实现了 Scan 方法（即它实现了 Scanner 接口）， 该操作数将使用该方法扫描其文本。此外，若已扫描的实参数少于所提供的实参数，就会返回一个错误。

所有需要被扫描的实参都必须是基本类型或 Scanner 接口的实现。

注意：Fscan 等函数会从输入中多读取一个字符（rune），因此，如果循环调用扫描函数，可能会跳过输入中的某些数据。一般只有在输入的数据中没有空白符时该问题才会出现。若提供给 Fscan 的读取器实现了 ReadRune，就会用该方法读取字符。若此读取器还实现了 UnreadRune 方法，就会用该方法保存字符，而连续的调用将不会丢失数据。若要为没有 ReadRune 和 UnreadRune 方法的读取器加上这些功能，需使用 bufio.NewReader。

### Print 序列函数

这里说的 Print 序列函数包括：Fprint/Fprintf/Fprintln/Sprint/Sprintf/Sprintln/Print/Printf/Println。之所以将放在一起介绍，是因为它们的使用方式类似、参数意思也类似。

一般的，我们将 Fprint/Fprintf/Fprintln 归为一类；Sprint/Sprintf/Sprintln 归为一类；Print/Printf/Println 归为另一类。其中，Print/Printf/Println 会调用相应的F开头一类函数。如：

```go
func Print(a ...interface{}) (n int, err error) {
    return Fprint(os.Stdout, a...)
}
```

Fprint/Fprintf/Fprintln 函数的第一个参数接收一个io.Writer类型，会将内容输出到 io.Writer 中去。而 Print/Printf/Println 函数是将内容输出到标准输出中，因此，直接调用 F类函数 做这件事，并将 os.Stdout 作为第一个参数传入。

Sprint/Sprintf/Sprintln 是格式化内容为 string 类型，而并不输出到某处，需要格式化字符串并返回时，可以用这组函数。

在这三组函数中，S/F/Printf函数通过指定的格式输出或格式化内容；S/F/Print函数只是使用默认的格式输出或格式化内容；S/F/Println函数使用默认的格式输出或格式化内容，同时会在最后加上"换行符"。

### Stringer 接口

Stringer 接口定义如下:

```go
type Stringer interface {
    String() string
}
```

根据 Go 语言中实现接口的定义，一个类型只要有 String() string 方法，我们就说它实现了 Stringer 接口。而在本节开始已经说到，如果格式化输出某种类型的值，只要它实现了 String() 方法，那么会调用 String() 方法进行处理。

我们定义如下struct：

```go
type Person struct {
	Name string
	Age  int
	Sex  int
}
```

```go
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
```

```bash
This is polaris, He is 28 years old.
```

可见，Stringer接口和Java中的ToString方法类似。

### Formatter 接口

Formatter 接口的定义如下:

```go
type Formatter interface {
	Format(f State, c rune)
}
```

> Formatter 接口由带有定制的格式化器的值所实现。 Format 的实现可调用 Sprintf 或 Fprintf(f) 等函数来生成其输出。

```go
func (p *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte(" Person has three fileds."))
	} else {
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}
```

这里需要解释以下几点：

1）fmt.State 是一个接口。由于 Format 方法是被 fmt 包调用的，它内部会实例化好一个 fmt.State 接口的实例，我们不需要关心该接口；

2）可以实现自定义占位符，同时 fmt 包中和类型相对应的预定义占位符会无效。因此例子中 Format 的实现加上了 else 子句；

3）实现了 Formatter 接口，相应的 Stringer 接口不起作用。但实现了 Formatter 接口的类型应该实现 Stringer 接口，这样方便在 Format 方法中调用 String() 方法。就像本例的做法；

4）Format 方法的第二个参数是占位符中%后的字母（有精度和宽度会被忽略，只保留字母）；

一般地，我们不需要实现 Formatter 接口。如果对 Formatter 接口的实现感兴趣，可以看看标准库 math/big 包中 Int 类型的 Formatter 接口实现。

### GoStringer 接口

GoStringer 接口定义如下:

```go
type GoStringer interface {
	GoString() string
}
```

该接口定义了类型的Go语法格式。用于打印(Printf)格式化占位符为 %#v 的值。

用前面的例子演示。执行：

```go
&main.Person{Name:"polaris", Age:28, Sex:0}%
```

```
&Person{Name is polaris, Age is 28, Sex is 0}
```

### Scan 序列函数

该序列函数和 Print 序列函数相对应，包括：Fscan/Fscanf/Fscanln/Sscan/Sscanf/Sscanln/Scan/Scanf/Scanln。

一般的，我们将Fscan/Fscanf/Fscanln归为一类；Sscan/Sscanf/Sscanln归为一类；Scan/Scanf/Scanln归为另一类。其中，Scan/Scanf/Scanln会调用相应的F开头一类函数。如：

```go
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```

Fscan/Fscanf/Fscanln 函数的第一个参数接收一个 io.Reader 类型，从其读取内容并赋值给相应的实参。而 Scan/Scanf/Scanln 正是从标准输入获取内容，因此，直接调用 F类函数 做这件事，并将 os.Stdin 作为第一个参数传入。

Sscan/Sscanf/Sscanln 则直接从字符串中获取内容。

对于Scan/Scanf/Scanln三个函数的区别，我们通过例子来说明，为了方便讲解，我们使用Sscan/Sscanf/Sscanln这组函数。

1. Scan/FScan/Sscan

```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscan("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscan("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
```

```
2 polaris 28
```

2. Scanf/FScanf/Sscanf

```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscanf("polaris 28", "%s%d", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscanf("polaris\n28", "%s%d", &name, &age)
	fmt.Println(n, name, age)
```

如果将"空格"分隔改为"\n"分隔，则输出为：1 polaris 0。可见，Scanf/FScanf/Sscanf 这组函数将连续由空格分隔的值存储为连续的实参， 其格式由 format 决定，换行符处停止扫描(Scan)。

3. Scanln/FScanln/Sscanln

```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscanln("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscanln("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
```

Scanln/FScanln/Sscanln表现和上一组一样，遇到"\n"停止（对于Scanln，表示从标准输入获取内容，最后需要回车）。

一般地，我们使用 Scan/Scanf/Scanln 这组函数。