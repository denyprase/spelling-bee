CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    display_time INTEGER NOT NULL,
    answer_time INTEGER NOT NULL
);


ALTER TABLE sessions
ADD COLUMN created_at TIMESTAMPTZ DEFAULT NOW(),
ADD COLUMN updated_at TIMESTAMPTZ DEFAULT NOW();

INSERT INTO sessions (name, display_time, answer_time)
VALUES
    ('Spelling Session 1', 5, 10),
    ('Spelling Session 2', 6, 12),
    ('Spelling Session 3', 4, 8);

CREATE TABLE rounds (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) not null,
    session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    limit_length BOOLEAN NOT NULL DEFAULT false,
    length INTEGER NOT NULL,
    playing BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Assuming sessions with IDs 1 and 2 already exist
INSERT INTO rounds (session_id, name, limit_length, length)
VALUES 
  (1, 'Introduction Round',true, 3),
  (1, 'Final Round', true, 4);


CREATE TABLE words (
    id SERIAL PRIMARY KEY,
    round_id INTEGER NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    text VARCHAR(50) NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO words (round_id, text) VALUES
(1, 'cat'),
(1, 'sun'),
(1, 'map');

INSERT INTO words (round_id, text) VALUES
(2, 'frog'),
(2, 'milk'),
(2, 'wind'),
(2, 'book');
