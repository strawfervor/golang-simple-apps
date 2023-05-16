/*
Math problems quiz, this program reads problems from csv file (or generete them) and ask user to solve them.
based on: Quiz Game by calhoun.io
written by: StrawFervor
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// open questions file, read it, and return as [][]array
func readQuestions(questionsFile *string) [][]string {
	questionsOpen, err := os.Open(*questionsFile)
	if err != nil {
		fmt.Printf("Failed to open file %s\n", *questionsFile)
		os.Exit(0)
	}

	questionsReader := csv.NewReader(questionsOpen)
	questions, err := questionsReader.ReadAll() //this 2 lines are creating new csv reader and reads all lines
	if err != nil {
		fmt.Println("Something went wrong with processing questions file :c")
		os.Exit(0)
	}
	return questions
}

type question struct { //structure that holds questions with form that is easy to work on
	q string
	a string
}

// parse questions [][]array to single array filled with items contains custom struct called question
func parseQuestions(questions [][]string) []question {
	var questionsNo int = len(questions)
	quizQuestions := make([]question, questionsNo)
	if questionsNo < 1 {
		fmt.Println("Questions file doesn't contain any valid questions")
		os.Exit(0)
		return nil
	} else {
		for i, q := range questions {
			quizQuestions[i].q = q[0]
			quizQuestions[i].a = q[1]
		}
		return quizQuestions
	}
}

// start quiz with prepared []question
func askQuestion(questions []question, timeFlag *int) int {
	quizTimer := time.NewTimer(time.Duration(*timeFlag) * time.Second) //create timer
	counter := 0                                                       //points counter
	for _, q := range questions {                                      //iterating throught elements of questions array
		select { //switch case that is checking if timer timeouted
		case <-quizTimer.C:
			fmt.Printf("Timeout, ")
			return counter
		default:
			fmt.Println(q.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if strings.TrimSpace(q.a) == strings.TrimSpace(answer) {
				fmt.Println("Correct!")
				counter++
			} else {
				fmt.Println("Wrong!")
			}
		}
	}
	return counter
}

func main() {
	questionsFile := flag.String("questions", "questions.csv", "file with questions and answers (format: questions,answer)") //flag allow to change default question "database" in terminal
	timeFlag := flag.Int("timeout", 30, "time to answer in seconds")                                                         //define custom timeout
	flag.Parse()

	questions := readQuestions(questionsFile)
	quizQuestions := parseQuestions(questions)

	var counter int
	counter = askQuestion(quizQuestions, timeFlag) //quiz function is called as output for counter varible
	fmt.Printf("Correct answers: %d of %d questions", counter, len(quizQuestions))
}
