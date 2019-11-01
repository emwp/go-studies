package mysql

import (
	"database/sql"

	"github.com/emwp/go-studies/snippetbox/pkg/models"
)

type SnippedModel struct {
	DB *sql.DB
}

func (m *SnippedModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippedModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippedModel) Latest(id int) ([]*models.Snippet, error) {
	return nil, nil
}
