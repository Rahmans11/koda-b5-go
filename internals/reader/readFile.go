package reader

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReaderFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover from panic %v\n", r)
		}
	}()

	file, err := os.Open("D:/KodaCourse/CodeSession/minitaskGo/minitask1/note")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic("failed to read file")
	}

	fmt.Printf("%s\n", b)

}
