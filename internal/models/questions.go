// Filename: internal/models/questions.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// let model the question table
type Question struct {
	QuestionID int64
	Body       string
	CreatedAt  time.Time
}

// Setup dependency injection
type QuestionModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// TODO
func (m *QuestionModel) Get() (*Question, error) {
	var q Question

	statement := `
				SELECT question_id, body
				FROM questions ORDER BY RANDOM()
				LIMIT 1 
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.QuestionID, &q.Body)
	if err != nil {
		return nil, err
	}
	return &q, err
}
