package main

import (
	"fmt"
	"io/fs"
	"os"
)

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello world!")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteBytes(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	size, err := file.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Wrote %d bytes to file", size)
}

func WriteByLines(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for _, val := range lines {
		_, err := fmt.Fprintln(file, val)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func Append(filename, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	CreateFile("hello.txt")
	WriteBytes("bytes.txt", []byte("Hello World with byte slice"))

	lines := []string{
		"This is the first line",
		"Second line",
		"Now is the third line!",
	}

	WriteByLines("lines.txt", lines)
	Append("lines.txt", "New line appended")
}
