package repositories

import (
	"database/sql"
	"errors"
	"fmt"
)

type ExampleModel struct {
	ID   int
	Name string
}

type ExampleRepository struct {
	db *sql.DB
}

func NewExampleRepository(db *sql.DB) *ExampleRepository {
	return &ExampleRepository{db: db}
}

func (r *ExampleRepository) GetExample(id int) (*ExampleModel, error) {
	var example ExampleModel
	query := "SELECT id, name FROM examples WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&example.ID, &example.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("example not found")
		}
		return nil, fmt.Errorf("error querying example: %v", err)
	}
	return &example, nil
}

func (r *ExampleRepository) SaveExample(example *ExampleModel) error {
	query := "INSERT INTO examples (name) VALUES (?)"
	_, err := r.db.Exec(query, example.Name)
	if err != nil {
		return fmt.Errorf("error saving example: %v", err)
	}
	return nil
}