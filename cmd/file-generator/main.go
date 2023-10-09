package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	totalFiles := 100
	if len(os.Args) > 1 {
		total, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		totalFiles = total
	}
	GenerateFiles(totalFiles)
}

func GenerateFiles(totalFiles int) {
	for i := 0; i < totalFiles; i++ {
		fileName := fmt.Sprintf("tmp/%d.txt", i)
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		_, err = file.WriteString("Hello, World!")
		if err != nil {
			log.Println(err)
		}
		func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
	}
}
