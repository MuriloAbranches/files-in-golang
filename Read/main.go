package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("File name: %s\n", stats.Name())
	fmt.Printf("Size of file: %v\n", stats.Size())
	fmt.Printf("Modified at: %v\n", stats.ModTime().Local())
}

func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(contents))
}

func ReadByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	buf := make([]byte, size)

	for {
		totalRead, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}

		fmt.Println(string(buf[:totalRead]))
	}
}

func ReadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=")
		key := raw[0]
		value := raw[1]
		fmt.Println(key)
		fmt.Println(value)
	}
}

func main() {
	filename := "test.txt"

	ReadStats(filename)
	ReadWholeFile(filename)
	ReadByLine(filename)
	ReadByWord(filename)
	ReadByBytes(filename, 8)
	ReadConfig("test.config")
}
