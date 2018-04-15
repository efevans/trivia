package trivia

import "time"

// Game conducts a game of triviaa
type Game struct {
	Writer         Writer
	QuestionGetter QuestionGetter

	questions []Question
	players   map[string]int

	currQuestion       *Question
	guessCh            chan guess
	hasQuestionRunning bool
	isRunning          bool
}

// Question defines the text and answer for a trivia question
type Question struct {
	Text   string
	Answer string
}

type player struct {
	name  string
	score int
}

type guess struct {
	player string
	guess  string
}

// QuestionGetter specifies a method to get questions for trivia
type QuestionGetter interface {
	GetQuestions() []Question
}

// Writer is an interface that implements a writing method that will be used by trivia to output the game
type Writer interface {
	write(string)
}

// Start begins the game of trivia
func (game *Game) Start() {
	// Initialize stuff like guessing channel and list of questions
	game.isRunning = true
	game.guessCh = make(chan guess, 10)
	game.players = make(map[string]int)
	game.questions = game.QuestionGetter.GetQuestions()

	// Go through each question
	for _, currQuestion := range game.questions {
		game.Writer.write("Next question in 3, 2, 1...")
		time.Sleep(time.Second * 3)
		game.Writer.write(currQuestion.Text)
		game.hasQuestionRunning = true

		// accept guesses in the allowed amount of time
	GuessTime:
		for {
			select {
			case guess := <-game.guessCh:
				if game.checkGuess(guess) {
					break GuessTime
				}
			case <-time.After(time.Second * 10):
				break GuessTime
			}
		}

		game.hasQuestionRunning = false
		time.Sleep(time.Second * 2)
	}
}

// Guess lets a player make a guess on the current answer. Answers are only allowed
func (game *Game) Guess(answer string, player string) {
	if game.isRunning && game.hasQuestionRunning {
		newGuess := guess{guess: answer, player: player}
		game.guessCh <- newGuess
	}
}

func (game *Game) checkGuess(guess guess) bool {
	// TODO: Track the guesser as having answered this question

	if game.currQuestion.Text == guess.guess {
		// TODO: say the player got it right and give them their score
		game.players[guess.player]++
		return true
	}

	// TODO: say the player sucks
	game.players[guess.player]--
	return false
}

func (player *player) addToScore(value int) {
	player.score += value
}

func (player *player) subractFromScore(value int) {
	player.score -= value
}
