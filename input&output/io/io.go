package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func PipeWrite(writer *io.PipeWriter) {
	data := []byte("Go语言中文网")
	for i := 0; i < 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节 %d\n", n)
	}
	writer.CloseWithError(errors.New("写入端已关闭"))
}

func PipeRead(reader *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接口端开始阻塞5秒钟...")
		time.Sleep(5 * time.Second)
		fmt.Println("接收端开始接受")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字节: %d\n buf内容: %s\n", n, buf)
	}
}

func main() {
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))

	fmt.Println("##### ReaderAt #######")
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s, %d\n", p, n)

	fmt.Println("##### WriterAt #######")
	file, err := os.Create("writeAt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	n, err = file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	fmt.Println("##### ReadFrom #######")
	file, err = os.Open("writeAt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()

	fmt.Println("##### Seeker #######")
	reader = strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)

	fmt.Println("##### ByteReader/ByteWriter ######")
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	err = buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	} else {
		fmt.Println("写入错误 ", err)
	}

	fmt.Println("###### UnreadByte ######")
	buffer = bytes.NewBuffer([]byte{'a', 'b'})
	err = buffer.UnreadByte()
	if err != nil {
		fmt.Println(err)
	}

	buffer = bytes.NewBuffer([]byte{'a', 'b'})
	buffer.ReadByte()
	err = buffer.UnreadByte()
	if err != nil {
		fmt.Println(err)
	}
	err = buffer.UnreadByte()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("###### LimitedReader #######")
	content := "This is LimitReader Example"
	reader = strings.NewReader(content)
	limitedReader := &io.LimitedReader{R: reader, N: 8}
	for limitedReader.N > 0 {
		tmp := make([]byte, 2)
		limitedReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}

	fmt.Println("####### PipeReader/PipeWriter #######")
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(30 * time.Second)

	fmt.Println("####### MultiReader/MultiWriter #######")
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	newreader := io.MultiReader(readers...)

	data = make([]byte, 0, 128)
	buf := make([]byte, 10)

	for n, err := newreader.Read(buf); err != io.EOF; n, err = newreader.Read(buf) {
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}
	fmt.Printf("%s\n", data)

	file, err = os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	newwriter := io.MultiWriter(writers...)
	newwriter.Write([]byte("Go语言中文网"))
}
