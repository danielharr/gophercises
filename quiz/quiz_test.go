package main

import (
	"strings"
	"testing"
)

func TestQuizCorrectAnswers(t *testing.T) {
	answers := `10
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
	q.Start(strings.NewReader(answers))
	if q.correctAnswers != expected {
		t.Errorf("expected %d correct answers, got %d", expected, q.correctAnswers)
	}
}
