package trivia

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
	ID   int
	Type QuestionType
}

// GetQuestion gets a trivia question
func GetQuestion() Question {
	return Question{ID: 4}
}
