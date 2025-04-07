package models

type Round struct {
	ID        string
	SessionID string
	Number    int
}

var dummyRounds = map[string][]Round{
	"1": {
		{ID: "r1", SessionID: "1", Number: 1},
		{ID: "r2", SessionID: "1", Number: 2},
	},
	"2": {},
}

func GetRoundsBySessionID(sessionID string) ([]Round, error) {
	return dummyRounds[sessionID], nil
}
