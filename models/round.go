package models

import "time"

type Round struct {
	ID          int
	SessionID   int
	Name        string
	LimitLength bool
	Length      int
	Playing     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (db *DB) CreateRound(sessionID int, name string) error {
	now := time.Now()
	_, err := db.Conn.Exec(`
		INSERT INTO rounds (session_id, name, limit, length)
		VALUES ($1, $2, $3, $4)
	`, sessionID, name, now, now)
	return err
}

func (db *DB) GetRoundsBySessionID(sessionID int) ([]Round, error) {
	rows, err := db.Conn.Query(`
		SELECT id, session_id, name, limit_length, length, playing
		FROM rounds
		WHERE session_id = $1
		ORDER BY id ASC
	`, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rounds []Round
	for rows.Next() {
		var r Round
		if err := rows.Scan(&r.ID, &r.SessionID, &r.Name, &r.LimitLength, &r.Length, &r.Playing); err != nil {
			return nil, err
		}
		rounds = append(rounds, r)
	}

	return rounds, nil
}

func (db *DB) GetRoundByID(id int) (*Round, error) {
	var r Round
	err := db.Conn.QueryRow(`
		SELECT id, session_id, name, limit_length, length, playing
		FROM rounds
		WHERE id = $1
	`, id).Scan(&r.ID, &r.SessionID, &r.Name, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *Round) WordsCount() int {
	return 1000
}

func (r *Round) PlayingStatus() string {
	if r.Playing {
		return "In Progress"
	}
	return "Not Started"
}
