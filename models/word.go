package models

type Word struct {
	ID      int
	RoundID int
	Text    string
	Used    bool
}

func (db *DB) InsertWord(roundID int, text string) (*Word, error) {
	query := `
		INSERT INTO words (round_id, text)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	var used bool
	err := db.Conn.QueryRow(query, roundID, text).Scan(&id, &used)
	if err != nil {
		return nil, err
	}
	return &Word{
		ID:      id,
		RoundID: roundID,
		Text:    text,
		Used:    used,
	}, nil
}

func (db *DB) GetWordsByRoundID(roundID int) ([]Word, error) {
	query := `
		SELECT id, round_id, text, used
		FROM words
		WHERE round_id = $1
		ORDER BY id ASC
	`
	rows, err := db.Conn.Query(query, roundID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []Word
	for rows.Next() {
		var w Word
		if err := rows.Scan(&w.ID, &w.RoundID, &w.Text, &w.Used); err != nil {
			return nil, err
		}
		words = append(words, w)
	}

	return words, rows.Err()
}

// func GetWords() ([]Word, error) {
// 	return dummyWords, nil
// }

// func GetWordsByRoundID(roundID string) ([]Word, error) {
// 	var words []Word
// 	for _, word := range dummyWords {
// 		if word.RoundID == roundID {
// 			words = append(words, word)
// 		}
// 	}
// 	if len(words) == 0 {
// 		return nil, fmt.Errorf("no words found for round ID: %s", roundID)
// 	}
// 	return words, nil
// }
