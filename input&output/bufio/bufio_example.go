package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// line, _ := reader.ReadSlice('\n')
	// fmt.Printf("the line:%s\n", line)
	// //这里可以换上任意的 bufio 的 Read/Write 操作
	// n, _ := reader.ReadSlice('\n')
	// fmt.Printf("the line:%s\n", line)
	// fmt.Println(string(n))

	// reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	// line, err := reader.ReadSlice('\n')
	// fmt.Printf("line:%s\terror:%s\n", line, err)
	// line, err = reader.ReadSlice('\n')
	// fmt.Printf("line:%s\terror:%s\n", line, err)

	// reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// line, _ := reader.ReadBytes('\n')
	// fmt.Printf("the line:%s\n", line)
	// // 这里可以换上任意的 bufio 的 Read/Write 操作
	// n, _ := reader.ReadBytes('\n')
	// fmt.Printf("the line:%s\n", line)
	// fmt.Println(string(n))

	// reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	// go Peek(reader)
	// go reader.ReadBytes('\t')
	// time.Sleep(1e8)

	// const input = "This is The Golang Standard Library.\nWelcome you!"
	// scanner := bufio.NewScanner(strings.NewReader(input))
	// scanner.Split(bufio.ScanWords)
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
	// fmt.Println(count)

	// scanner := bufio.NewScanner(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// if scanner.Scan() {
	// 	scanner.Scan()
	// 	fmt.Printf("%s", scanner.Text())
	// }

	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com\nIt is the home of gophers.\nIf you are studying golang, Welcome you!")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	// time.Sleep(1)
	fmt.Printf("%s\n", line)
}
