package models

type Round struct {
	ID          string
	SessionID   string
	LimitLength bool
	Length      int
	Playing     bool
}

var dummyRounds = map[string][]Round{
	"1": {
		{ID: "r1", SessionID: "1", LimitLength: true, Length: 3, Playing: true},
		{ID: "r2", SessionID: "1", LimitLength: true, Length: 4, Playing: false},
	},
	"2": {},
}

func GetRoundsBySessionID(sessionID string) ([]Round, error) {
	return dummyRounds[sessionID], nil
}

func (r *Round) WordsCount() int {
	words, _ := GetWordsByRoundID(r.ID)
	return len(words)
}

func (r *Round) PlayingStatus() string {
	if r.Playing {
		return "In Progress"
	}
	return "Not Started"
}
