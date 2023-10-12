# IO

---

`io` 包为 I/O 原语提供了基本的接口。它主要包装了这些原语的已有实现。

由于这些被接口包装的I/O原语是由不同的低级操作实现，因此，在另有声明之前不该假定它们的并发执行是安全的。

`io`包最终重要的两个接口是： `Reader` 和 `Writer` 接口。 只要满足这两个接口，就可以使用IO包的功能。

## Reader接口

Reader接口的定义如下：

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。若可读取的数据不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。

> 当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file)，它会返回读取的字节数。它可能会同时在本次的调用中返回一个non-nil错误,或在下一次的调用中返回这个错误（且 n 为 0）。 一般情况下, Reader会返回一个非0字节数n, 若 n = len(p) 个字节从输入源的结尾处由 Read 返回，Read可能返回 err == EOF 或者 err == nil。并且之后的 Read() 都应该返回 (n:0, err:EOF)。

> 调用者在考虑错误之前应当首先处理返回的数据。这样做可以正确地处理在读取一些字节后产生的 I/O 错误，同时允许EOF的出现。

## Writer接口

> Write 将 len(p) 个字节从 p 中写入到基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。若 Write 返回的 n < len(p)，它就必须返回一个 非nil 的错误。

同样的，所有实现了Write方法的类型都实现了 io.Writer 接口。

在fmt标准库中，有一组函数：Fprint/Fprintf/Fprintln，它们接收一个 io.Wrtier 类型参数（第一个参数），也就是说它们将数据格式化输出到 io.Writer 中。那么，调用这组函数时，该如何传递这个参数呢？

我们以 fmt.Fprintln 为例，同时看一下 fmt.Println 函数的源码。

```go
func Println(a ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}
```

## 实现了io.Reader接口或者io.Wrtier接口的类型

我们可以知道，os.File 同时实现了这两个接口。我们还看到 os.Stdin/Stdout 这样的代码，它们似乎分别实现了 io.Reader/io.Writer 接口。没错，实际上在 os 包中有这样的代码：

```go
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

- os.File 同时实现了 io.Reader 和 io.Writer
- strings.Reader 实现了 io.Reader
- bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
- bytes.Buffer 同时实现了 io.Reader 和 io.Writer
- bytes.Reader 实现了 io.Reader
- compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
- crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
- crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
- encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
- mime/multipart.Part 实现了 io.Reader
- net/conn 分别实现了 io.Reader 和 io.Writer(Conn接口定义了Read/Write)

## ReaderAt和WriterAt接口

**ReaderAt** 接口的定义如下:

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

> ReadAt 从基本输入源的偏移量 off 处开始，将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)）以及任何遇到的错误。

> 当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节。在这一点上，ReadAt 比 Read 更严格。

> 即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。若可读取的数据不到 len(p) 字节，ReadAt 就会阻塞,直到所有数据都可用或一个错误发生。 在这一点上 ReadAt 不同于 Read。

> 若 n = len(p) 个字节从输入源的结尾处由 ReadAt 返回，Read可能返回 err == EOF 或者 err == nil

> 若 ReadAt 携带一个偏移量从输入源读取，ReadAt 应当既不影响偏移量也不被它所影响。

> 可对相同的输入源并行执行 ReadAt 调用。

可见，ReaderAt 接口使得可以从指定偏移量处开始读取数据。

**WriterAt**接口的定义如下

```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

> WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off 处的基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。若 WriteAt 返回的 n < len(p)，它就必须返回一个 非nil 的错误。

> 若 WriteAt 携带一个偏移量写入到目标中，WriteAt 应当既不影响偏移量也不被它所影响。

> 若被写区域没有重叠，可对相同的目标并行执行 WriteAt 调用。

## ReaderFrom和WriteTo接口

ReaderFrom的定义如下:

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

> ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。

> 如果 ReadFrom 可用， Copy 函数就会使用它。

注意: ReadFrom 方法不会返回 err = EOF。

WriteTo的定义如下:

```go
type WriteTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

> WriteTo 将数据写入 w 中，直到没有数据可写或发生错误。其返回值 n 为写入的字节数。在写入过程中遇到的任何错误也将被返回。

> 如果 WriteTo 可用， Copy 函数就会使用它

其实 ReaderFrom 和 WriterTo 接口的方法接收的参数是 io.Reader 和 io.Writer 类型。根据 io.Reader 和 io.Writer 接口的讲解，对该接口的使用应该可以很好的掌握。

## Seeker 接口

接口定义如下:

```go
type Seeker interface {
    Seek(offset int64, whence int) (ret int64, err error)
}
```

> Seek 设置下一次 Read 或 Write 的偏移量为 offset，它的解释取决于 whence: 0 表示相对于文件的起始处，1 表示相对于当前的偏移， 而2 表示相对于其结尾处。Seek 返回新的偏移量和一个错误，如果有的话。

也就是说，Seek 方法是用于设置偏移量的，这样可以从某个特定位置开始操作数据流。听起来和 Reader/WriteAt 接口有些类似，不过 Seeker 接口更灵活，可以更好的控制读写数据流的位置。

whence 的值，在 io 包中定义了相应的常量
```go
const (
  SeekStart   = 0 // seek relative to the origin of the file
  SeekCurrent = 1 // seek relative to the current offset
  SeekEnd     = 2 // seek relative to the end
)
```

## Closer接口

接口定义如下:

```go
type Closer interface {
    Close() error
}
```

该接口比较简单，只有一个 Close() 方法，用于关闭数据流。

文件（os.File）、归档（压缩包）、数据库连接、Socket 等需要手动关闭连接的资源都实现了 Closer 接口。

实际编程中，经常将 Close 方法的调用放在 defer 里面。

```go
file, err := os.Open("studygolang.txt")
defer file.Close()
if err != nil {
    ...
}
```

当文件 studygolang.txt 不存在或找不到时， file.Close() 会 panic，因为 file 是 nil。因此，应该将 defer file.Close() 放在错误检查之后。

其实好像不会，但还是先处理错误比较好。

```go
func (f *File) Close() error {
	if f == nil {
		return ErrInvalid
	}
	return f.file.close()
}
```

## 其他接口

### ByteReader 和 ByteWriter

通过名称大概也能猜出这组接口的用途：读或写一个字节。接口定义如下：

```go
type ByteReader interface {
    ReadByte() (c byte, err error)
}

type ByteWriter interface {
    WriteByte(c byte) error
}
```

在标准库中，有如下类型实现了 io.ByteReader 和 io.ByteWriter:

- bufio.Reader/Writer 分别实现了io.ByteReader 和 io.ByteWriter
- bytes.Buffer 同时实现了 io.ByteReader 和 io.ByteWriter
- bytes.Reader 实现了 io.ByteReader
- strings.Reader 实现了 io.ByteReader

### ByteScanner、RuneReader和RuneScanner

将这三个接口放在一起，是考虑到与 ByteReader 相关或相应。

ByteScanner的接口定义如下:

```go
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}
```

可见，它内嵌了 ByteReader 接口（可以理解为继承了 ByteReader 接口），UnreadByte 方法的意思是：将上一次 ReadByte 的字节还原，使得再次调用 ReadByte 返回的结果和上一次调用相同，也就是说，UnreadByte 是重置上一次的 ReadByte。注意，UnreadByte 调用之前必须调用了 ReadByte，且不能连续调用 UnreadByte。即：

```go
buffer := bytes.NewBuffer([]byte{'a', 'b'})
err := buffer.UnreadByte()
```


```go
buffer := bytes.NewBuffer([]byte{'a', 'b'})
buffer.ReadByte()
err := buffer.UnreadByte()
err = buffer.UnreadByte()
```

### ReadCloser、ReadSeeker、ReadWriteCloser、ReadWriteSeeker、ReadWriter、WriteCloser 和 WriteSeeker 接口

这些接口是上面介绍的接口的两个或三个组合而成的新接口。例如 ReadWriter 接口：

```go
type ReadWriter interface {
    Reader
    Writer
}
```

这是 Reader 接口和 Writer 接口的简单组合（内嵌）。

这些接口的作用是：有些时候同时需要某两个接口的所有功能，即必须同时实现了某两个接口的类型才能够被传入使用。可见，io 包中有大量的“小接口”，这样方便组合为“大接口”。

### SectionReader 类型

SectionReader 是一个 struct（没有任何导出的字段），实现了 Read, Seek 和 ReadAt，同时，内嵌了 ReaderAt 接口。结构定义如下：

```go
type SectionReader struct {
	r     ReaderAt	// 该类型最终的 Read/ReadAt 最终都是通过 r 的 ReadAt 实现
	base  int64		// NewSectionReader 会将 base 设置为 off
	off   int64		// 从 r 中的 off 偏移处开始读取数据
	limit int64		// limit - off = SectionReader 流的长度
}
```

```go
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
```

> NewSectionReader 返回一个 SectionReader，它从 r 中的偏移量 off 处读取 n 个字节后以 EOF 停止。

也就是说，SectionReader 只是内部（内嵌）ReaderAt 表示的数据流的一部分：从 off 开始后的 n 个字节。

这个类型的作用是：方便重复操作某一段 (section) 数据流；或者同时需要 ReadAt 和 Seek 的功能。

### LimitedReader 类型

LimitedReader 结构定义如下

```go
type LimitedReader struct {
    R Reader // underlying reader，最终的读取操作通过 R.Read 完成
    N int64  // max bytes remaining
}
```

> 从 R 读取但将返回的数据量限制为 N 字节。每调用一次 Read 都将更新 N 来反应新的剩余数量。

也就是说，最多只能返回 N 字节数据。

LimitedReader 只实现了 Read 方法（Reader 接口）。

### PipeReader 和 PipeWriter 类型

PipeReader（一个没有任何导出字段的 struct）是管道的读取端。它实现了 io.Reader 和 io.Closer 接口。结构定义如下：

```go
type PipeReader struct {
	p *pipe
}
```

**关于 PipeReader.Read 方法的说明**： 从管道中读取数据。该方法会堵塞，直到管道写入端开始写入数据或写入端被关闭。如果写入端关闭时带有 error（即调用 CloseWithError 关闭），该Read返回的 err 就是写入端传递的error；否则 err 为 EOF。

PipeWriter（一个没有任何导出字段的 struct）是管道的写入端。它实现了 io.Writer 和 io.Closer 接口。结构定义如下：

```go
type PipeWriter struct {
    p *pipe
}
```

**关于 PipeReader.Read 方法的说明**： 写数据到管道中。该方法会堵塞，直到管道读取端读完所有数据或读取端被关闭。如果读取端关闭时带有 error（即调用 CloseWithError 关闭），该Write返回的 err 就是读取端传递的error；否则 err 为 ErrClosedPipe。

```go
func Pipe() (*PipeReader, *PipeWriter)
```

它将 io.Reader 连接到 io.Writer。一端的读取匹配另一端的写入，直接在这两端之间复制数据；它没有内部缓存。它对于并行调用 Read 和 Write 以及其它函数或 Close 来说都是安全的。一旦等待的 I/O 结束，Close 就会完成。并行调用 Read 或并行调用 Write 也同样安全：同种类的调用将按顺序进行控制。

正因为是同步的，因此不能在一个 goroutine 中进行读和写。

### Copy 和 CopyN 函数

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

> Copy 将 src 复制到 dst，直到在 src 上到达 EOF 或发生错误。它返回复制的字节数，如果有错误的话，还会返回在复制时遇到的第一个错误。

> 成功的 Copy 返回 err == nil，而非 err == EOF。由于 Copy 被定义为从 src 读取直到 EOF 为止，因此它不会将来自 Read 的 EOF 当做错误来报告。

> 若 dst 实现了 ReaderFrom 接口，其复制操作可通过调用 dst.ReadFrom(src) 实现。此外，若 src 实现了 WriterTo 接口，其复制操作可通过调用 src.WriteTo(dst) 实现。

```go
io.Copy(os.Stdout, strings.NewReader("Go语言中文网"))
```
直接将内容输出（写入 Stdout 中）。

```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

> CopyN 将 n 个字节(或到一个error)从 src 复制到 dst。 它返回复制的字节数以及在复制时遇到的最早的错误。当且仅当err == nil时,written == n 。

> 若 dst 实现了 ReaderFrom 接口，复制操作也就会使用它来实现。

```go
io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
```

输出

```bash
Go语言
```

### ReadAtLesat 和 ReadFull 函数

**ReadAtLeast**函数签名如下

```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

> ReadAtLeast 将 r 读取到 buf 中，直到读了最少 min 个字节为止。它返回复制的字节数，如果读取的字节较少，还会返回一个错误。若没有读取到字节，错误就只是 EOF。如果一个 EOF 发生在读取了少于 min 个字节之后，ReadAtLeast 就会返回 ErrUnexpectedEOF。若 min 大于 buf 的长度，ReadAtLeast 就会返回 ErrShortBuffer。对于返回值，当且仅当 err == nil 时，才有 n >= min。

**ReadFull**函数签名如下

```go
func ReadFull(r Reader, buf []byte) (n int, err error)
```

> ReadFull 精确地从 r 中将 len(buf) 个字节读取到 buf 中。它返回复制的字节数，如果读取的字节较少，还会返回一个错误。若没有读取到字节，错误就只是 EOF。如果一个 EOF 发生在读取了一些但不是所有的字节后，ReadFull 就会返回 ErrUnexpectedEOF。对于返回值，当且仅当 err == nil 时，才有 n == len(buf)。

注意该函数和 ReadAtLeast 的区别：ReadFull 将 buf 读满；而 ReadAtLeast 是最少读取 min 个字节。

### WriteString函数

```go
func WriteString(w Writer, s string) (n int, err error)
```

> WriteString 将s的内容写入w中，当 w 实现了 WriteString 方法时，会直接调用该方法，否则执行 w.Write([]byte(s))。


### MultiReader 和 MultiWriter 函数

```go
func MultiReader(readers ...Reader) Reader
func MultiWriter(writers ...Writer) Writer
```

它们接收多个 Reader 或 Writer，返回一个 Reader 或 Writer。我们可以猜想到这两个函数就是操作多个 Reader 或 Writer 就像操作一个。

事实上，在 io 包中定义了两个非导出类型：mutilReader 和 multiWriter，它们分别实现了 io.Reader 和 io.Writer 接口。类型定义为：

```go
type multiReader struct {
	readers []Reader
}

type multiWriter struct {
	writers []Writer
}
```

对于这两种类型对应的实现方法（Read 和 Write 方法）的使用，我们通过例子来演示。

**MultiReader的使用**：

```go
readers := []io.Reader{
	strings.NewReader("from strings reader"),
	bytes.NewBufferString("from bytes buffer"),
}
reader := io.MultiReader(readers...)
data := make([]byte, 0, 128)
buf := make([]byte, 10)
	
for n, err := reader.Read(buf); err != io.EOF ; n, err = reader.Read(buf){
	if err != nil{
		panic(err)
	}
	data = append(data,buf[:n]...)
}
fmt.Printf("%s\n", data)
```

输出

```bash
from strings readerfrom bytes buffer
```

### TeeReader 函数

```go
func TeeReader(r Reader, w Writer) Reader
```

TeeReader 返回一个 Reader，它将从 r 中读到的数据写入 w 中。所有经由它处理的从 r 的读取都匹配于对应的对 w 的写入。它没有内部缓存，即写入必须在读取完成前完成。任何在写入时遇到的错误都将作为读取错误返回。

```go
reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
reader.Read(make([]byte, 20))
```