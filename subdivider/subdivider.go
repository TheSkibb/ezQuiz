package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
 */
func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify a file to read")
	}

	arg1 := os.Args[1]
	file, err := os.Open(arg1)
	defer file.Close()

	if err != nil {
		log.Fatal("could not open file", err)
	}

	fmt.Println("subdividing file")

	scanner := bufio.NewScanner(file)

	fileToWriteTo := ""

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 && line[:2] == "# " {
			fileToWriteTo = line[2:]
			fmt.Println(fileToWriteTo)
		} else if fileToWriteTo == "" {
			fmt.Println("skipping line")
			continue
		} else {
			fmt.Println("writing line")
			lineBytes := []byte(line)
			err := os.WriteFile(fileToWriteTo, lineBytes, 0644)
			if err != nil {
				log.Fatal("cannot write file", err)
			}
		}
	}
}
