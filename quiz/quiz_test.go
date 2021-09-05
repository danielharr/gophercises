package main

import (
	"strings"
	"testing"
)

func TestQuizCorrectAnswers(t *testing.T) {
	answers := `
10
10
2
11
3
14
4
5
6
5
6
6
7
`
	expected := 13
	q := NewQuiz("problems.csv")
	q.Start(strings.NewReader(answers), 3)
	if q.correctAnswers != expected {
		t.Errorf("expected %d correct answers, got %d", expected, q.correctAnswers)
	}
}

func TestQuizIncorrectAnswers(t *testing.T) {
	answers := `
10
10
2
11
`
	expected := 4
	q := NewQuiz("problems.csv")
	q.Start(strings.NewReader(answers), 3)
	if q.correctAnswers != expected {
		t.Errorf("expected %d correct answers, got %d", expected, q.correctAnswers)
	}
}

func TestQuizTimeout(t *testing.T) {
	answers := `
10
10
2
11
3
14
4
5
6
5
6
6
7
`
	expected := 0
	q := NewQuiz("problems.csv")
	q.Start(strings.NewReader(answers), 0)
	if q.correctAnswers != expected {
		t.Errorf("expected %d correct answers, got %d", expected, q.correctAnswers)
	}
}
