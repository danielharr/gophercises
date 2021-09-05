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
	"time"
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

func (q *Quiz) Start(r io.Reader, timeoutSec int) {
	quizDone := make(chan bool)
	br := bufio.NewReader(r)
	fmt.Printf("You have %d seconds. Press enter to start.\n", timeoutSec)
	br.ReadString('\n')
	timer := time.NewTimer(time.Duration(timeoutSec) * time.Second)
	go q.readAnswers(br, quizDone)

	select {
	case <-timer.C:
		fmt.Println("\nTime is up!")
	case <-quizDone:
	}
	q.printCorrectAnswerCount()
}

func (q *Quiz) readAnswers(r io.Reader, done chan bool) {
	reader := bufio.NewReader(r)
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
	done <- true
}

func (q Quiz) printCorrectAnswerCount() {
	fmt.Printf("%d correct answers out of %d questions\n",
		q.correctAnswers, len(q.questions))
}

var filePath *string

func main() {
	filePath = flag.String("f", "problems.csv", "Quiz csv file")
	timeoutSec := flag.Int("t", 30, "Quiz time")
	flag.Parse()

	q := NewQuiz(*filePath)
	q.Start(os.Stdin, *timeoutSec)
}
