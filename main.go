package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type question struct {
	question string
	answer   string
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify a file to read")
	}
	arg1 := os.Args[1]

	file, err := os.Open(arg1)

	if err != nil {
		log.Fatal("could not open file:", arg1, err)
	}

	scanner := bufio.NewScanner(file)

	// index questions

	var questions []question

	questionTemp := ""
	answerTemp := ""

	for scanner.Scan() {
		currline := scanner.Text()
		if len(currline) > 2 && currline[:2] == "##" {
			if questionTemp != "" && answerTemp != "" {
				questions = append(questions, question{questionTemp, answerTemp})
			}
			questionTemp = currline[2:]
			answerTemp = ""
		} else if (len(currline) > 2 && currline[:1] != "#") || currline == "" {
			answerTemp += currline + "\n"
		}
	}

	startShell(questions)
}

func startShell(questions []question) {

	reader := bufio.NewReader(os.Stdin)
	message := ""

	fmt.Println("welcome to the flashcard program, press enter to select a random question! :-)")
	for message != "exit" {
		fmt.Print(">")
		var err error
		message, err = reader.ReadString('\n')
		message = message[:len(message)-1]
		if err != nil {
			log.Fatal("error reading input", err)
		}
		//fmt.Printf(message + "\n")
		message = handleInput(message, questions)
	}
}

func handleInput(input string, questions []question) string {
	switch input {
	case "":
		printRandomQuestion(questions)
	case "help", "h":
		fmt.Println("help is on the way")
	case "exit", "quit", "q":
		return "exit"
	default:
		// check if input is a number
		n, err := strconv.Atoi(input)
		if err == nil {
			printIndexedQuestion(questions, n)
		} else {
			fmt.Println("unregognized command")
		}
	}
	return ""
}

func printIndexedQuestion(questions []question, index int) {
	max := len(questions) - 1
	index = index % max

	fmt.Println(index, questions[index].question)

	//press anything to continue
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	answer := questions[index].answer

	if answer == "" {
		fmt.Println("no answer specified")
	} else {
		fmt.Println(answer)
	}

}

func printRandomQuestion(questions []question) {
	max := len(questions) - 1
	randIndex := rand.Int() % max

	fmt.Println(randIndex, questions[randIndex].question)

	//press anything to continue
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	answer := questions[randIndex].answer

	if answer == "" {
		fmt.Println("no answer specified")
	} else {
		fmt.Println(answer)
	}
}
