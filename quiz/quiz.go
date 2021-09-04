package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type question struct {
	text   string
	answer string
}

func (q question) isCorrect(answer string) bool {
	answer = strings.TrimSuffix(answer, "\n")
	return answer == q.answer
}

type Quiz struct {
	questions      []question
	correctAnswers int
}

func NewQuiz(filePath string) *Quiz {
	q := new(Quiz)
	q.correctAnswers = 0
	q.questions = make([]question, 0)
	q.parseQuestionsFile(filePath)
	return q
}

func (q *Quiz) parseQuestionsFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("cannot open file: %s", filePath)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	r := csv.NewReader(buf)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		q.questions = append(q.questions, question{record[0], record[1]})
	}
}

func (q *Quiz) Start() {
	reader := bufio.NewReader(os.Stdin)
	for _, question := range q.questions {
		fmt.Printf("Question: %s\n", question.text)
		fmt.Print("Your answer: ")
		answer, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if question.isCorrect(answer) {
			q.correctAnswers++
		}
	}
	q.printCorrectAnswerCount()
}

func (q Quiz) printCorrectAnswerCount() {
	fmt.Printf("%d correct answers out of %d questions\n",
		q.correctAnswers, len(q.questions))
}

var filePath *string

func init() {
	filePath = flag.String("f", "problems.csv", "Quiz csv file")
	flag.Parse()
}

func main() {
	q := NewQuiz(*filePath)
	q.Start()
}
