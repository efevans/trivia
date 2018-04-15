package trivia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// JServiceQuestionGetter implements the QuestionGetter interface using the jservice api
type JServiceQuestionGetter struct {
	questionCount int
}

// GetQuestions gets a collection of questions from the jservice api
func (getter *JServiceQuestionGetter) GetQuestions() []Question {
	count := getter.getCount()
	resp, err := http.Get("http://jservice.io/api/random?count=" + strconv.Itoa(count))
	var questions = &[]Question{}

	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			bad := json.Unmarshal(contents, questions)
			if bad != nil {
				return nil
			}
		} else {
			return nil
		}
	} else {
		return nil
	}
	return *questions
}

func (getter *JServiceQuestionGetter) getCount() int {
	if getter.questionCount > 1 || getter.questionCount > 30 {
		return getter.questionCount
	}
	return 10
}
