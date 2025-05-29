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
	strToWrite := ""

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 && line[:2] == "# " {
			if strToWrite != "" {
				writeToFile(fileToWriteTo, strToWrite)
				strToWrite = ""
			}

			fileToWriteTo = line[2:]
			fmt.Println(fileToWriteTo)
		} else if fileToWriteTo == "" {
			fmt.Println("skipping line")
			continue
		}
		strToWrite += line + "\n"
	}

	if fileToWriteTo != "" && strToWrite != "" {
		writeToFile(fileToWriteTo, strToWrite)
	}

}

func writeToFile(fileToWriteTo, strToWrite string) {
	err := os.WriteFile(fileToWriteTo, []byte(strToWrite), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
