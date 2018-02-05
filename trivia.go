package trivia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// QuestionType indicates the type of question (i.e. multiple choice, free type)
type QuestionType int

// Specifies valid question types
const (
	MultipleChoice QuestionType = 0
	FreeType       QuestionType = 1
)

func (qType QuestionType) String() string {
	types := [...]string{
		"MultipleChoice",
		"FreeType"}

	if qType < MultipleChoice || qType > FreeType {
		return "Unknown"
	}

	return types[qType]
}

// Question contains the information for a question
type Question struct {
	ID         int
	Question   string
	Value      int
	Answer     string
	isAnswered bool
}

// GetQuestion gets a trivia question
func GetQuestion() Question {
	question := getQuestion()
	return question
}

// Read prints out the question, and readies the Question for answering
func (question Question) Read() {
	fmt.Println(question.Question)
}

// Guess checks if a propsed answer is correct
func (question Question) Guess(answer string) {
	if !question.isAnswered {
		fmt.Println("Propsing answer with value: " + answer)
		correct := answer == question.Answer // might need to do some trimming of non-alphanumerics and lowercasing the guess and answer to avoid questions with odd answers

		if correct {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect. Correct Answer was: " + question.Answer)
		}
	}
}

func getQuestion() Question {
	resp, err := http.Get("http://jservice.io/api/random?count=1")
	var questions = &[]Question{}

	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			bad := json.Unmarshal(contents, questions)
			if bad != nil {
				fmt.Println("AHHHH")
			}
		} else {
			fmt.Println("woops again leeel")
		}
	} else {
		fmt.Println("woops lel")
	}

	(*questions)[0].isAnswered = false
	return (*questions)[0]
}
