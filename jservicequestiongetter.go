package trivia

// JServiceQuestionGetter implements the QuestionGetter interface using the jservice api
type JServiceQuestionGetter struct {
	count int
}

// GetQuestions gets a collection of questions from the jservice api
func (*JServiceQuestionGetter) GetQuestions() []Question {
	return nil
}
