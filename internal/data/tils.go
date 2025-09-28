package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrTilNotExists = errors.New("Til No exists")
)

type TIL struct {
	IDs       int
	Title     string
	Content   string
	CreatedAt time.Time
}

type TILModel struct {
	DB *sql.DB
}

func (m *TILModel) Insert(title, content string) (int, error) {

	query := `INSERT INTO tils (title,content,created_at) VALUES ($1,$2,NOW()) RETURNING id`

	var id int
	err := m.DB.QueryRow(query, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *TILModel) Latest() ([]*TIL, error) {

	query := `SELECT id, title,content,created_at
			 FROM tils
			 ORDER BY created_at DESC LIMIT 10`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tils := []*TIL{}

	for rows.Next() {
		t := &TIL{}
		err := rows.Scan(&t.IDs, &t.Title, &t.Content, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tils = append(tils, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tils, nil
}

func (m *TILModel) Get(id int) (*TIL, error) {

	query := `SELECT id, title, content, created_at FROM tils WHERE id = $1`

	var t TIL

	err := m.DB.QueryRow(query, id).Scan(
		&t.IDs,
		&t.Title,
		&t.Content,
		&t.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTilNotExists
		}
		return nil, err
	}
	return &t, nil
}
