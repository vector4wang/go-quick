package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//readFile("C:\\Users\\wdc43\\Desktop\\log.txt")
	//readFileUseUtil("C:\\Users\\wdc43\\Desktop\\log.txt")
	writeFile("C:\\Users\\wdc43\\Desktop\\log2.txt")
}

func writeFile(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	n, err1 := io.WriteString(f, "test")
	if err1 != nil {
		fmt.Println("write error", err1)
		return
	}
	fmt.Println("write byte is", n)
}

func readFileUseUtil(path string) {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	fmt.Println(string(f))
}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read fail")
	}

	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)

		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}

	fmt.Println("file is", string(chunk))
}
