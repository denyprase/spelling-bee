package models

import (
	"database/sql"
	"time"
)

type Session struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DisplayTime int       `json:"display_time"`
	AnswerTime  int       `json:"answer_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (db *DB) GetSessions() ([]Session, error) {
	rows, err := db.Conn.Query("SELECT id, name, display_time, answer_time FROM sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []Session
	for rows.Next() {
		var s Session
		if err := rows.Scan(&s.ID, &s.Name, &s.DisplayTime, &s.AnswerTime); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, nil
}

func (db *DB) InsertSession(name string, displayTime, answerTime int) (*Session, error) {
	query := `
		INSERT INTO sessions (name, display_time, answer_time)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	err := db.Conn.QueryRow(query, name, displayTime, answerTime).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &Session{
		ID:          id,
		Name:        name,
		DisplayTime: displayTime,
		AnswerTime:  answerTime,
	}, nil
}

func (db *DB) UpdateSession(sessionID int, name string, displayTime, answerTime int) (*Session, error) {
	now := time.Now()

	query := `UPDATE sessions
              SET name = $1, display_time = $2, answer_time = $3, updated_at = $4
              WHERE id = $5 RETURNING id, name, display_time, answer_time, created_at, updated_at`

	var session Session
	err := db.Conn.QueryRow(query, name, displayTime, answerTime, now, sessionID).
		Scan(&session.ID, &session.Name, &session.DisplayTime, &session.AnswerTime, &session.CreatedAt, &session.UpdatedAt)

	if err != nil {
		return nil, err
	}

	session.UpdatedAt = now

	return &session, nil
}

func (db *DB) GetSessionByID(sessionID int) (*Session, error) {
	query := `SELECT id, name, display_time, answer_time, created_at, updated_at
              FROM sessions WHERE id = $1`

	var session Session
	err := db.Conn.QueryRow(query, sessionID).Scan(&session.ID, &session.Name, &session.DisplayTime,
		&session.AnswerTime, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if no session found with the given ID
		}
		return nil, err // Return error if there was an issue querying the database
	}

	return &session, nil
}
