package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type questionList struct {
	questions []question
	titles    []string
}

type question struct {
	question   string
	answer     string
	answered   bool
	titleIndex int
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify a file to read")
	}
	arg1 := os.Args[1]

	file, err := os.Open(arg1)
	defer file.Close()

	if err != nil {
		log.Fatal("could not open file:", arg1, err)
	}

	scanner := bufio.NewScanner(file)

	// index questions

	var q questionList

	questionTemp := ""
	answerTemp := ""
	titleCounter := 0

	for scanner.Scan() {
		currline := scanner.Text()
		if len(currline) > 2 && currline[:2] == "##" {
			if questionTemp != "" && answerTemp != "" {
				q.questions = append(q.questions, question{questionTemp, answerTemp, false, titleCounter})
			}
			questionTemp = currline[2:]
			answerTemp = ""
		} else if len(currline) > 2 && currline[:2] == "# " {
			q.titles = append(q.titles, currline[2:])
			titleCounter++
		} else if (len(currline) > 2 && currline[:1] != "#") || currline == "" {
			answerTemp += currline + "\n"
		}
	}

	startShell(q)
}

func startShell(q questionList) {

	reader := bufio.NewReader(os.Stdin)
	message := ""

	fmt.Println("welcome to the flashcard program, press enter to select a random question! :-)")
	fmt.Println("There are", len(q.questions), "in the current file")
	for message != "exit" {
		fmt.Print(">")
		var err error
		message, err = reader.ReadString('\n')
		message = message[:len(message)-1]
		if err != nil {
			log.Fatal("error reading input", err)
		}
		//fmt.Printf(message + "\n")
		message = handleInput(message, q)
	}
}

func handleInput(input string, questions questionList) string {
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

func printIndexedQuestion(questions questionList, index int) {
	max := len(questions.questions) - 1
	index = index % max //make sure it is withing bounds of array

	fmt.Println(index, questions.questions[index].question)

	//press anything to continue
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	answer := questions.questions[index].answer

	if answer == "" {
		fmt.Println("no answer specified")
	} else {
		fmt.Println(answer)
	}
}

func printRandomQuestion(questions questionList) {
	max := len(questions.questions) - 1
	randIndex := rand.Intn(max)

	//make sure that the question has not already been shown in this session
	//if questions[randIndex].answered {
	//	for !questions[randIndex].answered {
	//		randIndex = rand.Int() % max
	//	}
	//}

	randQuestion := questions.questions[randIndex]

	fmt.Println(randIndex, randQuestion.question)

	//press anything to continue
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	answer := randQuestion.answer

	if answer == "" {
		fmt.Println("no answer specified")
	} else {
		fmt.Println(answer)
	}
}
