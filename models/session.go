package models

import "fmt"

type Session struct {
	ID          string
	Name        string
	DisplayTime int
	AnswerTime  int
}

var dummySessions = []Session{
	{ID: "1", Name: "Morning Buzz", DisplayTime: 3, AnswerTime: 10},
	{ID: "2", Name: "Afternoon Jam", DisplayTime: 5, AnswerTime: 8},
}

func GetSessionByID(id string) (Session, error) {
	for _, s := range dummySessions {
		if s.ID == id {
			return s, nil
		}
	}
	return Session{}, fmt.Errorf("session not found")
}

func GetSessions() ([]Session, error) {
	return dummySessions, nil
}
